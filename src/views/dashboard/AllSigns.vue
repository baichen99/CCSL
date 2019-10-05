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
      background
      layout="prev, pager, next"
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
          :data="data"
          :mode="mode"
        />
        <div class="form-drawer__footer">
          <el-button @click="handleClose">取 消</el-button>
          <el-button
            v-if="checkDiff"
            type="primary"
            :loading="loading"
            @click="handleSave"
          >{{ loading ? '保存中 ...' : '确 定' }}</el-button>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script>
import SignForm from "@/views/dashboard/forms/SignForm";
import { getSigns } from "@/api/signs";
export default {
  name: "AllSigns",
  components: {
    SignForm
  },
  data() {
    return {
      show: false,
      mode: "",
      data: {},
      params: {
        name: "",
        limit: 8,
        page: 1
      },
      total: 0,
      loading: false,
      list: []
    };
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
      getSigns(this.params)
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
      this.mode = "";
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
      this.data = data;
    },
    handleDelete(id) {
      console.log(id);
    },
    handleClose() {
      this.show = false;
      this.clearData();
    },
    handleSave() {
      if (this.mode === "edit") {
        console.log("UPDATE DATA");
      } else if (this.mode === "create") {
        console.log("CREATE DATA");
      } else {
        console.error("NO AVALIABLE OPERATION");
      }
    },
    checkDiff() {
      return true;
    }
  }
};
</script>

<style lang="scss" scoped>
.el-input {
  width: 200px;
}
</style>