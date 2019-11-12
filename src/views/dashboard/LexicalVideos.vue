<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        :placeholder="$t('chineseTip')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
      <gender-selector v-model="params.gender" @update="handleSearch" />
      <sign-selector v-model="params.signID" @update="handleSearch" />
      <sign-selector v-model="params.leftSignID" orientation="left" @update="handleSearch" />
      <sign-selector v-model="params.rightSignID" orientation="right" @update="handleSearch" />
      <word-construct-selector v-model="params.constructType" @update="handleSearch" />
      <lexemes-input v-model="params.constructWords" @update="handleSearch" @enter="handleSearch" />
      <el-button type="primary" plain @click="handleNew">
        增加
        <i class="el-icon-plus el-icon--right" />
      </el-button>
      <el-button type="primary" plain @click="handleExport">
        导出
        <i class="el-icon-download el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column :label="$t('Initial')" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lexicalWord.initial }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Region')" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.performer.region.name }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Gender')" align="center" width="100px">
          <template slot-scope="{row}">
            <span>{{ $t(genderTypes[row.performer.gender].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Chinese')" align="center" min-width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lexicalWord.chinese }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('WordFormation')" align="center" width="150px">
          <template slot-scope="{row}">
            <span v-if="row.constructType">{{ $t(constructTypes[row.constructType].name) }}</span>
            <span v-else>{{ $t("NoData") }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('CompoundLexemes')" align="center" min-width="150px">
          <template slot-scope="{row}">
            <div v-if="row.constructType==='single'">
              <el-tag type="danger">{{ $t("Unavailable") }}</el-tag>
            </div>
            <div v-else-if="row.constructWords.length > 0">
              <el-tag
                v-for="value in row.constructWords"
                :key="value"
                class="construct-tags"
              >{{ value }}</el-tag>
            </div>
            <div v-else>
              <el-tag type="info">{{ $t("NoData") }}</el-tag>
            </div>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Name')" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.performer.name }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="200px" fixed="right">
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
        <lexical-video-form ref="form" :data="data" />
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

<i18n>
{
  "zh-CN": {
    "chineseTip": "请输入中文"
  },
  "en-US": {
    "chineseTip": "Input Chinese"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import CitySelector from "@/components/form/CitySelector.vue";
import GenderSelector from "@/components/form/GenderSelector";
import WordConstructSelector from "@/components/form/WordConstructSelector";
import SignSelector from "@/components/form/SignSelector.vue";
import LexemesInput from "@/components/form/LexemesInput.vue";
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
    GenderSelector,
    SignSelector,
    WordConstructSelector,
    LexemesInput
  },
  mixins: [listMixin],
  data() {
    return {
      removeProperties: ["performer", "lexicalWord", "leftSigns", "rightSigns"],
      params: {
        chinese: "",
        regionID: undefined,
        gender: "",
        signID: "",
        leftSignID: "",
        rightSignID: "",
        constructType: "",
        constructWords: "",
        orderBy: "performers.gender"
      }
    };
  },
  computed: {
    ...mapGetters(["genderTypes", "constructTypes"])
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
    },
    handleExport() {
      if (!this.checkParam()) {
        return;
      }
      const params = JSON.parse(JSON.stringify(this.params));
      params.limit = 0;
      GetLexicalVideosList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          const leftSigns = [];
          const rightSigns = [];
          item.leftSigns.map(sign => {
            leftSigns.push(sign.name);
          });
          item.rightSigns.map(sign => {
            rightSigns.push(sign.name);
          });
          return {
            创建时间: new Date(item.createdAt),
            上次更新: new Date(item.updatedAt),
            汉语拼音音序: item.lexicalWord.initial,
            中文转写: item.lexicalWord.chinese,
            英文转写: item.lexicalWord.english,
            词性: item.lexicalWord.pos,
            构词类型: item.constructType,
            构词词语: item.constructWords.join(","),
            左手手形: leftSigns.join(","),
            右手手形: rightSigns.join(","),
            被试姓名: item.performer.name,
            被试性别: this.genderTypes[item.performer.gender].name,
            被试地区: item.performer.region.name,
            视频文件: "https://ccsl.shu.edu.cn/public/" + item.videoPath
          };
        });
        this.handleDownloadSheet(sheetData, "video");
      });
    },
    checkParam() {
      if (
        !this.params.chinese &&
        !this.params.regionID &&
        !this.params.gender &&
        !this.params.signID &&
        !this.params.leftSignID &&
        !this.params.rightSignID &&
        !this.params.constructType &&
        !this.params.constructWords
      ) {
        this.$notify({
          title: "导出失败",
          message: "请至少选择一样查询条件",
          type: "warning"
        });
        return false;
      }
      return true;
    }
  }
};
</script>

<style lang="scss" scoped>
.construct-tags {
  margin: 4px;
}
</style>