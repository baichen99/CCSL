<template>
  <div class="app-container flex-column">

    <div class="table-toolbar">
      <el-input
        v-model="params.title"
        prefix-icon="el-icon-search"
        placeholder="请输入图片标题以查找"
        clearable
        @clear="handleSearch"
      />
      <news-state-selector
        v-model="params.state"
        @clear="handleSearch"
      />
      <el-button
        type="primary"
        plain
        @click="handleSearch"
      >查找</el-button>
      <el-button
        type="primary"
        plain
        @click="handleCreate"
      >增加</el-button>
    </div>

    <div class="table-content">
      <el-table
        v-loading="loading"
        :data="list"
        stripe
        border
      >
        <el-table-column
          label="创建时间"
          align="center"
          width="180px"
        >
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="上次更新"
          align="center"
          width="180px"
        >
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.updatedAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="重要性"
          align="center"
          width="150px"
        >
          <template slot-scope="{row}">
            <el-rate
              v-model="row.importance"
              disabled
            />
          </template>
        </el-table-column>

        <el-table-column
          label="状态"
          align="center"
          width="100px"
        >
          <template slot-scope="{row}">
            <el-tag
              v-if="row.state==='published'"
              type="success"
            >发布</el-tag>
            <el-tag
              v-if="row.state==='draft'"
              type="info"
            >草稿</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          label="图片标题"
          align="center"
        >
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="操作"
          align="center"
          width="250px"
          fixed="right"
        >
          <template slot-scope="{row}">
            <el-button
              v-if="row.state==='draft'"
              type="success"
              size="mini"
              plain
              @click="handlePublish(row.id)"
            >发布</el-button>
            <el-button
              v-if="row.state==='published'"
              type="info"
              size="mini"
              plain
              @click="handleDraft(row.id)"
            >撤回</el-button>
            <el-button
              type="primary"
              size="mini"
              plain
              @click="handleEdit(row)"
            >编辑</el-button>
            <el-button
              type="danger"
              size="mini"
              plain
              @click="handleDelete(row.id)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-pagination
      v-if="total>params.limit"
      background
      layout="total,prev, pager, next"
      :total="total"
      :page-size.sync="params.limit"
      :current-page.sync="params.page"
    />

    <el-drawer
      ref="drawer"
      size="40%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <carousel-form
          :data="originalData"
          :mode="mode"
        />
        <div class="form-drawer__footer">
          <el-button @click="handleClose">取 消</el-button>
          <el-button
            v-if="checkDiff"
            type="primary"
            :loading="loading"
            @click="handleSave"
          >{{ loading ? '保存中 ...' : '保 存' }}</el-button>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script>
import CarouselForm from "@/views/dashboard/form/CarouselForm";
import NewsStateSelector from "@/components/form/NewsStateSelector";
import {
  GetCarouselsList,
  CreateCarousel,
  DeleteCarousel,
  UpdateCarousel
} from "@/api/carousel";

export default {
  name: "AllCarousels",
  components: {
    CarouselForm,
    NewsStateSelector
  },
  data() {
    return {
      show: false,
      mode: "",
      originalData: {},
      data: {},
      params: {
        limit: 8,
        page: 1,
        order: "desc",
        title: "",
        state: ""
      },
      total: 0,
      loading: false,
      list: []
    };
  },
  computed: {
    checkDiff() {
      if (this.mode === "create") {
        return true;
      } else {
        let diffFound = false;
        for (let key in this.data) {
          if (this.originalData[key] !== this.data[key]) {
            diffFound = true;
          }
        }
        return diffFound;
      }
    }
  },
  watch: {
    "params.page"() {
      this.getList();
    }
  },
  created() {
    this.getList();
  },
  methods: {
    getList() {
      GetCarouselsList(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    },
    clearData() {
      this.data = {};
      this.originalData = {};
      this.mode = "";
    },
    checkData() {
      if (!this.originalData.title) {
        this.$message({
          type: "warning",
          message: "请填写完整信息"
        });
        return false;
      } else {
        this.originalData.title = this.originalData.title.trim();
        return true;
      }
    },
    handleSearch() {
      this.params.page = 1;
      this.getList();
    },
    handleCreate() {
      this.mode = "create";
      this.show = true;
    },
    handleEdit(data) {
      this.mode = "edit";
      this.show = true;
      // Use JSON.parse(JSON.stringify(data)) to deep copy a object
      this.originalData = JSON.parse(JSON.stringify(data));
      this.data = JSON.parse(JSON.stringify(data));
    },
    handleClose() {
      this.show = false;
      this.clearData();
    },
    handleDelete(id) {
      this.$confirm("此操作将永久删除, 是否继续?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "error"
      })
        .then(() => {
          DeleteCarousel(id).then(() => {
            this.$message({
              type: "success",
              message: "操作成功"
            });
            this.handleSearch();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消操作"
          });
        });
    },
    handleSave() {
      this.loading = true;
      if (this.mode === "edit") {
        let updateData = {};
        for (let key in this.data) {
          if (this.originalData[key] !== this.data[key]) {
            updateData[key] = this.originalData[key];
          }
        }
        this.handleUpdate(this.data.id, updateData);
      } else if (this.mode === "create") {
        if (this.checkData()) {
          CreateCarousel(this.originalData)
            .then(() => {
              this.$message({
                type: "success",
                message: "操作成功"
              });
              this.handleSearch();
              this.loading = false;
            })
            .catch(() => {
              this.loading = false;
            });
        } else {
          this.loading = false;
          return;
        }
      } else {
        console.error("NO AVALIABLE OPERATION");
      }
      this.handleClose();
    },
    handlePublish(id) {
      const updateData = { state: "published" };
      this.handleUpdate(id, updateData);
    },
    handleDraft(id) {
      const updateData = { state: "draft" };
      this.handleUpdate(id, updateData);
    },
    handleUpdate(id, updateData) {
      UpdateCarousel(id, updateData)
        .then(() => {
          this.$message({
            type: "success",
            message: "操作成功"
          });
          this.handleSearch();
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    }
  }
};
</script>
