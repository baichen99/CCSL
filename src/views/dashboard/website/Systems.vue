<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-button type="primary" plain @click="handleDump">
        {{ $t("backup") }}
        <i class="el-icon-s-cooperation el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column :label="$t('CreatedAt')" align="center" width="200">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('url')" align="center" width="300">
          <template slot-scope="{row}">
            <span>{{ row.url }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('source')" align="center" width="150">
          <template slot-scope="{row}">
            <span>{{ row.info }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('stack')" align="center" min-width="200">
          <template slot-scope="{row}">
            <el-popover placement="top" width="400" trigger="hover">
              <vue-json-pretty :deep="1" :data="row.err" />
              <span slot="reference" class="stack">{{ row.err }}</span>
            </el-popover>
          </template>
        </el-table-column>

        <el-table-column :label="$t('storage')" align="center" min-width="200">
          <template slot-scope="{row}">
            <el-popover placement="top" width="400" trigger="hover">
              <vue-json-pretty :deep="1" :data="row.store" />
              <span slot="reference" class="stack">{{ row.store }}</span>
            </el-popover>
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
  </div>
</template>


<i18n>
{
  "zh-CN": {
    "url": "地址",
    "source": "错误来源",
    "stack": "堆栈错误信息",
    "storage": "存储数据",
    "backup": "备份数据库"
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
import VueJsonPretty from "vue-json-pretty";
import { GetDatabaseDump, GetJsErrorList } from "@/api/systems";
import listMixin from "@/views/dashboard/listMixin";
export default {
  name: "Systems",
  components: {
    VueJsonPretty
  },
  mixins: [listMixin],
  methods: {
    getList() {
      this.loading = true;
      GetJsErrorList(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDump() {
      GetDatabaseDump(true).then(res => {
        const url = window.URL.createObjectURL(new Blob([res]));
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", `${new Date().toISOString()}.dump`);
        document.body.appendChild(link);
        link.click();
      });
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