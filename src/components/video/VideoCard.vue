<template>
  <el-card shadow="hover" :body-style="{ padding: '0px' }">
    <video-player :src="video.videoPath " />
    <div class="tags-container" @click="onVideoClick(video)">
      <div class="tags">
        <span v-if="showRegion" class="tag-value">{{ video.performer.region.name }}</span>
        <span
          v-if="$i18n.locale==='zh-CN'"
          class="tag-value"
          v-html="$options.filters.addNumberSup(video.lexicalWord.chinese.split('（')[0])"
        />
        <span v-else-if="$i18n.locale==='en-US'" class="tag-value">{{ video.lexicalWord.english }}</span>
      </div>
    </div>
  </el-card>
</template>

<script>
import VideoPlayer from "@/components/video/VideoPlayer";
import { mapGetters } from "vuex";
export default {
  name: "VideoCard",
  components: {
    VideoPlayer
  },
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
  computed: {
    ...mapGetters(["settings"])
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
