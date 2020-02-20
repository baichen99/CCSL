<template>
  <div class="app-container flex-column">
    <div class="table-content">
      <el-table
        v-loading="loading"
        :data="list"
        @filter-change="handleFilter"
        @row-click="handleEditItem"
        @selection-change="handleSelectItems"
      >
        <el-table-column
          v-if="deleteItemMethod"
          type="selection"
          width="45"
          align="center"
          fixed="left"
        />
        <el-table-column>
          <template slot="header">
            <div class="table-toolbar">
              <span>
                <el-button
                  v-if="createItemMethod"
                  size="mini"
                  type="primary"
                  plain
                  @click="handleNewItem"
                >{{ $t("New") }}</el-button>
                <el-button
                  v-if="exportListConfig"
                  size="mini"
                  type="info"
                  plain
                  @click="handleExportList"
                >{{ $t("Export") }}</el-button>
                <el-button
                  v-if="deleteItemMethod"
                  :disabled="!showDeleteButton"
                  size="mini"
                  type="danger"
                  plain
                  @click="handleDeleteItems"
                >{{ $t("Delete") }}</el-button>
                <slot
                  name="toolbar-button"
                  :selection="selection"
                  :handleUpdateItems="handleUpdateItems"
                ></slot>
              </span>
              <span>
                <slot name="toolbar-search" :params="params" :handleSearch="handleSearch"></slot>
              </span>
            </div>
          </template>
          <template v-for="item in columns">
            <el-table-column
              v-if="item.slot"
              #default="{row}"
              :key="item.slot"
              :prop="item.slot"
              :column-key="item.slot"
              :label="item.label"
              :min-width="item.width"
              :filters="item.filters"
              :formatter="item.formatter"
              :filter-multiple="false"
              :show-overflow-tooltip="item.hideOverflow"
              align="center"
            >
              <slot :name="item.slot" :row="row"></slot>
            </el-table-column>
            <el-table-column
              v-else
              :key="item.prop"
              :prop="item.prop"
              :column-key="item.prop"
              :label="item.label"
              :min-width="item.width"
              :filters="item.filters"
              :formatter="item.formatter"
              :filter-multiple="false"
              :show-overflow-tooltip="item.hideOverflow"
              align="center"
            ></el-table-column>
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
      @current-change="handleChangePage"
      @size-change="handleChangeLimit"
    />

    <el-drawer
      v-if="allowDetailForm"
      ref="drawer"
      :size="formDrawerSize"
      :before-close="handleCloseDrawer"
      :show-close="false"
      :visible.sync="openFormDrawer"
      :append-to-body="true"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <component
          :is="listFormComponent"
          v-if="openFormDrawer"
          ref="form"
          :data="itemData"
          :mode="mode"
        />
        <div class="form-drawer__footer">
          <el-button @click="handleCloseDrawer">{{ $t("Cancel") }}</el-button>
          <el-button
            v-if="showSaveButton"
            type="primary"
            :loading="loading"
            @click="handleSaveItem"
          >{{ loading ? $t("Saving") : $t("Save") }}</el-button>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import lodash from "lodash";
import xlsx from "xlsx";
export default {
  name: "ListView",
  props: {
    allowDetailForm: {
      type: Boolean,
      default: true
    },
    getListMethod: {
      type: Function,
      required: true,
      default: null
    },
    exportListConfig: {
      type: Function,
      default: null
    },
    createItemMethod: {
      type: Function,
      default: null
    },
    updateItemMethod: {
      type: Function,
      default: null
    },
    deleteItemMethod: {
      type: Function,
      default: null
    },
    listFormComponent: {
      type: Object,
      default: null
    },
    formDrawerSize: {
      type: String,
      default: "70%"
    },
    ignoredProperties: {
      type: Array,
      default: () => []
    },
    deleteWarning: {
      type: String,
      default: () => "此操作将永久删除, 是否继续？"
    },
    columns: {
      type: Array,
      default: () => []
    },
    order: {
      type: String,
      default: () => ""
    },
    orderBy: {
      type: String,
      default: () => ""
    },
    entity: {
      type: String,
      default: () => "data"
    }
  },
  data() {
    return {
      openFormDrawer: false,
      mode: "",
      list: [],
      selection: [],
      loading: false,
      total: 0,
      params: {
        limit: 10,
        page: 1,
        order: this.order,
        orderBy: this.orderBy
      },
      itemData: {},
      originalItemData: {}
    };
  },
  computed: {
    showSaveButton() {
      if (this.mode === "create") {
        return true;
      } else {
        // CHECK DIFF HERE
        let diffFound = false;
        for (let key in this.originalItemData) {
          if (!this.ignoredProperties.includes(key)) {
            diffFound = !lodash.isEqual(
              this.itemData[key],
              this.originalItemData[key]
            );
            if (diffFound) {
              return true;
            }
          }
        }
        return false;
      }
    },
    showDeleteButton() {
      if (this.deleteItemMethod) {
        if (this.selection.length > 0) {
          return true;
        }
      }
      return false;
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.getList();
    });
  },
  methods: {
    async getList() {
      this.loading = true;
      try {
        const res = await this.getListMethod(this.params);
        this.list = res.data;
        this.total = res.total;
      } finally {
        this.loading = false;
      }
    },
    handleSearch() {
      if (this.params.page !== 1) {
        this.params.page = 1;
        this.getList();
      } else {
        this.getList();
      }
    },
    handleFilter(val) {
      for (const key in val) {
        this.params[key] = val[key][0];
      }
      this.handleSearch();
    },
    handleChangePage(page) {
      this.params.page = page;
      this.getList();
    },
    handleChangeLimit(limit) {
      this.params.limit = limit;
      this.getList();
    },
    clearData() {
      this.itemData = {};
      this.originalItemData = {};
      this.mode = "";
    },
    makeUpdatedData() {
      let updatedData = {};
      for (let key in this.originalItemData) {
        if (!this.ignoredProperties.includes(key)) {
          if (!lodash.isEqual(this.itemData[key], this.originalItemData[key])) {
            updatedData[key] = this.itemData[key];
            if (lodash.isString(updatedData[key])) {
              updatedData[key] = updatedData[key].trim();
            }
          }
        }
      }
      return updatedData;
    },
    handleCloseDrawer() {
      this.clearData();
      const form = this.$refs.form;
      if (form) {
        form.resetFields();
      }
      this.openFormDrawer = false;
    },
    handleNewItem() {
      this.openFormDrawer = true;
      this.mode = "create";
    },
    handleEditItem(data) {
      if (this.allowDetailForm) {
        this.openFormDrawer = true;
        this.mode = "edit";
        this.originalItemData = lodash.cloneDeep(data);
        this.itemData = lodash.cloneDeep(data);
      }
    },
    async handleSaveItem() {
      switch (this.mode) {
        case "edit": {
          const updatedData = this.makeUpdatedData();
          const id = this.originalItemData.id;
          this.handleUpdateItem(id, updatedData);
          break;
        }

        case "create": {
          try {
            await this.$refs.form.checkFields();
            this.handleCreateItem(this.itemData);
          } catch (err) {
            return;
          }
          break;
        }

        default: {
          console.error("NO AVALIABLE OPERATION");
          this.handleCloseDrawer();
        }
      }
    },
    async handleUpdateItem(id, updatedData) {
      this.loading = true;
      try {
        await this.updateItemMethod(id, updatedData);
        this.handleModifyItem();
      } finally {
        this.loading = false;
      }
    },
    async handleCreateItem(data) {
      this.loading = true;
      try {
        await this.createItemMethod(data);
        this.handleModifyItem();
      } finally {
        this.loading = false;
      }
    },
    handleSelectItems(value) {
      this.selection = value;
    },
    async handleDeleteItems() {
      this.loading = true;
      try {
        const itemWarning = `将删除${this.selection.length}条数据，`;
        await this.$confirm(itemWarning + this.deleteWarning, "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        });
        for (const item of this.selection) {
          await this.deleteItemMethod(item.id);
        }
        this.handleModifyItem();
      } catch (_) {
        this.showFailedNotification();
      } finally {
        this.loading = false;
      }
    },
    async handleUpdateItems(updatedData) {
      this.loading = true;
      try {
        const itemWarning = `将修改${this.selection.length}条数据，是否继续？`;
        await this.$confirm(itemWarning, "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        });
        for (const item of this.selection) {
          await this.updateItemMethod(item.id, updatedData);
        }
        this.handleModifyItem();
      } catch (_) {
        this.showFailedNotification();
      } finally {
        this.loading = false;
      }
    },
    async handleExportList() {
      this.loading = true;
      const params = lodash.cloneDeep(this.params);
      params.limit = 0;
      try {
        const res = await this.getListMethod(params);
        const sheetData = res.data.map(this.exportListConfig);
        this.handleDownloadSheet(sheetData, this.entity);
      } catch (error) {
        console.error(error);
        this.showFailedNotification();
      } finally {
        this.loading = false;
      }
    },
    handleDownloadSheet(sheetData, filename) {
      const workbook = xlsx.utils.book_new();
      const worksheet = xlsx.utils.json_to_sheet(sheetData);
      xlsx.utils.book_append_sheet(workbook, worksheet, "Sheet1");
      xlsx.writeFile(workbook, filename + ".xlsx");
    },
    handleModifyItem(forceUpdateActionName) {
      this.showSuccessNotification();
      this.handleCloseDrawer();
      this.getList();
      if (forceUpdateActionName) {
        this.$store.dispatch(forceUpdateActionName, true);
      }
    },
    showSuccessNotification() {
      this.$notify({
        type: "success",
        title: this.$t("SuccessfulOperation")
      });
    },
    showFailedNotification() {
      this.$notify({
        type: "error",
        title: this.$t("FailedOperation")
      });
    }
  }
};
</script>

<style lang="scss">
//main-container全局样式
.app-container {
  padding: 20px;
  .table-content {
    margin: 20px;
  }

  .el-pagination {
    text-align: center;
  }
  .el-tag {
    margin: 3px;
  }
  .table-toolbar {
    .el-button,
    .el-select,
    .el-input {
      margin: 2px 5px;
    }
  }
}
</style>
