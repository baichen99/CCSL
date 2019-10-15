<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-input
        v-model="params.username"
        prefix-icon="el-icon-search"
        placeholder="请输入账号"
        clearable
        @clear="handleSearch"
      />
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        placeholder="请输入姓名"
        clearable
        @clear="handleSearch"
      />
      <user-type-selector v-model="params.userType" @clear="handleSearch" />
      <el-button type="primary" plain @click="handleSearch">查找</el-button>
      <el-button type="primary" plain @click="handleNew">增加</el-button>
    </div>
    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="创建时间" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column label="上次更新" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.updatedAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column label="用户角色" align="center" width="150px">
          <template slot-scope="{row}">
            <el-tag :type="userTypes[row.userType].color">{{ userTypes[row.userType].name }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="用户账号" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.username }}</span>
          </template>
        </el-table-column>

        <el-table-column label="用户姓名" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="250px" fixed="right">
          <template slot-scope="{row}">
            <el-button type="primary" size="mini" plain @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" size="mini" plain @click="handleDelete(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-pagination
      background
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      :page-size.sync="params.limit"
      :current-page.sync="params.page"
      :hide-on-single-page="true"
    />

    <el-drawer
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <user-form ref="form" :data="data" :mode="mode" />
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
import { mapGetters } from "vuex";
import UserForm from "@/views/dashboard/form/UserForm";
import UserTypeSelector from "@/components/form/UserTypeSelector";
import listMixin from "./listMixin";
import { GetUsersList, CreateUser, UpdateUser, DeleteUser } from "@/api/users";
export default {
  name: "Users",
  components: {
    UserForm,
    UserTypeSelector
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        username: "",
        name: "",
        userType: ""
      }
    };
  },
  computed: { ...mapGetters(["userTypes"]) },
  methods: {
    getList() {
      this.loading = true;
      GetUsersList(this.params)
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
      CreateUser(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateUser(id, updateData)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm("此操作将永久删除, 是否继续?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "error"
      })
        .then(() => {
          DeleteUser(id).then(() => {
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
