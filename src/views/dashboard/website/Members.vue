<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-button type="primary" plain @click="handleNew">
        {{ $t("New") }}
        <i class="el-icon-plus el-icon--right" />
      </el-button>
    </div>

    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border @filter-change="handleFilter">
        <el-table-column :label="$t('CreatedAt')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Name')" align="center" width="150px">
          <template slot-scope="{row}">
            <span v-if="$i18n.locale==='en-US'">{{ row.nameEn }}</span>
            <span v-else>{{ row.nameZh }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Type')" align="center" width="200px">
          <template slot-scope="{row}">
            <span>{{ $t(memberTypes[row.type].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Employer')" align="center" min-width="300px">
          <template slot-scope="{row}">
            <span v-if="$i18n.locale==='en-US'">{{ row.employerEn }}</span>
            <span v-else>{{ row.employerZh }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="180px" fixed="right">
          <template slot-scope="{row}">
            <el-button type="primary" size="mini" plain @click="handleEdit(row)">{{ $t("Edit") }}</el-button>
            <el-button
              type="danger"
              size="mini"
              plain
              @click="handleDelete(row.id)"
            >{{ $t("Delete") }}</el-button>
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

    <el-drawer
      v-if="show"
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <member-form ref="form" :data="data" :mode="mode" />
        <div class="form-drawer__footer">
          <el-button @click="handleClose">{{ $t("Cancel") }}</el-button>
          <el-button
            v-if="checkDiff"
            type="primary"
            :loading="loading"
            @click="handleSave"
          >{{ loading ? $t("Saving") : $t("Save") }}</el-button>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import MemberForm from "@/views/dashboard/form/MemberForm";
import listMixin from "@/views/dashboard/listMixin";
import {
  GetMembersList,
  CreateMember,
  UpdateMember,
  DeleteMember
} from "@/api/members";
import { mapGetters } from "vuex";

export default {
  name: "Members",
  components: {
    MemberForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        nameZh: ""
      }
    };
  },
  computed: {
    ...mapGetters(["memberTypes"])
  },
  methods: {
    getList() {
      this.loading = true;
      GetMembersList(this.params)
        .then(res => {
          this.list = res.data;
          this.total = res.total;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleCreate(data) {
      CreateMember(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateMember(id, updateData)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm("此操作将永久删除, 是否继续?", this.$t("Warning"), {
        confirmButtonText: this.$t("Confirm"),
        cancelButtonText: this.$t("Cancel"),
        type: "error"
      })
        .then(() => {
          DeleteMember(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    }
  }
};
</script>