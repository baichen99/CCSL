<template>
  <div class="app-container flex-colunm">
    <el-button
      class="new"
      type="primary"
      @click="create"
    >增加</el-button>

    <el-dialog
      width="100px"
      :visible.sync="show"
      center
      @closed="onDialogClose"
    >
      <slot name="form"></slot>
    </el-dialog>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      stripe
      fit
      highlight-current-row
    >
      <slot name="column"></slot>
    </el-table>

    <el-pagination
      background
      layout="prev, pager, next"
      :page-size.sync="limit"
      :current-page.sync="page"
      :hide-on-single-page="true"
      :total="total"
    >
    </el-pagination>

  </div>
</template>

<script>
export default {
  name: "TableList",
  data() {
    return {
      show: false,
      list: null,
      listLoading: true,
      total: 0,
      page: 1,
      limit: 10
    };
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
.app-container {
  .new {
    margin: 10px 0;
  }

  .el-table {
    .edit-input {
      padding-right: 100px;
    }
    .cancel-btn {
      position: absolute;
      right: 15px;
      top: 12px;
    }
  }
}
</style>