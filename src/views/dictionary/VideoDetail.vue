<template>
  <div>
    <video-player
      :options="{
        sources: [{ type: 'video/mp4', src: 'public/videos/'+video.videoPath }]
      }"
      width="100%"
      class="vjs-big-play-centered"
    />
    <div class="tags-container">
      <div class="tags">
        <span class="tag-label">地区</span>
        <span class="tag-value">{{ video.region }}</span>
        <span class="tag-label">性别</span>
        <span class="tag-value">{{ video.gender }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">汉语转写</span>
        <span class="tag-value">{{ video.chinese }}</span>
        <span class="tag-label">英语转写</span>
        <span class="tag-value">{{ video.english }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">汉语音序</span>
        <span class="tag-value">{{ video.initial }}</span>
        <span class="tag-label">词性</span>
        <span class="tag-value">{{ video.type }}</span>
      </div>
      <div class="tags">
        <span class="tag-label">构词方式</span>
        <span class="tag-value">{{ video.constructType || "暂无数据" }}</span>
        <span class="tag-label">复合词构词词语</span>
        <span class="tag-value">{{ video.constructWords || "暂无数据" }}</span>
      </div>
      <div
        v-if="video.leftSign"
        class="tags"
      >
        <span class="tag-label">左手手型</span>
        <span class="tag-value">
          <span
            v-for="item in video.leftSign.split(',')"
            :key="item"
          >
            <el-tooltip
              effect="dark"
              :content="item"
              placement="top"
            >
              <img
                style="width:100px"
                :src="'sign/'+item+'.jpg'"
                :alt="item"
              >
            </el-tooltip>
          </span>
        </span>
      </div>
      <div
        v-if="video.rightSign"
        class="tags"
      >
        <span class="tag-label">右手手型</span>
        <span class="tag-value">
          <span
            v-for="item in video.rightSign.split(',')"
            :key="item"
          >
            <el-tooltip
              effect="dark"
              :content="item"
              placement="top"
            >
              <img
                style="width:100px"
                :src="'sign/'+item+'.jpg'"
                :alt="item"
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
.tags {
  padding: 10px 20px;
  display: flex;
  align-items: center;
  justify-content: space-evenly;
}
.tag-label {
  font-weight: bold;
}
</style>
