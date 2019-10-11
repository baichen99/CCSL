<template>
  <div class="app-container">
    <div class="table-toolbar"></div>
  </div>
</template>

<script>
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import listMixin from "./listMixin";
import {
  GetLexicalVideosList,
  CreateLexicalVideo,
  UpdateLexicalVideo,
  DeleteLexicalVideo
} from "@/api/videos";
export default {
  name: "LexicalVideos",
  components: { LexicalVideoForm },
  mixins: [listMixin],
  data() {
    return {
      params: {
        performerID: "",
        lexicalWordID: "",
        videoPath: "",
        constructType: "",
        constructWords: [],
        leftSigns: [],
        rightSigns: []
      }
    };
  },
  methods: {
    getList() {
      this.loading = true;
      GetLexicalVideosList(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleCreate(data) {
      CreateLexicalVideo(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateLexicalVideo(id, updateData)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm(
        "此操作将永久删除该视频以及其相关的标注信息，是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteLexicalVideo(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    }
  }
};
</script>

<style lang="scss" scoped>
</style>