<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        placeholder="请输入姓名"
        clearable
        @clear="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
      <gender-selector v-model="params.gender" @update="handleSearch" />
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

        <el-table-column label="被试姓名" align="center">
          <template slot-scope="{row}">
            <span>{{ row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="被试性别" align="center">
          <template slot-scope="{row}">
            <span>{{ genderTypes[row.gender].name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="被试所在地区" align="center">
          <template slot-scope="{row}">
            <span>{{ row.region.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="180px" fixed="right">
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
        <performer-form ref="form" :data="data" :mode="mode" />
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
import PerformerForm from "@/views/dashboard/form/PerformerForm";
import GenderSelector from "@/components/form/GenderSelector";
import CitySelector from "@/components/form/CitySelector";
import listMixin from "./listMixin";
import {
  GetPerformersList,
  CreatePerformer,
  UpdatePerformer,
  DeletePerformer
} from "@/api/performers";
export default {
  name: "Performers",
  components: {
    PerformerForm,
    GenderSelector,
    CitySelector
  },
  mixins: [listMixin],
  data() {
    return {
      removeProperties: ["region"],
      params: {
        regionID: undefined,
        name: "",
        gender: ""
      }
    };
  },
  computed: {
    ...mapGetters(["genderTypes"])
  },
  methods: {
    getList() {
      this.loading = true;
      GetPerformersList(this.params)
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
      CreatePerformer(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdatePerformer(id, updateData)
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
        "此操作会删除所有视频中含有该被试信息的标注，此操作将永久删除, 是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeletePerformer(id).then(() => {
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