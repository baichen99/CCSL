<template>
  <div @click="handleShow">
    <svg-icon icon-class="message" />
    <el-dialog
      :title="$t('Notifications')"
      :append-to-body="true"
      :center="true"
      :visible.sync="show"
    >
      <el-table v-loading="loading" :data="list" border>
        <el-table-column fixed :label="$t('CreatedAt')" align="center" width="200px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Title')" align="center" min-width="80px">
          <template slot-scope="{row}">
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

        <el-table-column :label="$t('MessageContent')" align="center" min-width="100px">
          <template slot-scope="{row}">
            <span class="message">{{ row.message }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="200px" fixed="right">
          <template slot-scope="{row}">
            <el-button
              type="primary"
              size="mini"
              plain
              @click="handleView(row,index)"
            >{{ $t("View") }}</el-button>
            <el-button
              type="danger"
              size="mini"
              plain
              @click="handleDelete(row.id)"
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
    total: 0
  }),
  watch: {
    "params.page"() {
      this.getList();
    }
  },
  created() {
    this.$nextTick(() => {
      this.getList();
    });
  },
  methods: {
    getList() {
      this.loading = true;
      GetNotificationsList(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleShow() {
      this.show = true;
    },
    handleView(data) {
      data.readAt = new Date().toISOString();
      this.$alert(data.message, data.title);
      GetNotification(data.id);
    },
    handleDelete(id) {
      this.loading = true;
      DeleteNotification(id)
        .then(() => {
          this.getList();
        })
        .catch(() => {
          this.loading = false;
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.el-pagination {
  text-align: center;
  margin: 20px auto;
}

.message {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
</style>