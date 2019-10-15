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
      <city-selector v-model="params.regionID" @update="handleSearch" />
      <word-construct-selector v-model="params.constructType" @update="handleSearch" />
      <el-button type="primary" plain @click="handleSearch">查找</el-button>
      <el-button type="primary" plain @click="handleNew">增加</el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column label="音序" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lexicalWord.initial }}</span>
          </template>
        </el-table-column>

        <el-table-column label="中文转写" align="center" min-width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lexicalWord.chinese }}</span>
          </template>
        </el-table-column>

        <el-table-column label="构词" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.constructType }}</span>
          </template>
        </el-table-column>

        <el-table-column label="复合词构词词语" align="center" min-width="150px">
          <template slot-scope="{row}">
            <div v-if="row.constructType==='单纯词'">
              <el-tag type="danger">不可用</el-tag>
            </div>
            <div v-else-if="row.constructWords">
              <el-tag
                v-for="value in row.constructWords"
                :key="value"
                class="construct-tags"
              >{{ value }}</el-tag>
            </div>
            <div v-else>
              <el-tag type="info">暂无数据</el-tag>
            </div>
          </template>
        </el-table-column>

        <el-table-column label="被试姓名" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.performer.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="被试地区" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.performer.region.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="被试性别" align="center" width="100px">
          <template slot-scope="{row}">
            <span>{{ genderTypes[row.performer.gender].name }}</span>
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
        <lexical-video-form ref="form" :data="data" :mode="mode" />
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
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import CitySelector from "@/components/form/CitySelector.vue";
import WordConstructSelector from "@/components/form/WordConstructSelector";
import listMixin from "./listMixin";
import {
  GetLexicalVideosList,
  CreateLexicalVideo,
  UpdateLexicalVideo,
  DeleteLexicalVideo
} from "@/api/videos";
export default {
  name: "LexicalVideos",
  components: {
    LexicalVideoForm,
    CitySelector,
    WordConstructSelector
  },
  mixins: [listMixin],
  data() {
    return {
      removeProperties: ["performer", "lexicalWord", "leftSigns", "rightSigns"],
      params: {
        chinese: "",
        regionID: undefined,
        // lexicalWordID: "",
        constructType: "",
        orderBy: "performers.gender"
      }
    };
  },
  computed: {
    ...mapGetters(["genderTypes"])
  },
  methods: {
    getList() {
      this.loading = true;
      GetLexicalVideosList(this.params)
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
      CreateLexicalVideo(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateLexicalVideo(id, updateData)
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
        "此操作将永久删除该视频以及其相关的标注信息，是否继续?",
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        }
      )
        .then(() => {
          DeleteLexicalVideo(id).then(() => {
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

<style lang="scss" scoped>
.construct-tags {
  margin: 4px;
}
</style>