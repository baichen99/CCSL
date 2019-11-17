<template>
  <el-upload
    ref="uploader"
    class="upload-container"
    drag
    :accept="acceptType"
    :limit="1"
    :headers="{
      Authorization: `Bearer ${$store.getters.token}`
    }"
    :action="'/api/files?dir='+dir"
    :on-success="onUploadSuccess"
    :on-remove="clearFile"
    :style="{height}"
    :show-file-list="false"
  >
    <div v-if="imageUrl === ''" class="el-upload__text">
      <i class="el-icon-upload" />
      <div>
        {{ $t("drag") }}
        <em>{{ $t("click") }}</em>
        <br />
        {{ $t("size") }}
      </div>
    </div>
    <img v-else :src="settings.publicURL + imageUrl" alt="image" />
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
    type: {
      type: String,
      default: "all",
      validator: function(value) {
        return ["all", "svg"].indexOf(value) !== -1;
      }
    },
    height: {
      type: String,
      default: "500px"
    }
  },
  computed: {
    ...mapGetters(["settings"]),
    imageUrl: {
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
    }
  },
  beforeDestroy() {
    this.clearFile();
  },
  methods: {
    onUploadSuccess(res) {
      const fileUrl = res.data;
      this.imageUrl = fileUrl;
      this.$refs.uploader.clearFiles();
    },
    clearFile() {
      this.imageUrl = "";
      this.$refs.uploader.clearFiles();
    }
  }
};
</script>

<style lang="scss" scoped>
img {
  max-height: 100%;
  max-width: 100%;
  padding: 5px;
}
</style>