<template>
  <transition name="slide-fade">
    <el-card
      v-if="videos.length!==0"
      shadow="hover"
    >
      <el-row :gutter="20">
        <el-col
          v-for="video in videos"
          :key="video.id"
          :sm="sm"
          :md="md"
          :lg="lg"
        >
          <video-card
            :show-region="showRegion"
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
  </transition>

</template>

<script>
import VideoCard from "@/components/VideoCard.vue";
import VideoDetail from "@/components/VideoDetail.vue";
export default {
  name: "VideoSearchResult",
  components: {
    VideoCard,
    VideoDetail
  },
  props: {
    showRegion: {
      type: Boolean,
      default: true
    },
    sm: {
      type: Number,
      default: 12
    },
    md: {
      type: Number,
      default: 8
    },
    lg: {
      type: Number,
      default: 6
    },
    videos: {
      type: Array,
      default: () => []
    },
    limit: {
      type: Number,
      default: 12
    },
    page: {
      type: Number,
      default: 1
    },
    total: {
      type: Number,
      default: 0
    }
  },
  data: () => ({
    videoDetail: {},
    showDetail: false
  }),
  methods: {
    changePage(page) {
      this.$emit("change-page", page);
    },
    showDetailModal(video) {
      this.videoDetail = video;
      this.showDetail = true;
    }
  }
};
</script>

<style lang="scss" scoped>
.el-card {
  margin: 10px 0;
  .el-select {
    margin: 5px 0;
  }
}

.el-pagination {
  text-align: center;
}
</style>