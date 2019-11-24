<template>
  <div class="app-container flex-column">
    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border @filter-change="handleFilter">
        <el-table-column :label="$t('CreatedAt')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Name')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.user.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="IP" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.ip }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Region')" align="center" min-width="300px">
          <template slot-scope="{row}">
            <div>
              {{ row.country }}
              <span v-if="row.regionName">- {{ row.regionName }}</span>
              <span v-if="row.city">- {{ row.city }}</span>
            </div>
          </template>
        </el-table-column>

        <el-table-column :label="$t('ISP')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.isp }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Longitude')" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lon }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Latitude')" align="center" width="150px">
          <template slot-scope="{row}">
            <span>{{ row.lat }}</span>
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

<script>
import listMixin from "@/views/dashboard/listMixin";
import { GetLoginHistory } from "@/api/systems";
export default {
  name: "LoginHistory",
  mixins: [listMixin],
  methods: {
    getList() {
      this.loading = true;
      GetLoginHistory(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    }
  }
};
</script>