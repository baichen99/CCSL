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
      <language-selector v-model="params.language" @update="handleSearch" />
      <news-state-selector v-model="params.state" @update="handleSearch" />
      <news-type-selector v-model="params.type" @update="handleSearch" />
      <news-column-selector v-model="params.column" @update="handleSearch" />
      <el-button type="primary" plain @click="handleNew">
        增加
        <i class="el-icon-plus el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="日期" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.date), 'short') }}</span>
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

        <el-table-column label="新闻栏目" align="center" width="120px">
          <template slot-scope="{row}">
            <span>{{ $t(newsColumns[row.column].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column label="新闻标题" align="center" min-width="300px">
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
        <news-form ref="form" :data="data" :mode="mode" />
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
import NewsForm from "@/views/dashboard/form/NewsForm";
import NewsStateSelector from "@/components/form/NewsStateSelector";
import NewsTypeSelector from "@/components/form/NewsTypeSelector";
import NewsColumnSelector from "@/components/form/NewsColumnSelector";
import LanguageSelector from "@/components/form/LanguageSelector";
import listMixin from "./listMixin";
import { GetNewsList, CreateNews, UpdateNews, DeleteNews } from "@/api/news";

export default {
  name: "News",
  components: {
    NewsForm,
    NewsStateSelector,
    NewsTypeSelector,
    NewsColumnSelector,
    LanguageSelector
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        order: "desc",
        title: "",
        type: "",
        column: "",
        language: "",
        state: ""
      }
    };
  },
  computed: { ...mapGetters(["newsColumns"]) },
  methods: {
    getList() {
      this.loading = true;
      GetNewsList(this.params)
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
      CreateNews(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateNews(id, updateData)
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
          DeleteNews(id).then(() => {
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
