<template>
  <div>
    <div class="video">
      <video-player
        :options="{
          sources: [{ type: 'video/mp4', src: 'https://ccsl.shu.edu.cn/public/videos/'+video.videoPath }]
        }"
        width="100%"
        class="vjs-big-play-centered"
        @contextmenu.native.prevent=""
      />
    </div>

    <div class="tags-container">
      <div class="tags">
        <span class="tag-label">地区</span>
        <span class="tag-value">{{ video.performer.region.name }}</span>
        <span class="tag-label">汉语转写</span>
        <span class="tag-value">{{ video.word.chinese }}</span>

      </div>
      <div class="tags">
        <span class="tag-label">词性</span>
        <span class="tag-value">{{ video.word.pos }}</span>
        <span class="tag-label">英语转写</span>
        <span class="tag-value">{{ video.word.english }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">构词方式</span>
        <span class="tag-value">{{ video.constructType || "暂无数据" }}</span>
        <span class="tag-label">构词词根</span>
        <span class="tag-value">{{ video.constructWords || "暂无数据" }}</span>
      </div>
      <div
        v-if="video.leftSigns.length !== 0"
        class="tags"
      >
        <span class="tag-label">左手手型</span>
        <span
          class="tag-value"
          style="width:50%"
        >
          <span
            v-for="item in video.leftSigns"
            :key="item.id"
          >
            <el-tooltip
              effect="dark"
              :content="item.name"
              placement="top"
            >
              <img
                style="width:100px"
                :src="'https://ccsl.shu.edu.cn/public/signs/'+item.image"
                :alt="item.name"
              >
            </el-tooltip>
          </span>
        </span>
      </div>
      <div
        v-if="video.rightSigns.length !== 0"
        class="tags"
      >
        <span class="tag-label">右手手型</span>
        <span
          class="tag-value"
          style="width:50%"
        >
          <span
            v-for="item in video.rightSigns"
            :key="item.id"
          >
            <el-tooltip
              effect="dark"
              :content="item.name"
              placement="top"
            >
              <img
                style="width:100px"
                :src="'https://ccsl.shu.edu.cn/public/signs/'+item.image"
                :alt="item.name"
              >
            </el-tooltip>
          </span>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "VideoDetail",
  props: {
    video: {
      type: Object,
      required: true,
      default: () => ({})
    }
  }
};
</script>

<style lang="scss" scoped>
.video {
  margin: 5px;
  border-radius: 5px;
  overflow: hidden;
}

.tags {
  text-align: center;
  padding: 10px 20px;
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
      height: 120px;
      padding: 10px;
    }
    width: 25%;
  }
}
</style>
