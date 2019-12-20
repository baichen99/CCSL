<template>
  <div class="learning-platform">
    <word-search start="A" end="L" @word-selected="searchByWord" />

    <div class="search">
      <h2>{{ $t("learningPlatform") }}</h2>

      <video-search-input v-model="params" :show-advance="false" @search="searchByButton" />

      <video-search-result
        v-model="params.page"
        v-loading="loading"
        :videos="videos"
        :limit="params.limit"
        :total="total"
        :sm="24"
        :md="12"
        :lg="12"
        :show-region="false"
        @page-change="changePage"
      />
    </div>

    <word-search start="M" end="Z" @word-selected="searchByWord" />
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "learningPlatform": "国家通用手语学习平台"
  },
  "en-US": {
    "learningPlatform": "Chinese National Sign Language Learning Platform"
  }
}
</i18n>

<script>
import VideoSearchInput from "@/components/video/VideoSearchInput.vue";
import VideoSearchResult from "@/components/video/VideoSearchResult.vue";
import WordSearch from "@/components/form/WordSearch.vue";
import { GetLexicalVideosList } from "@/api/videos";

export default {
  name: "LearningPlatform",
  components: {
    VideoSearchResult,
    VideoSearchInput,
    WordSearch
  },
  data() {
    return {
      loading: false,
      videos: [],
      total: 0,
      params: {
        chinese: "",
        lexiconID: "",
        gender: "M",
        regionID: 100000,
        page: 1,
        limit: 4
      }
    };
  },
  methods: {
    clearParams() {
      this.params.lexiconID = "";
      this.params.chinese = "";
    },
    getData() {
      let params = this.params;
      GetLexicalVideosList(params)
        .then(res => {
          this.loading = false;
          this.videos = res.data;
          this.params.page = res.page;
          this.params.limit = res.limit;
          this.total = res.total;
          if (this.total === 0) {
            this.$notify.info({
              title: this.$t("VideoNotFound"),
              duration: 2000
            });
          }
        })
        .catch(() => {
          this.loading = false;
        });
    },
    changePage() {
      this.getData();
    },
    searchByButton() {
      this.params.lexiconID = "";
      this.params.page = 1;
      this.getData();
    },
    searchByWord(lexiconID) {
      this.clearParams();
      this.params.lexiconID = lexiconID;
      this.params.page = 1;
      this.getData();
    }
  }
};
</script>

<style lang="scss" scoped>
.learning-platform {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 10px 0;

  .search {
    width: 60%;
    padding: 40px;
    h2 {
      text-align: center;
    }

    .options {
      margin: 5px 0;
      text-align: center;
    }
  }
}
</style>
