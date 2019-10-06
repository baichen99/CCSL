<template>
  <div class="lexical-database">
    <sign-search @sign-selected="searchBySign" />

    <div class="search">
      <h3>国家通用手语比对语料库</h3>

      <video-search-input
        v-model="params"
        @search-clicked="searchByButton"
      />

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
        word: "",
        chinese: "",
        english: "",
        gender: "",
        region: undefined,
        leftSign: "", // 左手手形
        rightSign: "", // 右手手形
        sign: "", // 任意手形
        pos: "", // 词性
        initial: "",
        constructType: "",
        constructWords: ""
      }
    };
  },
  methods: {
    clearParams() {
      this.params.word = "";
      this.params.gender = "";
      this.params.region = "";
      this.params.leftSign = "";
      this.params.rightSign = "";
      this.params.sign = "";
      this.params.pos = "";
      this.params.chinese = "";
      this.params.english = "";
      this.params.initial = "";
      this.params.constructType = "";
      this.params.constructWords = "";
    },
    getData() {
      GetLexicalVideosList(this.params).then(res => {
        this.videos = res.data;
        this.params.page = res.page;
        this.params.limit = res.limit;
        this.total = res.total;
        if (this.total === 0) {
          this.$message("没有找到相关的数据哦～");
        }
      });
    },
    changePage(page) {
      this.params.page = page;
      this.getData();
    },
    searchByButton() {
      this.params.word = "";
      this.params.sign = "";
      this.params.page = 1;
      this.getData();
    },
    searchByWord(word) {
      this.clearParams();
      this.params.word = word;
      this.params.page = 1;
      this.getData();
    },
    searchBySign(sign) {
      this.clearParams();
      this.params.sign = sign;
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
    h3 {
      text-align: center;
    }
  }
}
</style>
