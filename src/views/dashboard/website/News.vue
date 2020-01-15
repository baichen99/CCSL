<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-input
        v-model="params.title"
        prefix-icon="el-icon-search"
        :placeholder="$t('tipTitle')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-button type="primary" plain @click="handleNew">
        {{ $t("New") }}
        <i class="el-icon-plus el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border @filter-change="handleFilter">
        <el-table-column :label="$t('Date')" align="center" width="130px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.date), 'short') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Publisher')" align="center" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.creator.name }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Importance')" align="center" width="150px">
          <template slot-scope="{row}">
            <el-rate v-model="row.importance" disabled />
          </template>
        </el-table-column>

        <el-table-column
          column-key="state"
          :filters="[
            { text: $t('Draft'), value: 'draft'}, 
            { text: $t('Published'), value: 'published'},
          ]"
          :filter-multiple="false"
          :label="$t('State')"
          align="center"
          width="100px"
        >
          <template slot-scope="{row}">
            <el-tag
              size="small"
              :type="newsState[row.state].color"
            >{{ $t(newsState[row.state].name) }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          column-key="column"
          :filters="[
            { text: $t('NewsColumn'), value: 'news'}, 
            { text: $t('ActivityColumn'), value: 'activity'},
            { text: $t('NoticeColumn'), value: 'notice'},
            { text: $t('ResearchColumn'), value: 'research'},
          ]"
          :filter-multiple="false"
          :label="$t('Column')"
          align="center"
          width="120px"
        >
          <template slot-scope="{row}">
            <span>{{ $t(newsColumns[row.column].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Title')" align="center" min-width="300px">
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
          </template>
        </el-table-column>

        <el-table-column
          column-key="type"
          :filters="[
            { text: $t('Link'), value: 'link'}, 
            { text: $t('Document'), value: 'document'},
          ]"
          :filter-multiple="false"
          :label="$t('Type')"
          align="center"
          width="120px"
        >
          <template slot-scope="{row}">
            <span>{{ $t(newsTypes[row.type].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column
          column-key="language"
          :filters="[
            { text: $t('Chinese'), value: 'zh-CN'}, 
            { text: $t('English'), value: 'en-US'},
          ]"
          :filter-multiple="false"
          :label="$t('Language')"
          align="center"
          width="120px"
        >
          <template slot-scope="{row}">
            <span>{{ $t(languageTypes[row.language].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="250px" fixed="right">
          <template slot-scope="{row}">
            <el-button
              v-if="row.state==='draft'"
              type="success"
              size="mini"
              plain
              @click="handlePublish(row.id)"
            >{{ $t("Publish") }}</el-button>
            <el-button
              v-if="row.state==='published'"
              type="warning"
              size="mini"
              plain
              @click="handleDraft(row.id)"
            >{{ $t("Recall") }}</el-button>
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
      v-if="show"
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <news-form ref="form" :data="data" :mode="mode" />
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

<i18n>
{
  "zh-CN": {
    "tipTitle": "请输入标题"
  },
  "en-US": {
    "tipTitle": "Input title"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import NewsForm from "@/views/dashboard/form/NewsForm";
import listMixin from "@/views/dashboard/listMixin";
import { GetNewsList, CreateNews, UpdateNews, DeleteNews } from "@/api/news";

export default {
  name: "News",
  components: {
    NewsForm
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
  computed: {
    ...mapGetters(["newsColumns", "newsTypes", "newsState", "languageTypes"])
  },
  created() {
    this.$nextTick(() => {
      this.getList();
    });
  },
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
          confirmButtonText: "确认",
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
