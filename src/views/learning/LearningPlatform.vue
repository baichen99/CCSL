<template>
  <div class="learning-platform">
    <word-selector
      start="A"
      end="L"
      @word-selected="searchByWord"
    />

    <div class="search">
      <h3>国家通用手语学习平台</h3>

      <el-input
        v-model="keyword"
        class="options"
        placeholder="请输入关键词以搜索"
        @keyup.enter.native="searchByButton"
      >
        <el-button
          slot="append"
          icon="el-icon-search"
          @click="searchByButton"
        />
      </el-input>

      <video-search-result
        :videos="videos"
        :limit="limit"
        :page="page"
        :total="total"
        :sm="24"
        :md="12"
        :lg="12"
        @change-page="changePage"
      />

    </div>

    <word-selector
      start="M"
      end="Z"
      @word-selected="searchByWord"
    />
  </div>
</template>

<script>
import VideoSearchResult from "@/components/VideoSearchResult.vue";
import WordSelector from "@/components/WordSelector.vue";
import { getUniversalContrastVideos } from "@/api/videos";

export default {
  name: "LearningPlatform",
  components: {
    VideoSearchResult,
    WordSelector
  },
  data() {
    return {
      advancedSearch: false,
      videos: [],
      wordID: "",
      keyword: "",
      page: 1,
      limit: 4,
      total: 0
    };
  },
  methods: {
    clearParams() {
      this.wordID = "";
      this.keyword = "";
    },
    getData() {
      let params = {
        word: this.wordID,
        page: this.page,
        limit: this.limit,
        gender: "男",
        region: "通用手语",
        chinese: this.keyword
      };
      getUniversalContrastVideos(params).then(res => {
        this.videos = res.data;
        this.page = res.page;
        this.limit = res.limit;
        this.total = res.total;
        if (this.videos.length === 0) {
          this.$message("没有找到相关的数据哦～");
        }
      });
    },
    searchVideos() {
      this.page = 1;
      this.getData();
    },
    changePage(page) {
      this.page = page;
      this.getData();
    },
    searchByButton() {
      this.wordID = "";
      this.searchVideos();
    },
    searchByWord(wordID) {
      this.clearParams();
      this.wordID = wordID;
      this.searchVideos();
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
    h3 {
      text-align: center;
    }

    .options {
      margin: 5px 0;
      text-align: center;
    }
  }
}
</style>
