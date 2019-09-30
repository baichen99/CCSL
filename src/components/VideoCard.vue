<template>
  <el-card
    shadow="hover"
    :body-style="{ padding: '0px' }"
  >
    <video-player
      class="vjs-big-play-centered"
      :options="{
        sources: [{ type: 'video/mp4', src: 'https://ccsl.shu.edu.cn/public/videos/' + video.videoPath }]
      }"
      @contextmenu.native.prevent=""
    />
    <div
      class="tags-container"
      @click="onVideoClick(video)"
    >
      <div class="tags">
        <span
          v-if="showRegion"
          class="tag-value"
        >{{ video.region }}</span>
        <span class="tag-value">{{ video.chinese.split("（")[0] }}</span>
      </div>
    </div>
  </el-card>
</template>

<script>
export default {
  name: "VideoCard",
  props: {
    video: {
      type: Object,
      required: true,
      default: () => ({})
    },
    showRegion: {
      type: Boolean,
      default: true
    }
  },
  methods: {
    onVideoClick(video) {
      this.$emit("on-video-click", video);
    }
  }
};
</script>

<style lang="scss" scoped>
.el-card {
  background: #efefef;
  .tags-container {
    cursor: pointer;
    padding: 5px 0;
  }
  .tags {
    padding: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
  }
  .tag-label {
    font-weight: bold;
  }
}
</style>
