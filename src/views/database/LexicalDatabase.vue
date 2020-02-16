<template>
  <div class="lexical-database">
    <handshape-search @handshape-selected="searchByHandshape" />

    <div class="search">
      <h2>{{ $t("LexicalDatabase") }}</h2>

      <video-search-input v-model="params" @search="searchByButton" />

      <video-search-result
        v-model="params.page"
        v-loading="loading"
        :videos="videos"
        :limit="params.limit"
        :total="total"
        :sm="24"
        :md="12"
        :lg="12"
        @page-change="changePage"
      />
    </div>

    <word-search @word-selected="searchByWord" />
  </div>
</template>

<script>
import VideoSearchInput from "@/components/video/VideoSearchInput.vue";
import VideoSearchResult from "@/components/video/VideoSearchResult.vue";
import WordSearch from "@/components/form/WordSearch.vue";
import HandshapeSearch from "@/components/form/HandshapeSearch.vue";

import { GetLexicalVideosList } from "@/api/videos";

export default {
  name: "LexicalDatabase",
  components: {
    VideoSearchResult,
    VideoSearchInput,
    WordSearch,
    HandshapeSearch
  },
  data() {
    return {
      loading: false,
      videos: [],
      total: 0,
      params: {
        page: 1,
        limit: 4,
        lexiconID: "",
        chinese: "",
        english: "",
        gender: "",
        regionID: undefined,
        leftHandshapeID: "", // 左手手形
        rightHandshapeID: "", // 右手手形
        handshapeID: "", // 任意手形
        pos: "", // 词性
        initial: "",
        wordFormation: "",
        morpheme: ""
      }
    };
  },
  methods: {
    clearParams() {
      this.params.lexiconID = "";
      this.params.gender = "";
      this.params.regionID = undefined;
      this.params.leftHandshapeID = "";
      this.params.rightHandshapeID = "";
      this.params.handshapeID = "";
      this.params.pos = "";
      this.params.chinese = "";
      this.params.english = "";
      this.params.initial = "";
      this.params.wordFormation = "";
      this.params.morpheme = "";
    },
    async getData() {
      this.loading = true;
      try {
        const res = await GetLexicalVideosList(this.params);
        if (res.total === 0) {
          this.$notify.info({
            title: this.$t("VideoNotFound"),
            duration: 2000
          });
        }
        this.videos = res.data;
        this.params.page = res.page;
        this.params.limit = res.limit;
        this.total = res.total;
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    changePage() {
      this.getData();
    },
    searchByButton() {
      this.params.lexiconID = "";
      this.params.handshapeID = "";
      this.params.page = 1;
      this.getData();
    },
    searchByWord(lexiconID) {
      this.clearParams();
      this.params.lexiconID = lexiconID;
      this.params.page = 1;
      this.getData();
    },
    searchByHandshape(handshapeID) {
      this.clearParams();
      this.params.handshapeID = handshapeID;
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
