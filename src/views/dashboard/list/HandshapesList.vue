<template>
  <list-view
    :get-list-method="GetHandshapesList"
    :create-item-method="CreateHandshape"
    :update-item-method="UpdateHandshape"
    :delete-item-method="DeleteHandshape"
    :list-form-component="HandshapeForm"
    :export-list-config="exportConfig"
    :columns="columns"
    delete-warning="删除手形会删除视频中该手形对应的所有标注，是否继续？"
    entity="handshape"
  >
    <template #toolbar="{params,handleSearch}">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        :placeholder="$t('tip')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
    </template>

    <template #image="{row}">
      <img :src="settings.publicURL + row.image" alt="handshape" style="width:80px" />
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
    "tip": "请输入手形名称"
  },
  "en-US": {
    "tip": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import ListView from "@/components/ListView";
import HandshapeForm from "@/views/dashboard/form/HandshapeForm";
import {
  GetHandshapesList,
  CreateHandshape,
  UpdateHandshape,
  DeleteHandshape
} from "@/api/handshapes";
export default {
  name: "HandshapeList",
  components: {
    ListView
  },
  data() {
    return {
      HandshapeForm,
      GetHandshapesList,
      CreateHandshape,
      UpdateHandshape,
      DeleteHandshape,
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
          label: this.$t("HandshapeName"),
          width: "100px"
        },
        {
          prop: "glyph",
          label: this.$t("HandshapeGlyph"),
          width: "120px"
        },
        {
          slot: "image",
          label: this.$t("HandshapeImage"),
          width: "150px"
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "90px",
          fixed: "right"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  methods: {
    exportConfig(item) {
      return {
        [this.$t("CreatedAt")]: new Date(item.createdAt),
        [this.$t("UpdatedAt")]: new Date(item.updatedAt),
        [this.$t("HandshapeName")]: item.name,
        [this.$t("HandshapeGlyph")]: item.glyph
      };
    }
  }
};
</script>