<template>
  <div class="universal-contrast">
    <sign-selector />

    <div class="search">
      <h3>国家通用手语比对语料库</h3>
      <el-input
        v-model="keyword"
        class="options"
        placeholder="请输入内容"
        @keyup.enter.native="searchByButton"
      >
        <el-select
          slot="prepend"
          v-model="keywordType"
          placeholder="请选择"
          style="width:100px"
        >
          <el-option
            label="中文"
            value="chinese"
          />
          <el-option
            label="English"
            value="english"
          />
        </el-select>
        <el-button
          slot="append"
          icon="el-icon-search"
          @click="searchByButton"
        />
      </el-input>

      <div class="options">
        <span>高级检索</span>
        <el-radio-group v-model="advancedSearch">
          <el-radio :label="false">否</el-radio>
          <el-radio :label="true">是</el-radio>
        </el-radio-group>
      </div>

      <transition
        name="fade"
        :duration="{ enter: 500, leave: 800 }"
      >
        <el-card
          v-if="advancedSearch"
          class="advance"
        >
          <div class="options">
            <span>性别</span>
            <el-radio-group v-model="gender">
              <el-radio label="">无限制</el-radio>
              <el-radio label="男">男</el-radio>
              <el-radio label="女">女</el-radio>
            </el-radio-group>
          </div>
          <div class="options">
            <el-select
              v-model="region"
              clearable
              placeholder="请选择区域"
            >
              <el-option
                v-for="(item,index) in $store.getters.regions"
                :key="index"
                :label="item"
                :value="item"
              >
                {{ item }}
              </el-option>
            </el-select>
            <el-select
              v-model="leftSign"
              clearable
              placeholder="请选择左手手型"
            >
              <el-option
                v-for="(item,index) in $store.getters.sign"
                :key="index"
                :label="item"
                :value="item"
              >
                <span>{{ item }}</span>
                <img
                  style="float: right;height:100%"
                  :src="'sign/'+item+'.jpg'"
                  :alt="item"
                >
              </el-option>
            </el-select>
            <el-select
              v-model="rightSign"
              clearable
              placeholder="请选择右手手型"
            >
              <el-option
                v-for="(item,index) in $store.getters.sign"
                :key="index"
                :label="item"
                :value="item"
              >
                <span>{{ item }}</span>
                <img
                  style="float: right;height:100%"
                  :src="'sign/'+item+'.jpg'"
                  :alt="item"
                >
              </el-option>
            </el-select>
            <el-select
              v-model="type"
              clearable
              placeholder="请选择词性"
            >
              <el-option
                v-for="(item,index) in $store.getters.wordTypes"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
            <el-select
              v-model="constructType"
              clearable
              placeholder="请选择构词方式"
            >
              <el-option
                label="单纯词"
                value="单纯词"
              />
              <el-option
                label="复合词"
                value="复合词"
              />
            </el-select>
            <el-input
              v-model="constructWords"
              clearable
              style="width:200px"
              placeholder="请输入复合词构成词语"
            />
          </div>
        </el-card>
      </transition>

      <el-card
        v-if="videos.length!==0"
        shadow="never"
      >
        <el-row :gutter="20">
          <el-col
            v-for="video in videos"
            :key="video.id"
            :sm="12"
            :md="8"
            :lg="6"
          >
            <video-card
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

    <word-selector />
  </div>
</template>

<script>
import VideoCard from "@/components/VideoCard.vue";
import VideoDetail from "@/components/VideoDetail.vue";
import WordSelector from "@/components/WordSelector.vue";
import SignSelector from "@/components/SignSelector.vue";
import { getUniversalContrastVideos } from "@/api/videos";
import { mapGetters } from "vuex";

export default {
  name: "UniversalContrast",
  components: {
    VideoCard,
    VideoDetail,
    WordSelector,
    SignSelector
  },
  data() {
    return {
      showDetail: false,
      videoDetail: {},
      advancedSearch: false,
      videos: [],
      wordID: "",
      keywordType: "",
      keyword: "",
      page: 1,
      limit: 12,
      total: 0,
      gender: "",
      region: "",
      leftSign: "",
      rightSign: "",
      type: "",
      initial: "",
      constructType: "",
      constructWords: ""
    };
  },
  methods: {
    clearParams() {
      this.wordID = "";
      this.gender = "";
      this.region = "";
      this.leftSign = "";
      this.rightSign = "";
      this.type = "";
      this.keywordType = "";
      this.keyword = "";
      this.initial = "";
      this.constructType = "";
      this.constructWords = "";
    },
    getData() {
      let params = {
        word: this.wordID,
        page: this.page,
        limit: this.limit,
        gender: this.gender,
        leftSign: this.leftSign,
        rightSign: this.rightSign,
        initial: this.initial,
        region: this.region,
        type: this.type,
        constructType: this.constructType,
        constructWords: this.constructWords
      };
      if (this.keywordType === "chinese" || this.keywordType === "") {
        params.chinese = this.keyword;
      } else if (this.keywordType === "english") {
        params.english = this.keyword;
      }
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
    },
    searchBySign(sign) {
      this.clearParams();
      this.rightSign = sign;
      this.searchVideos();
    }
  }
};
</script>

<style lang="scss" scoped>
.universal-contrast {
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
