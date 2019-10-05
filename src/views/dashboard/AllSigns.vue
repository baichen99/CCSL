<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        placeholder="请输入名称以查找"
        clearable
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
        >
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="上次更新"
          align="center"
        >
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.updatedAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column
          label="手形名称"
          align="center"
        >
          <template slot-scope="{row}">
            <el-tag>{{ row.name }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          label="手形图片"
          align="center"
        >
          <template slot-scope="{row}">
            <img
              :src="'https://ccsl.shu.edu.cn/public/'+row.image"
              alt="sign"
              style="height:50px"
            >
          </template>
        </el-table-column>

        <el-table-column
          label="操作"
          align="center"
        >
          <template slot-scope="{row}">
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
        <sign-form
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
import SignForm from "@/views/dashboard/form/SignForm";
import { GetSignsList, DeleteSign, CreateSign, UpdateSign } from "@/api/signs";
export default {
  name: "AllSigns",
  components: {
    SignForm
  },
  data() {
    return {
      show: false,
      mode: "",
      originalData: {},
      data: {},
      params: {
        title: "",
        state: "",
        limit: 8,
        page: 1
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
      this.loading = true;
      GetSignsList(this.params)
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
      if (!this.originalData.name || !this.originalData.image) {
        this.$message({
          type: "warning",
          message: "请填写完整信息"
        });
        return false;
      } else {
        this.originalData.name = this.originalData.name.trim();
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
      this.$confirm(
        "删除该手形会删除所有视频中含有该手形的标注，此操作将永久删除, 是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteSign(id).then(() => {
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
        UpdateSign(this.data.id, updateData)
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
      } else if (this.mode === "create") {
        if (this.checkData()) {
          CreateSign(this.originalData)
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
    }
  }
};
</script>

