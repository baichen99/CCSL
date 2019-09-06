<template>
  <el-upload
    class="upload-container"
    drag
    accept="image/*"
    :limit="1"
    :headers="{
      Authorization: `Bearer ${$store.getters.token}`
    }"
    action="/api/files"
    :on-success="onUploadSuccess"
    :on-remove="clearFile"
  >
    <i class="el-icon-upload" />
    <div
      v-if="imageUrl === ''"
      class="el-upload__text"
    >
      将文件拖到此处，或<em>点击上传</em><br>只能上传jpg/png文件，且大小不超过5Mb<br>建议图片尺寸为16:9，否则会拉伸变形
    </div>
    <div v-else>
      <img
        :src="'/public/files/'+imageUrl"
        alt="image"
      >
    </div>
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
</style>