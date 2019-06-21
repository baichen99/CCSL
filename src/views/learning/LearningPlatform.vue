<template>
  <div class="learning-platform">
    <div class="search">
      <h3>通用手语学习平台</h3>
      <el-card class="initial-search">
        <div class="initial">
          <el-collapse accordion>
            <el-collapse-item
              v-for="letter in $store.getters.letters"
              :key="letter"
              :title="letter"
            >
              <el-link
                v-for="word in words[letter]"
                :key="word.id"
                type="primary"
                @click="searchByWord(word.id)"
              >{{ word.chinese }}</el-link>
            </el-collapse-item>
          </el-collapse>
        </div>
      </el-card>
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
            :lg="8"
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
      limit: 12,
      total: 0
    };
  },
  computed: {
    words() {
      return this.$store.state.signlang.words;
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
        this.videos = res.data.videos;
        this.page = res.data.page;
        this.limit = res.data.limit;
        this.total = res.data.total;
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
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
  padding: 10px;

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
    .el-link {
      padding: 5px 10px;
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
