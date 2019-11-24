<template>
  <div>
    <div ref="watermark" class="video-watermark" @contextmenu.prevent>
      <div>{{ $store.state.user.username }}</div>
      <svg-icon icon-class="logo" class="logo" />
      <div>{{ $t("copyright") }}</div>
    </div>
    <video-player
      :class="{
        'vjs-big-play-centered': centerPlayButton
      }"
      :options="options"
      width="100%"
      @contextmenu.native.prevent
    />
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "copyright":"版权所有"
  },
  "en-US": {
    "copyright":"Copyright Reserved."
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import { videoPlayer } from "vue-video-player";
import "video.js/dist/video-js.css";

export default {
  name: "VideoPlayer",
  components: {
    videoPlayer
  },
  props: {
    centerPlayButton: {
      type: Boolean,
      default: true
    },
    src: {
      type: String,
      default: () => ""
    }
  },
  computed: {
    ...mapGetters(["settings"]),
    options() {
      return {
        controls: true,
        language: this.$i18n.locale,
        autoplay: false,
        fluid: true,
        controlBar: {
          volumeBar: false,
          volumePanel: false,
          currentTimeDisplay: false,
          fullscreenToggle: false
        },
        sources: [
          {
            type: "video/mp4",
            src: this.settings.publicURL + this.src
          }
        ],
        languages: {
          "zh-CN": {
            "The media could not be loaded, either because the server or network failed or because the format is not supported.":
              "暂无数据",
            "Play Video": "播放视频",
            Fullscreen: "全屏",
            Play: "播放",
            Pause: "暂停",
            Replay: "重播"
          },
          "en-US": {
            "The media could not be loaded, either because the server or network failed or because the format is not supported.":
              "No Data",
            "Play Video": "Play Video",
            Fullscreen: "Fullscreen",
            Play: "Play",
            Pause: "Pause",
            Replay: "Replay"
          }
        }
      };
    }
  }
};
</script>


<style lang="scss" scoped>
.video-watermark {
  font-size: 1vw;
  font-weight: bolder;
  color: rgba($color: #000, $alpha: 0.5);
  text-align: center;
  user-select: none;
  position: relative;
  top: 10px;
  right: 10px;
  float: right;
  z-index: 500;
  .logo {
    font-size: 3em;
  }
}
</style>