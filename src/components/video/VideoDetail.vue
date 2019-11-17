<template>
  <div>
    <div class="video-container">
      <video-player
        :options="{ 
          sources: [{ 
            type: 'video/mp4',
            src: settings.publicURL + video.videoPath 
          }] 
        }"
        width="100%"
        class="vjs-big-play-centered"
        @contextmenu.native.prevent
      />
    </div>

    <div class="tags-container">
      <div class="tags">
        <span class="tag-label">{{ $t("Region") }}</span>
        <span class="tag-value">{{ video.performer.region.name }}</span>
        <span class="tag-label">{{ $t("Chinese") }}</span>
        <span class="tag-value" v-html="$options.filters.addNumberSup(video.lexicalWord.chinese) "></span>
      </div>
      <div class="tags">
        <span class="tag-label">{{ $t("PoS") }}</span>
        <span class="tag-value">{{ video.lexicalWord.pos }}</span>
        <span class="tag-label">{{ $t("English") }}</span>
        <span class="tag-value">{{ video.lexicalWord.english }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">{{ $t("WordFormation") }}</span>
        <span
          v-if="video.constructType"
          class="tag-value"
        >{{ $t(constructTypes[video.constructType].name) }}</span>
        <span v-else class="tag-value">
          <el-tag type="info">{{ $t("NoData") }}</el-tag>
        </span>
        <span class="tag-label">{{ $t("CompoundLexemes") }}</span>
        <span v-if="video.constructType==='simple'" class="tag-value">
          <el-tag type="danger">{{ $t("Unavailable") }}</el-tag>
        </span>
        <span v-else-if="video.constructWords.length > 0" class="tag-value">
          <el-tag
            v-for="(item,index) in video.constructWords"
            :key="index"
            class="tag-words"
          >{{ item }}</el-tag>
        </span>
        <span v-else class="tag-value">
          <el-tag type="info">{{ $t("NoData") }}</el-tag>
        </span>
      </div>
      <div v-if="video.leftSigns.length !== 0" class="tags">
        <span class="tag-label">{{ $t("LeftSign") }}</span>
        <span class="tag-value" style="width:50%">
          <span v-for="item in video.leftSigns" :key="item.id">
            <el-tooltip effect="dark" :content="item.name" placement="top">
              <img
                style="width:150px"
                :src="settings.publicURL + item.image"
                :alt="item.name"
              />
            </el-tooltip>
          </span>
        </span>
      </div>
      <div v-if="video.rightSigns.length !== 0" class="tags">
        <span class="tag-label">{{ $t("RightSign") }}</span>
        <span class="tag-value" style="width:50%">
          <span v-for="item in video.rightSigns" :key="item.id">
            <el-tooltip effect="dark" :content="item.name" placement="top">
              <img
                style="width:150px"
                :src="settings.publicURL + item.image"
                :alt="item.name"
              />
            </el-tooltip>
          </span>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "VideoDetail",
  props: {
    video: {
      type: Object,
      required: true,
      default: () => ({})
    }
  },
  computed: {
    ...mapGetters(["constructTypes", "settings"])
  }
};
</script>

<style lang="scss" scoped>
.tags {
  text-align: center;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  line-height: 30px;
  .tag-label {
    width: 25%;
    font-weight: bold;
  }
  .tag-value {
    img {
      height: 100px;
      padding: 10px;
    }
    width: 25%;
  }
  .tag-words {
    margin: 0 10px;
  }
}
</style>
