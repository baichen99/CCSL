<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        placeholder="请输入中文"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-input
        v-model="params.english"
        prefix-icon="el-icon-search"
        placeholder="请输入英文"
        clearable
        @change="handleSearch"
      />
      <word-pos-selector v-model="params.pos" @update="handleSearch" />
      <el-button type="primary" plain @click="handleNew">
        {{ $t("New") }}
        <i class="el-icon-plus el-icon--right" />
      </el-button>
      <el-button type="primary" plain @click="handleExport">
        {{ $t("Export") }}
        <i class="el-icon-download el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border @filter-change="handleFilter">
        <el-table-column
          column-key="initial"
          :filters="initialFilters"
          :filter-multiple="false"
          :label="$t('Initial')"
          align="center"
          width="80px"
        >
          <template slot-scope="{row}">
            <span>{{ row.initial }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Chinese')" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span class="word-sup" v-html="$options.filters.addNumberSup(row.chinese) " />
          </template>
        </el-table-column>

        <el-table-column :label="$t('English')" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.english }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('PoS')" align="center" min-width="100px">
          <template slot-scope="{row}">
            <span>{{ row.pos }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="200px" fixed="right">
          <template slot-scope="{row}">
            <el-button type="primary" size="mini" plain @click="handleEdit(row)">{{ $t("Edit") }}</el-button>
            <el-button
              type="danger"
              size="mini"
              plain
              @click="handleDelete(row.id)"
            >{{ $t("Delete") }}</el-button>
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
      size="30%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <lexicon-form ref="form" :data="data" :mode="mode" />
        <div class="form-drawer__footer">
          <el-button @click="handleClose">{{ $t("Cancel") }}</el-button>
          <el-button
            v-if="checkDiff"
            type="primary"
            :loading="loading"
            @click="handleSave"
          >{{ loading ? $t("Saving") : $t("Save") }}</el-button>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import LexiconForm from "@/views/dashboard/form/LexiconForm";
import WordPosSelector from "@/components/form/WordPosSelector";
import listMixin from "@/views/dashboard/listMixin";
import { mapGetters } from "vuex";
import {
  GetWordsList,
  CreateWord,
  UpdateWord,
  DeleteWord
} from "@/api/lexicons";

export default {
  name: "Lexicons",
  components: {
    WordPosSelector,
    LexiconForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        initial: "",
        chinese: "",
        english: "",
        pos: "",
        orderBy: "initial, id"
      },
      initialFilters: []
    };
  },
  computed: {
    ...mapGetters(["wordInitial"])
  },
  created() {
    this.wordInitial.map(item => {
      this.initialFilters.push({
        text: item,
        value: item
      });
    });
  },
  methods: {
    getList() {
      this.loading = true;
      GetWordsList(this.params)
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
      CreateWord(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateWord(id, updateData)
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
          DeleteWord(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handleExport() {
      const params = JSON.parse(JSON.stringify(this.params));
      params.limit = 0;
      GetWordsList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          return {
            [this.$t("CreatedAt")]: new Date(item.createdAt),
            [this.$t("UpdatedAt")]: new Date(item.updatedAt),
            [this.$t("Initial")]: item.initial,
            [this.$t("Chinese")]: item.chinese,
            [this.$t("English")]: item.english,
            [this.$t("PoS")]: item.pos
          };
        });
        this.handleDownloadSheet(sheetData, "lexicon");
      });
    }
  }
};
</script>