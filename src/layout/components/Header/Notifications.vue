<template>
  <div @click="handleShow">
    <svg-icon icon-class="message" />
    <el-badge v-if="unread>0" :value="unread" />

    <el-dialog
      :title="$t('Notifications')"
      :append-to-body="true"
      :center="true"
      :visible.sync="show"
      @close="getList"
    >
      <el-table ref="table" v-loading="loading" :data="list" @row-click="handleView">
        <el-table-column type="expand">
          <template #default="{row}">
            <el-form label-position="left" inline>
              <el-form-item :label="$t('Title')">
                <span>{{ row.title }}</span>
              </el-form-item>
              <el-form-item :label="$t('MessageContent')">
                <span v-html="row.message"></span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>

        <el-table-column :label="$t('CreatedAt')" align="center" width="200px">
          <template #default="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Title')" align="center" min-width="80px">
          <template #default="{row}">
            <span>
              <el-tag
                v-if="!row.readAt"
                :disable-transitions="true"
                type="danger"
                size="small"
              >{{ $t('Unread') }}</el-tag>
              {{ row.title }}
            </span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="100px" fixed="right">
          <template #default="{row}">
            <el-button
              type="danger"
              size="mini"
              plain
              @click.stop="handleDelete(row.id)"
            >{{ $t("Delete") }}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="total, prev, pager, next"
        :total="total"
        :current-page.sync="params.page"
        :hide-on-single-page="true"
      />
    </el-dialog>
  </div>
</template>

<script>
import {
  GetNotificationsList,
  GetNotification,
  DeleteNotification
} from "@/api/notifications";

export default {
  name: "Notifications",
  data: () => ({
    loading: false,
    show: false,
    list: [],
    params: {
      page: 1,
      order: "desc"
    },
    total: 0,
    unread: 0
  }),
  watch: {
    "params.page"() {
      this.getList();
    }
  },
  created() {
    this.getList();
  },
  methods: {
    async getList() {
      try {
        this.loading = true;
        const res = await GetNotificationsList(this.params);
        this.list = res.data;
        this.total = res.total;
        this.unread = res.unread;
      } finally {
        this.loading = false;
      }
    },
    handleShow() {
      this.show = true;
      this.getList();
    },
    async handleView(row) {
      await GetNotification(row.id);
      this.$refs.table.toggleRowExpansion(row);
      row.readAt = new Date().toISOString();
    },
    async handleDelete(id) {
      try {
        this.loading = true;
        await DeleteNotification(id);
        await this.getList();
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.el-pagination {
  text-align: center;
  margin: 20px auto;
}

.el-badge {
  float: right;
  position: relative;
  top: -5px;
  right: 5px;
}
</style>