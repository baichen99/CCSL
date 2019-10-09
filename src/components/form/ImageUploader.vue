<template>
  <el-upload
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
  >
    <div v-if="imageUrl === ''" class="el-upload__text">
      <i class="el-icon-upload" />
      <div>
        将文件拖到此处，或
        <em>点击上传</em>
        <br />大小不超过5Mb
      </div>
    </div>

    <img v-else :src="'https://ccsl.shu.edu.cn/public/'+imageUrl" alt="image" />
  </el-upload>
</template>

<script>
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
    },
    clearFile() {
      this.imageUrl = "";
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