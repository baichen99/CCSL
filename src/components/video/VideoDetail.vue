<template>
  <div>
    <div class="video-container">
      <video-player :src="video.videoPath" />
    </div>
    <div class="tags-container">
      <div class="tags">
        <span class="tag-label">{{ $t("Region") }}</span>
        <span class="tag-value">{{ performers[video.performerID].regionID | getRegionName }}</span>
        <span class="tag-label">{{ $t("Chinese") }}</span>
        <span
          class="tag-value word-sup"
          v-html="$options.filters.addNumberSup(lexicons[video.lexiconID].chinese)"
        ></span>
      </div>
      <div class="tags">
        <span class="tag-label">{{ $t("PoS") }}</span>
        <span class="tag-value">
          <el-tag v-for="value in lexicons[video.lexiconID].pos" :key="value" size="small">
            <span>{{ $t($options.filters.getObjectItem(partOfSpeech,value).text) }}</span>
          </el-tag>
        </span>
        <span class="tag-label">{{ $t("English") }}</span>
        <span class="tag-value">{{ lexicons[video.lexiconID].english }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">{{ $t("WordFormation") }}</span>
        <span
          v-if="video.wordFormation"
          class="tag-value"
        >{{ $t($options.filters.getObjectItem(wordFormations,video.wordFormation).text) }}</span>
        <span v-else class="tag-value">
          <el-tag type="info" size="small">{{ $t("NoData") }}</el-tag>
        </span>
        <span class="tag-label">{{ $t("Morpheme") }}</span>
        <span v-if="video.morpheme.length > 0" class="tag-value">
          <el-tag v-for="(item,index) in video.morpheme" :key="index" size="small" class="tag">
            <span class="word-sup" v-html="$options.filters.addNumberSup(item)"></span>
          </el-tag>
        </span>
        <span v-else class="tag-value">
          <el-tag type="info" size="small">{{ $t("NoData") }}</el-tag>
        </span>
      </div>
      <div v-if="video.leftHandshapesID.length !== 0" class="tags">
        <span class="tag-label">{{ $t("LeftHandshape") }}</span>
        <span class="tag-value" style="width:85%">
          <span v-for="item in video.leftHandshapesID" :key="item">
            <el-tooltip :content="handshapes[item].name">
              <img :src="settings.publicURL + handshapes[item].image" :alt="handshapes[item].name" />
            </el-tooltip>
          </span>
        </span>
      </div>
      <div v-if="video.rightHandshapesID.length !== 0" class="tags">
        <span class="tag-label">{{ $t("RightHandshape") }}</span>
        <span class="tag-value" style="width:85%">
          <span v-for="item in video.rightHandshapesID" :key="item">
            <el-tooltip :content="handshapes[item].name" placement="top">
              <img :src="settings.publicURL + handshapes[item].image" :alt="handshapes[item].name" />
            </el-tooltip>
          </span>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import VideoPlayer from "@/components/video/VideoPlayer";
import { mapGetters } from "vuex";
export default {
  name: "VideoDetail",
  components: {
    VideoPlayer
  },
  props: {
    video: {
      type: Object,
      required: true,
      default: () => ({})
    }
  },
  computed: {
    ...mapGetters([
      "wordFormations",
      "partOfSpeech",
      "settings",
      "lexicons",
      "performers",
      "handshapes"
    ])
  }
};
</script>

<style lang="scss" scoped>
.video-container {
  margin: 0 40px;
}

.tags-container {
  padding: 20px 40px 0 40px;
  .tags {
    text-align: center;
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    line-height: 30px;
    .tag-label {
      width: 15%;
      font-weight: bold;
    }
    .tag-value {
      img {
        height: 60px;
        padding: 0 20px;
      }
      width: 35%;
    }
  }
}

.el-tag {
  margin: 0 5px;
}
</style>
