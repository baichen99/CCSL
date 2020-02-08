<template>
  <list-view
    :get-list-method="GetPerformersList"
    :create-item-method="CreatePerformer"
    :update-item-method="UpdatePerformer"
    :delete-item-method="DeletePerformer"
    :list-form-component="PerformerForm"
    :export-list-config="exportConfig"
    :columns="columns"
    delete-warning="此操作会删除所有视频中含有该被试信息的标注，此操作将永久删除, 是否继续?"
    entity="performers"
  >
    <template #toolbar="{params,handleSearch}">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        :placeholder="$t('tipName')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
    </template>

    <template #action="{row,handleDeleteItem}">
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
    "tipName": "请输入姓名"
  },
  "en-US": {
    "tipName": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import PerformerForm from "@/views/dashboard/form/PerformerForm";
import ListView from "@/components/ListView";
import CitySelector from "@/components/form/CitySelector";
import {
  GetPerformersList,
  CreatePerformer,
  UpdatePerformer,
  DeletePerformer
} from "@/api/performers";
export default {
  name: "PerformersList",
  components: {
    ListView,
    CitySelector
  },
  data() {
    return {
      GetPerformersList,
      CreatePerformer,
      UpdatePerformer,
      DeletePerformer,
      PerformerForm,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "updatedAt",
          label: this.$t("UpdatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.updatedAt), "long")
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "150px"
        },
        {
          prop: "gender",
          label: this.$t("Gender"),
          width: "100px",
          formatter: row => this.$t(this.genderTypes[row.gender].name)
        },
        {
          prop: "region",
          label: this.$t("Region"),
          width: "200px",
          hideOverflow: true,
          formatter: row => this.$options.filters.getRegionName(row.regionID)
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "100px"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["genderTypes"])
  },
  methods: {
    exportConfig(item) {
      return {
        [this.$t("CreatedAt")]: new Date(item.createdAt),
        [this.$t("UpdatedAt")]: new Date(item.updatedAt),
        [this.$t("Name")]: item.name,
        [this.$t("Gender")]: this.$t(this.genderTypes[item.gender].name),
        [this.$t("Region")]: this.$options.filters.getRegionName(item.regionID)
      };
    }
  }
};
</script>