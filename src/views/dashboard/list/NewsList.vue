<template>
  <list-view
    :get-list-method="GetNewsList"
    :create-item-method="CreateNews"
    :update-item-method="UpdateNews"
    :delete-item-method="DeleteNews"
    :list-form-component="NewsForm"
    :columns="columns"
    order="desc"
  >
    <template #state="{row}">
      <el-tag size="small" :type="newsState[row.state].color">{{ $t(newsState[row.state].name) }}</el-tag>
    </template>
    <template #importance="{row}">
      <el-rate v-model="row.importance" disabled />
    </template>

    <template #action="{row,handleUpdateItem,handleDeleteItem}">
      <el-button
        v-if="row.state==='draft'"
        type="success"
        size="mini"
        plain
        @click.stop="handleUpdateItem(row.id, {
          state: 'published'
        })"
      >{{ $t("Publish") }}</el-button>
      <el-button
        v-if="row.state==='published'"
        type="warning"
        size="mini"
        plain
        @click.stop="handleUpdateItem(row.id, {
          state: 'draft'
        })"
      >{{ $t("Recall") }}</el-button>
      <el-button
        type="danger"
        size="mini"
        plain
        @click.stop="handleDeleteItem(row.id)"
      >{{ $t("Delete") }}</el-button>
    </template>
  </list-view>
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
import ListView from "@/components/ListView";
import NewsForm from "@/views/dashboard/form/NewsForm";
import { GetNewsList, CreateNews, UpdateNews, DeleteNews } from "@/api/news";
export default {
  name: "NewsList",
  components: {
    ListView
  },
  data() {
    return {
      GetNewsList,
      CreateNews,
      UpdateNews,
      DeleteNews,
      NewsForm,
      columns: [
        {
          prop: "date",
          label: this.$t("Date"),
          width: "130px",
          formatter: row => this.$d(new Date(row.date), "short")
        },
        {
          slot: "state",
          label: this.$t("State"),
          width: "100px",
          filters: [
            { text: this.$t("Draft"), value: "draft" },
            { text: this.$t("Published"), value: "published" }
          ]
        },
        {
          prop: "creator.name",
          label: this.$t("Publisher"),
          width: "100px"
        },
        {
          prop: "column",
          label: this.$t("Column"),
          width: "120px",
          filters: [
            { text: this.$t("NewsColumn"), value: "news" },
            { text: this.$t("ActivityColumn"), value: "activity" },
            { text: this.$t("NoticeColumn"), value: "notice" },
            { text: this.$t("ResearchColumn"), value: "research" }
          ],
          formatter: row => this.$t(this.newsColumns[row.column].name)
        },
        {
          prop: "title",
          label: this.$t("Title"),
          width: "300px",
          hideOverflow: true
        },
        {
          prop: "type",
          label: this.$t("Type"),
          width: "120px",
          filters: [
            { text: this.$t("Link"), value: "link" },
            { text: this.$t("Document"), value: "document" }
          ],
          formatter: row => this.$t(this.newsTypes[row.type].name)
        },
        {
          prop: "language",
          label: this.$t("Language"),
          width: "120px",
          filters: [
            { text: this.$t("Chinese"), value: "zh-CN" },
            { text: this.$t("English"), value: "en-US" }
          ],
          formatter: row => this.$t(this.languageTypes[row.language].name)
        },
        {
          slot: "importance",
          label: this.$t("Importance"),
          width: "150px"
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "180px",
          fixed: "right"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["newsColumns", "newsTypes", "newsState", "languageTypes"])
  }
};
</script>