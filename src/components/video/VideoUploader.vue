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
    :on-success="onSuccess"
    :on-error="onError"
    :on-remove="clearFile"
    :before-upload="beforeUpload"
    :show-file-list="false"
    :with-credentials="true"
  >
    <div v-if="fileUrl===''" class="el-upload__text">
      <i class="el-icon-upload" />
      <div class="upload-text">
        {{ $t("drag") }}
        <em>{{ $t("click") }}</em>
        <br />
        {{ $t("size", { size }) }}
      </div>
    </div>
    <div v-else class="video-container" @click.stop>
      <video-player :src="fileUrl" />
    </div>
  </el-upload>
</template>

<i18n>
{
  "zh-CN": {
    "drag": "将文件拖到此处，或",
    "click": "点击上传",
    "size": "大小不超过{size}Mb",
    "sizeError": "文件太大，请压缩后再上传"
  },
  "en-US": {
    "drag": "Drag file to here, or",
    "click": "click to upload",
    "size": "Maximal file size {size}Mb",
    "sizeError": "File size exceed limitation, please compress it"
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
    },
    size: {
      type: Number,
      default: 20
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
    onSuccess(res) {
      const fileUrl = res.data;
      this.fileUrl = fileUrl;
      this.$refs.uploader.clearFiles();
    },
    onError(err) {
      const error = JSON.parse(err.message);
      this.$notify.error({
        title: error.message,
        message: error.error
      });
    },
    beforeUpload(file) {
      const lessThanMaxSize = file.size / 1024 / 1024 < this.size;
      if (!lessThanMaxSize) {
        this.$notify.error({
          title: this.$t("sizeError")
        });
      }
      return lessThanMaxSize;
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
  margin: 40px;
}
</style>