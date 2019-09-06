<template>
  <div class="app-container flex-column">
    <el-button
      class="new"
      type="primary"
      @click="create"
    >增加</el-button>

    <el-dialog
      :visible.sync="show"
      center
      @closed="onDialogClose"
    >
      <image-uploader
        v-model="image"
      />
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

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      stripe
      fit
      highlight-current-row
      style="width: 100%"
    >

      <el-table-column
        width="180px"
        align="center"
        label="发布时间"
      >
        <template slot-scope="{row}">
          <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="标题"
        align="center"
      >
        <template slot-scope="{row}">
          <template v-if="row.edit">
            <el-input
              v-model="row.title"
              class="edit-input"
              size="small"
            />
            <el-button
              class="cancel-btn"
              size="small"
              icon="el-icon-refresh"
              type="warning"
              @click="cancelEdit(row)"
            >
              取消
            </el-button>
          </template>
          <span v-else>{{ row.title }}</span>
        </template>
      </el-table-column>

      <el-table-column
        align="center"
        label="操作"
        width="200"
        fixed="right"
      >
        <template slot-scope="{row}">
          <el-button
            v-if="row.edit"
            type="success"
            size="small"
            icon="el-icon-edit-outline"
            @click="confirmEdit(row)"
          >
            保存
          </el-button>
          <el-button
            v-else
            type="primary"
            size="small"
            icon="el-icon-edit"
            @click="row.edit=!row.edit"
          >
            编辑
          </el-button>
          <el-button
            type="danger"
            size="small"
            icon="el-icon-delete"
            @click="confirmDelete(row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {
  GetCarouselList,
  CreateCarousel,
  DeleteCarousel
} from "@/api/carousel";
import ImageUploader from "@/components/ImageUploader";
export default {
  name: "CarouselsSetting",
  components: {
    ImageUploader
  },
  data() {
    return {
      show: false,
      image: "",
      title: "",
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10
      }
    };
  },
  created() {
    this.getList();
  },
  methods: {
    save() {
      const data = { image: this.image, title: this.title };
      CreateCarousel(data).then(() => {
        this.$message({
          message: "上传成功",
          type: "success"
        });
        this.getList();
      });
      this.show = false;
    },
    create() {
      this.show = true;
    },
    onDialogClose() {
      this.image = "";
      this.title = "";
    },
    async getList() {
      this.listLoading = true;
      const { data } = await GetCarouselList(this.listQuery);
      const items = data;
      this.list = items.map(v => {
        this.$set(v, "edit", false);
        v.originalTitle = v.title;
        return v;
      });
      this.listLoading = false;
    },
    cancelEdit(row) {
      row.title = row.originalTitle;
      row.edit = false;
      this.$message({
        message: "已撤销修改",
        type: "warning"
      });
    },
    confirmEdit(row) {
      row.edit = false;
      row.originalTitle = row.title;
      const id = row.id;
      console.log(id);
      this.$message({
        message: "已保存修改",
        type: "success"
      });
    },
    confirmDelete(row) {
      const id = row.id;
      DeleteCarousel(id).then(() => {
        this.$message({
          message: "删除成功",
          type: "success"
        });
        this.getList();
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.app-container {
  .new {
    margin: 10px 0;
  }

  .el-table {
    .edit-input {
      padding-right: 100px;
    }
    .cancel-btn {
      position: absolute;
      right: 15px;
      top: 12px;
    }
  }
}
</style>
