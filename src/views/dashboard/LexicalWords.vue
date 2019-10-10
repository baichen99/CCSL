<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        placeholder="请输入中文"
        clearable
        @clear="handleSearch"
      />
      <el-input
        v-model="params.english"
        prefix-icon="el-icon-search"
        placeholder="请输入英文"
        clearable
        @clear="handleSearch"
      />
      <word-pos-selector v-model="params.pos" @clear="handleSearch" />
      <word-initial-selector v-model="params.initial" @clear="handleSearch" />
      <el-button type="primary" plain @click="handleSearch">查找</el-button>
      <el-button type="primary" plain @click="handleNew">增加</el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="音序" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.initial }}</span>
          </template>
        </el-table-column>

        <el-table-column label="词性" align="center" min-width="100px">
          <template slot-scope="{row}">
            <span>{{ row.pos }}</span>
          </template>
        </el-table-column>

        <el-table-column label="汉语转写" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.chinese }}</span>
          </template>
        </el-table-column>

        <el-table-column label="英语转写" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.english }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="200px" fixed="right">
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
      size="30%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <lexical-word-form ref="form" :data="data" :mode="mode" />
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
import LexicalWordForm from "@/views/dashboard/form/LexicalWordForm";
import WordPosSelector from "@/components/form/WordPosSelector";
import WordInitialSelector from "@/components/form/WordInitialSelector";
import listMixin from "./listMixin";
import {
  GetLexicalWordsList,
  CreateLexicalWord,
  UpdateLexicalWord,
  DeleteLexicalWord
} from "@/api/words";

export default {
  name: "LexicalWords",
  components: {
    WordPosSelector,
    WordInitialSelector,
    LexicalWordForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        limit: 10,
        initial: "",
        chinese: "",
        english: "",
        pos: ""
      }
    };
  },
  methods: {
    getList() {
      this.loading = true;
      GetLexicalWordsList(this.params)
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
      CreateLexicalWord(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateLexicalWord(id, updateData)
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
        "此操作将永久删除所有视频中该词语的标注信息，是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteLexicalWord(id).then(() => {
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