<template>
  <list-view :get-list-method="GetJsErrorList" :allow-detail-form="false" :columns="columns">
    <template #toolbar-button>
      <el-button type="primary" size="mini" plain @click="handleDump">{{ $t("backup") }}</el-button>
    </template>
    <template #stack="{row}">
      <el-popover placement="top" width="400" trigger="hover">
        <vue-json-pretty :deep="1" :data="row.err" />
        <span slot="reference" class="stack">{{ row.err }}</span>
      </el-popover>
    </template>
    <template #storage="{row}">
      <el-popover placement="top" width="400" trigger="hover">
        <vue-json-pretty :deep="1" :data="row.store" />
        <span slot="reference" class="stack">{{ row.store }}</span>
      </el-popover>
    </template>
  </list-view>
</template>

<i18n>
{
  "zh-CN": {
    "url": "地址",
    "source": "错误来源",
    "stack": "堆栈错误信息",
    "storage": "存储数据",
    "backup": "备份"
  },
  "en-US": {
    "url": "URL",
    "source": "Error Source",
    "stack": "Stack Info",
    "storage": "Local Storage",
    "backup": "Backup Database"
  }
}
</i18n>

<script>
import { GetDatabaseDump, GetJsErrorList } from "@/api/systems";
import ListView from "@/components/ListView";
import VueJsonPretty from "vue-json-pretty";
export default {
  name: "SystemsList",
  components: {
    ListView,
    VueJsonPretty
  },
  data() {
    return {
      GetJsErrorList,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "url",
          label: this.$t("url"),
          width: "300px"
        },
        {
          prop: "info",
          label: this.$t("source"),
          width: "100px"
        },
        {
          slot: "stack",
          label: this.$t("stack"),
          width: "300px"
        },
        {
          slot: "storage",
          label: this.$t("storage"),
          width: "300px"
        }
      ]
    };
  },
  methods: {
    async handleDump() {
      try {
        const res = await GetDatabaseDump(true);
        const url = window.URL.createObjectURL(new Blob([res]));
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", `${new Date().toISOString()}.dump`);
        document.body.appendChild(link);
        link.click();
      } catch (err) {
        console.error(err);
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.stack {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}
</style>