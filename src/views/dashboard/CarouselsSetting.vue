<template>
  <div class="app-container flex-row">
    <el-button
      type="primary"
      @click="create"
    >上传新图片</el-button>
    <el-dialog
      :visible.sync="show"
      center
      @closed="onDialogClose"
    >
      <el-upload
        ref="uploader"
        drag
        :limit="1"
        :headers="{
          Authorization: `Bearer ${$store.getters.token}`
        }"
        action="/api/files"
        :on-success="onUploadSuccess"
        :on-remove="onImageRemove"
      >
        <i class="el-icon-upload" />
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em><br>只能上传jpg/png文件，且大小不超过5Mb<br>建议图片尺寸为16:9，否则会拉伸变形</div>
      </el-upload>

      <el-input
        v-model="title"
        placeholder="请输入滚动图片标题"
      />

      <span slot="footer">
        <el-button @click="show=false">取 消</el-button>
        <el-button
          type="primary"
          @click="save"
        >保 存</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { CreateCarousel } from "@/api/carousel";
export default {
  name: "CarouselsSetting",
  data() {
    return {
      show: false,
      image: "",
      title: "",
      carousels: []
    };
  },
  methods: {
    save() {
      const data = { image: this.image, title: this.title };
      CreateCarousel(data).then(() => {
        this.$message({
          message: "上传成功",
          type: "success"
        });
      });
      this.show = false;
    },
    create() {
      this.show = true;
    },
    onDialogClose() {
      this.image = "";
      this.title = "";
      this.$refs.uploader.clearFiles();
    },
    onUploadSuccess(data) {
      const filename = data.data;
      this.image = filename;
    },
    onImageRemove() {
      this.image = "";
    }
  }
};
</script>

<style lang="scss" scoped>
.app-container {
  .el-button {
    width: 200px;
  }
  .uploader {
    .el-upload {
      width: 100%;
    }
  }
}
</style>
