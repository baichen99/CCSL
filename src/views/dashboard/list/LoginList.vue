<template>
  <list-view
    :get-list-method="GetLoginHistory"
    :allow-detail-form="false"
    :columns="columns"
    order="desc"
  >
    <template #region="{row}">
      {{ row.country }}
      <span v-if="row.regionName">- {{ row.regionName }}</span>
      <span v-if="row.city">- {{ row.city }}</span>
    </template>
  </list-view>
</template>

<script>
import ListView from "@/components/ListView";
import { GetLoginHistory } from "@/api/systems";
export default {
  name: "LoginList",
  components: {
    ListView
  },
  data() {
    return {
      GetLoginHistory,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "user.name",
          label: this.$t("Name"),
          width: "180px"
        },
        {
          prop: "ip",
          label: "IP",
          width: "180px"
        },
        {
          slot: "region",
          label: this.$t("Region"),
          width: "200px"
        },
        {
          prop: "isp",
          label: this.$t("ISP"),
          width: "180px",
          formatter: row => (row.isp === "" ? this.$t("NoData") : row.isp)
        },
        {
          prop: "lon",
          label: this.$t("Longitude"),
          width: "150px",
          formatter: row => (row.lon === 0 ? this.$t("NoData") : row.lon)
        },
        {
          prop: "lat",
          label: this.$t("Latitude"),
          width: "150px",
          formatter: row => (row.lat === 0 ? this.$t("NoData") : row.lat)
        }
      ]
    };
  }
};
</script>