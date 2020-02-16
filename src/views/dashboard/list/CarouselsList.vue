<template>
  <list-view
    :get-list-method="GetCarouselsList"
    :create-item-method="CreateCarousel"
    :update-item-method="UpdateCarousel"
    :delete-item-method="DeleteCarousel"
    :list-form-component="CarouselForm"
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
import CarouselForm from "@/views/dashboard/form/CarouselForm";
import {
  GetCarouselsList,
  CreateCarousel,
  DeleteCarousel,
  UpdateCarousel
} from "@/api/carousel";
export default {
  name: "CarouselsList",
  components: {
    ListView,
    SearchInput
  },
  data() {
    return {
      GetCarouselsList,
      CreateCarousel,
      UpdateCarousel,
      DeleteCarousel,
      CarouselForm,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "creator.name",
          label: this.$t("Publisher"),
          width: "100px"
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
          prop: "title",
          label: this.$t("Title"),
          width: "300px",
          formatter: row =>
            this.$i18n.locale === "en-US" ? row.titleEn : row.titleZh,
          hideOverflow: true
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
    ...mapGetters(["newsState"])
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

<style lang="scss" scoped>
.el-tag {
  margin: 3px;
}
</style>