<template>
  <div class="team flex-column">
    <h2 class="title">{{ $t("TeamMenu") }}</h2>
    <el-card>
      <el-table :data="tableData">
        <el-table-column :label="$t('Name')" align="center" width="120px">
          <template slot-scope="{row}">
            <span v-if="$i18n.locale==='en-US'">{{ row.nameEn }}</span>
            <span v-else>{{ row.nameZh }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Type')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $t($options.filters.getObjectItem(memberTypes,row.type).text) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Employer')" align="center">
          <template slot-scope="{row}">
            <span v-if="$i18n.locale==='en-US'">{{ row.employerEn }}</span>
            <span v-else>{{ row.employerZh }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Degree')" align="center" width="100px">
          <template slot-scope="{row}">
            <span
              v-if="row.degree"
            >{{ $t($options.filters.getObjectItem(memberDegrees,row.degree).text) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="80px" fixed="right">
          <template slot-scope="{row}">
            <el-link
              v-if="showView(row)"
              type="primary"
              @click="showDetail(row.id)"
            >{{ $t("View") }}</el-link>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>


<script>
import { mapGetters } from "vuex";
import { GetMembersList } from "@/api/members";
export default {
  name: "Team",
  data() {
    return {
      tableData: []
    };
  },
  computed: {
    ...mapGetters(["memberTypes", "memberDegrees"])
  },
  created() {
    this.getData();
  },
  methods: {
    async getData() {
      try {
        const res = await GetMembersList({
          limit: 0,
          orderBy: "type,created_at"
        });
        this.tableData = res.data;
      } catch (err) {
        console.error(err);
      }
    },
    showDetail(id) {
      this.$router.push(`/about/team-detail/${id}`);
    },
    showView(member) {
      if (this.$i18n.locale === "en-US") {
        return member.descriptionEn;
      } else {
        return member.descriptionZh;
      }
    }
  }
};
</script>


<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.team {
  align-items: center;

  .title {
    color: $--color-primary;
    margin: 20px;
  }

  .el-card {
    width: 900px;
    margin: 1rem auto;
    padding: 1rem;
  }
}
</style>
