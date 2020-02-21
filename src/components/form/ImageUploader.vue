<template>
  <el-upload
    ref="uploader"
    class="upload-container"
    drag
    :multiple="false"
    :accept="acceptType"
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
    <div v-if="isEmpty">
      <i class="el-icon-upload" />
      <div class="upload-text">
        {{ $t("drag") }}
        <em>{{ $t("click") }}</em>
        <br />
        {{ $t("size", {size}) }}
      </div>
    </div>
    <div v-else class="uploaded-image">
      <img :src="settings.publicURL + fileUrl" alt="image" />
    </div>
  </el-upload>
</template>

<i18n>
{
  "zh-CN": {
    "drag": "将文件拖到此处，或",
    "click": "点击上传",
    "size": "文件大小不超过{size}Mb",
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
import { mapGetters } from "vuex";
export default {
  name: "ImageUploader",
  model: {
    prop: "url",
    event: "update"
  },
  props: {
    url: {
      type: String,
      default: () => ""
    },
    dir: {
      type: String,
      required: true
    },
    // Allow all images or svg image
    type: {
      type: String,
      default: "all",
      validator: function(value) {
        return ["all", "svg"].indexOf(value) !== -1;
      }
    },
    // Max size, default 3MB
    size: {
      type: Number,
      default: 3
    }
  },
  computed: {
    ...mapGetters(["settings"]),
    fileUrl: {
      get() {
        return this.url;
      },
      set(val) {
        this.$emit("update", val);
      }
    },
    acceptType() {
      if (this.type === "svg") {
        return "image/svg+xml";
      } else {
        return "image/*";
      }
    },
    isEmpty() {
      return this.fileUrl === "";
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
.upload-container {
  .upload-text {
    margin: 30px;
  }
  .uploaded-image {
    padding: 30px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  img {
    max-height: 100%;
    max-width: 100%;
    border-radius: 4px;
  }
}
</style>