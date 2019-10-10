<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        placeholder="请输入名称"
        clearable
        @clear="handleSearch"
      />
      <el-button type="primary" plain @click="handleSearch">查找</el-button>
      <el-button type="primary" plain @click="handleNew">增加</el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="创建时间" align="center">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column label="上次更新" align="center">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.updatedAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column label="手形名称" align="center">
          <template slot-scope="{row}">
            <el-tag>{{ row.name }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="手形图片" align="center">
          <template slot-scope="{row}">
            <img :src="'https://ccsl.shu.edu.cn/public/'+row.image" alt="sign" style="height:50px" />
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center">
          <template slot-scope="{row}">
            <el-button type="primary" size="mini" plain @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" size="mini" plain @click="handleDelete(row.id)">删除</el-button>
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
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <sign-form ref="form" :data="data" :mode="mode" />
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
import listMixin from "./listMixin";
import { GetSignsList, DeleteSign, CreateSign, UpdateSign } from "@/api/signs";

export default {
  name: "Signs",
  components: {
    SignForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        title: "",
        state: ""
      }
    };
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
    handleCreate(data) {
      CreateSign(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateSign(id, updateData)
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
        "此操作会删除所有视频中含有该手形的标注，此操作将永久删除, 是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteSign(id).then(() => {
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

