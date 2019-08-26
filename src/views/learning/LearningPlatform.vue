<template>
  <div class="learning-platform">
    <el-card class="sidebar">
      <h3 class="title">音序检索</h3>
      <div class="initial">
        <el-collapse accordion>
          <el-collapse-item
            v-for="letter in $store.getters.letters.slice(0,11)"
            :key="letter"
            :title="letter"
          >
            <div
              v-for="word in words[letter]"
              :key="word.id"
            >
              <el-link
                type="primary"
                @click="searchByWord(word.id)"
              >{{ word.chinese }}</el-link>
            </div>
          </el-collapse-item>
        </el-collapse>
      </div>
    </el-card>
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
      <el-card
        v-if="videos.length!==0"
        shadow="never"
      >
        <el-row :gutter="20">
          <el-col
            v-for="video in videos"
            :key="video.id"
            :sm="24"
            :md="12"
            :lg="12"
          >
            <video-card
              :show-region="false"
              :video="video"
              @on-video-click="showDetailModal"
            />
          </el-col>
        </el-row>
        <el-pagination
          background
          layout="total, prev, pager, next"
          :current-page.sync="page"
          :total="total"
          :page-size="limit"
          @current-change="changePage"
        />
        <el-dialog
          :visible.sync="showDetail"
          top="40px"
        >
          <video-detail :video="videoDetail" />
        </el-dialog>
      </el-card>
    </div>
    <el-card class="sidebar">
      <h3 class="title">音序检索</h3>
      <div class="initial">
        <el-collapse accordion>
          <el-collapse-item
            v-for="letter in $store.getters.letters.slice(11)"
            :key="letter"
            :title="letter"
          >
            <div
              v-for="word in words[letter]"
              :key="word.id"
            >
              <el-link
                type="primary"
                @click="searchByWord(word.id)"
              >{{ word.chinese }}</el-link>
            </div>
          </el-collapse-item>
        </el-collapse>
      </div>
    </el-card>
  </div>
</template>

<script>
import VideoCard from "@/components/VideoCard.vue";
import VideoDetail from "@/components/VideoDetail.vue";
import { getUniversalContrastVideos } from "@/api/videos";

export default {
  name: "LearningPlatform",
  components: {
    VideoCard,
    VideoDetail
  },
  data() {
    return {
      showDetail: false,
      videoDetail: {},
      advancedSearch: false,
      videos: [],
      wordID: "",
      keyword: "",
      page: 1,
      limit: 4,
      total: 0
    };
  },
  computed: {
    words() {
      return this.$store.state.sign.words;
    }
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
    showDetailModal(video) {
      this.videoDetail = video;
      this.showDetail = true;
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

  .sidebar {
    width: 18%;
    height: 600px;
    overflow: scroll;
    .title {
      text-align: center;
    }
  }

  .initial-search {
    height: 300px;
    overflow: scroll;
  }

  .initial {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    .el-collapse {
      width: 100%;
    }
  }

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

    .el-card {
      margin: 10px 0;
      .el-select {
        margin: 5px 0;
      }
    }

    .el-pagination {
      text-align: center;
    }
  }
}
</style>
