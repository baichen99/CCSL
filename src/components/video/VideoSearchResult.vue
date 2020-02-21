<template>
  <el-card v-if="videos.length!==0" shadow="hover">
    <el-row :gutter="20">
      <el-col v-for="video in videos" :key="video.id" :sm="sm" :md="md" :lg="lg">
        <transition name="el-collapse-transition">
          <video-card :show-region="showRegion" :video="video" @on-video-click="showDetailModal" />
        </transition>
      </el-col>
    </el-row>
    <el-pagination
      background
      layout="total, prev, pager, next"
      :current-page.sync="pageNumber"
      :total="total"
      :page-size.sync="limit"
      @current-change="$emit('page-change', page)"
    />
    <el-dialog
      v-if="showDetail"
      :visible.sync="showDetail"
      :show-close="false"
      :destroy-on-close="true"
      top="100px"
    >
      <video-detail :video="videoDetail" />
    </el-dialog>
  </el-card>
</template>

<script>
import VideoCard from "@/components/video/VideoCard.vue";
import VideoDetail from "@/components/video/VideoDetail.vue";
export default {
  name: "VideoSearchResult",
  components: {
    VideoCard,
    VideoDetail
  },
  model: {
    prop: "page",
    event: "update"
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
  computed: {
    pageNumber: {
      get() {
        return this.page;
      },
      set(page) {
        this.$emit("update", page);
      }
    }
  },
  created() {
    this.$nextTick(() => {
      this.$store.dispatch("data/getPerformers");
      this.$store.dispatch("data/getLexicons");
      this.$store.dispatch("data/getHandshapes");
    });
  },
  methods: {
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