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
    <template #toolbar-button="{selection,handleUpdateItems}">
      <el-button
        :disabled="!showButton(selection,'draft')"
        type="success"
        size="mini"
        plain
        @click="handleUpdateItems({state:'published'})"
      >{{ $t("Publish") }}</el-button>
      <el-button
        :disabled="!showButton(selection,'published')"
        type="warning"
        size="mini"
        plain
        @click="handleUpdateItems({state:'draft'})"
      >{{ $t("Recall") }}</el-button>
    </template>

    <template #toolbar-search="{params, handleSearch}">
      <search-input v-model="params.titleZh" :placeholder="$t('tipTitle')" @update="handleSearch" />
    </template>

    <template #state="{row}">
      <el-tag
        size="small"
        :type="$options.filters.getObjectItem(newsState,row.state).color"
      >{{ $t($options.filters.getObjectItem(newsState,row.state).text) }}</el-tag>
    </template>

    <template #importance="{row}">
      <el-rate v-model="row.importance" disabled />
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
import SearchInput from "@/components/form/SearchInput";
import NewsForm from "@/views/dashboard/form/NewsForm";
import { GetNewsList, CreateNews, UpdateNews, DeleteNews } from "@/api/news";
export default {
  name: "NewsList",
  components: {
    ListView,
    SearchInput
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
          width: "80px"
        },
        {
          prop: "column",
          label: this.$t("Column"),
          width: "100px",
          filters: [
            { text: this.$t("NewsColumn"), value: "news" },
            { text: this.$t("ActivityColumn"), value: "activity" },
            { text: this.$t("NoticeColumn"), value: "notice" },
            { text: this.$t("ResearchColumn"), value: "research" }
          ],
          formatter: row =>
            this.$t(
              this.newsColumns.data.find(item => row.column === item.value).text
            )
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
          width: "100px",
          filters: [
            { text: this.$t("Link"), value: "link" },
            { text: this.$t("Document"), value: "document" }
          ],
          formatter: row =>
            this.$t(
              this.$options.filters.getObjectItem(this.newsTypes, row.type).text
            )
        },
        {
          prop: "language",
          label: this.$t("Language"),
          width: "100px",
          filters: [
            { text: this.$t("Chinese"), value: "zh-CN" },
            { text: this.$t("English"), value: "en-US" }
          ],
          formatter: row =>
            this.$t(
              this.$options.filters.getObjectItem(
                this.languageTypes,
                row.language
              ).text
            )
        },
        {
          slot: "importance",
          label: this.$t("Importance"),
          width: "150px"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["newsColumns", "newsTypes", "newsState", "languageTypes"])
  },
  methods: {
    showButton(selection, state) {
      if (selection.length > 0) {
        const allowEnable = selection.filter(item => item.state === state);
        return allowEnable.length === selection.length;
      } else {
        return false;
      }
    }
  }
};
</script>