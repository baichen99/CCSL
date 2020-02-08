<template>
  <list-view
    :get-list-method="GetCarouselsList"
    :create-item-method="CreateCarousel"
    :update-item-method="UpdateCarousel"
    :delete-item-method="DeleteCarousel"
    :list-form-component="CarouselForm"
    :columns="columns"
  >
    <template #toolbar="{params, handleSearch}">
      <el-input
        v-model="params.titleZh"
        prefix-icon="el-icon-search"
        :placeholder="$t('tipTitle')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
    </template>
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
    ListView
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
            this.$i18n.locale === "en-US" ? row.titleEn : row.titleZh
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
    ...mapGetters(["newsState"])
  }
};
</script>

<style lang="scss" scoped>
.el-tag {
  margin: 3px;
}
</style>