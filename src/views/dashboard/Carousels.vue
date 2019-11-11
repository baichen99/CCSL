<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-input
        v-model="params.title"
        prefix-icon="el-icon-search"
        placeholder="请输入标题"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <news-state-selector v-model="params.state" @update="handleSearch" />
      <el-button type="primary" plain @click="handleNew">
        增加
        <i class="el-icon-plus el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="创建时间" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column label="发布人" align="center" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.creator.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="重要性" align="center" width="150px">
          <template slot-scope="{row}">
            <el-rate v-model="row.importance" disabled />
          </template>
        </el-table-column>

        <el-table-column label="状态" align="center" width="100px">
          <template slot-scope="{row}">
            <el-tag v-if="row.state==='published'" type="success">发布</el-tag>
            <el-tag v-if="row.state==='draft'" type="warning">草稿</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="图片标题" align="center" min-width="300px">
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="250px" fixed="right">
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
              type="warning"
              size="mini"
              plain
              @click="handleDraft(row.id)"
            >撤回</el-button>
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
        <carousel-form ref="form" :data="data" />
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
import listMixin from "./listMixin";
import {
  GetCarouselsList,
  CreateCarousel,
  DeleteCarousel,
  UpdateCarousel
} from "@/api/carousel";

export default {
  name: "Carousels",
  components: {
    CarouselForm,
    NewsStateSelector
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        order: "desc",
        title: "",
        state: ""
      }
    };
  },
  methods: {
    getList() {
      this.loading = true;
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
    handleCreate(data) {
      CreateCarousel(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateCarousel(id, updateData)
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
        "此操作将永久删除, 如果想暂时不显示请选择撤回, 是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteCarousel(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handlePublish(id) {
      const updateData = { state: "published" };
      this.handleUpdate(id, updateData);
    },
    handleDraft(id) {
      const updateData = { state: "draft" };
      this.handleUpdate(id, updateData);
    }
  }
};
</script>
