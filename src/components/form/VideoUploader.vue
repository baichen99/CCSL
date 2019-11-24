<template>
  <el-upload
    ref="uploader"
    class="upload-container"
    drag
    accept="video/mp4"
    :limit="1"
    :headers="{
      Authorization: `Bearer ${$store.getters.token}`
    }"
    :action="'/api/files?dir='+dir"
    :on-success="onUploadSuccess"
    :on-remove="clearFile"
    :show-file-list="false"
  >
    <div v-if="fileUrl === ''" class="el-upload__text">
      <i class="el-icon-upload" />
      <div class="upload-text">
        {{ $t("drag") }}
        <em>{{ $t("click") }}</em>
        <br />
        {{ $t("size") }}
      </div>
    </div>
    <div v-else class="video-container">
      <video-player :src="fileUrl" />
    </div>
  </el-upload>
</template>

<i18n>
{
  "zh-CN": {
    "drag": "将文件拖到此处，或",
    "click": "点击上传",
    "size": "大小不超过5Mb"
  },
  "en-US": {
    "drag": "Drag file to here, or",
    "click": "click to upload",
    "size": "Maximal file size 5Mb"
  }
}
</i18n>

<script>
import VideoPlayer from "@/components/video/VideoPlayer";
import { mapGetters } from "vuex";
export default {
  name: "VideoUploader",
  components: {
    VideoPlayer
  },
  model: {
    prop: "src",
    event: "update"
  },
  props: {
    src: {
      type: String,
      default: () => ""
    },
    dir: {
      type: String,
      required: true
    }
  },
  computed: {
    ...mapGetters(["settings"]),
    fileUrl: {
      get() {
        return this.src;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  beforeDestroy() {
    this.clearFile();
  },
  methods: {
    onUploadSuccess(res) {
      const fileUrl = res.data;
      this.fileUrl = fileUrl;
      this.$refs.uploader.clearFiles();
    },
    clearFile() {
      this.fileUrl = "";
      this.$refs.uploader.clearFiles();
    }
  }
};
</script>

<style lang="scss" scoped>
.upload-text {
  margin: 30px;
}

.video-container {
  margin: 20px;
}
</style>