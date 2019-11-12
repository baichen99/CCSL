<template>
  <div class="lexical-database">
    <sign-search @sign-selected="searchBySign" />

    <div class="search">
      <h2>{{ $t("LexicalDatabase") }}</h2>

      <video-search-input v-model="params" @search="searchByButton" />

      <video-search-result
        :videos="videos"
        :limit="params.limit"
        :page="params.page"
        :total="total"
        :sm="24"
        :md="12"
        :lg="12"
        @change-page="changePage"
      />
    </div>

    <word-search @word-selected="searchByWord" />
  </div>
</template>

<script>
import VideoSearchInput from "@/components/video/VideoSearchInput.vue";
import VideoSearchResult from "@/components/video/VideoSearchResult.vue";
import WordSearch from "@/components/form/WordSearch.vue";
import SignSearch from "@/components/form/SignSearch.vue";

import { GetLexicalVideosList } from "@/api/videos";

export default {
  name: "LexicalDatabase",
  components: {
    VideoSearchResult,
    VideoSearchInput,
    WordSearch,
    SignSearch
  },
  data() {
    return {
      videos: [],
      total: 0,
      params: {
        page: 1,
        limit: 4,
        wordID: "",
        chinese: "",
        english: "",
        gender: "",
        regionID: undefined,
        leftSignID: "", // 左手手形
        rightSignID: "", // 右手手形
        signID: "", // 任意手形
        pos: "", // 词性
        initial: "",
        constructType: "",
        constructWords: "",
        orderBy: "performers.gender"
      }
    };
  },
  methods: {
    clearParams() {
      this.params.word = "";
      this.params.gender = "";
      this.params.regionID = undefined;
      this.params.leftSignID = "";
      this.params.rightSignID = "";
      this.params.signID = "";
      this.params.pos = "";
      this.params.chinese = "";
      this.params.english = "";
      this.params.initial = "";
      this.params.constructType = "";
      this.params.constructWords = "";
    },
    getData() {
      GetLexicalVideosList(this.params, true).then(res => {
        this.videos = res.data;
        this.params.page = res.page;
        this.params.limit = res.limit;
        this.total = res.total;
        if (this.total === 0) {
          this.$notify({
            title: "没有找到相关的数据哦～",
            type: "warning",
            duration: 2000
          });
        }
      });
    },
    changePage(page) {
      this.params.page = page;
      this.getData();
    },
    searchByButton() {
      this.params.wordID = "";
      this.params.signID = "";
      this.params.page = 1;
      this.getData();
    },
    searchByWord(wordID) {
      this.clearParams();
      this.params.wordID = wordID;
      this.params.page = 1;
      this.getData();
    },
    searchBySign(signID) {
      this.clearParams();
      this.params.signID = signID;
      this.params.page = 1;
      this.getData();
    }
  }
};
</script>

<style lang="scss" scoped>
.lexical-database {
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
  }
}
</style>
