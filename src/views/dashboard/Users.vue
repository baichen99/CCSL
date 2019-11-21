<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-input
        v-model="params.username"
        :placeholder="$t('tipAccount')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-input
        v-model="params.name"
        :placeholder="$t('tipName')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-button type="primary" plain @click="handleNew">
        {{ $t("New") }}
        <i class="el-icon-plus el-icon--right" />
      </el-button>
      <el-button type="primary" plain @click="handleExport">
        {{ $t("Export") }}
        <i class="el-icon-download el-icon--right" />
      </el-button>
    </div>
    <div class="table-content">
      <el-table v-loading="loading" :data="list" stripe border @filter-change="handleFilter">
        <el-table-column :label="$t('CreatedAt')" align="center" width="180px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column
          column-key="state"
          :label="$t('AccountState')"
          :filters="[
            { text: $t('Active'), value: 'active'}, 
            { text: $t('Inactive'), value: 'inactive'},
          ]"
          :filter-multiple="false"
          align="center"
          width="150px"
        >
          <template slot-scope="{row}">
            <el-tag :type="userState[row.state].color">{{ $t(userState[row.state].name) }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          column-key="userType"
          :filters="[
            { text: $t('SuperAdmin'), value: 'super'}, 
            { text: $t('Admin'), value: 'admin'},
            { text: $t('Learner'), value: 'learner'},
            { text: $t('User'), value: 'user'},
          ]"
          :filter-multiple="false"
          :label="$t('UserRole')"
          align="center"
          width="150px"
        >
          <template slot-scope="{row}">
            <el-tag :type="userTypes[row.userType].color">{{ $t(userTypes[row.userType].name) }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Account')" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.username }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Name')" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="250px" fixed="right">
          <template slot-scope="{row}">
            <el-button
              v-if="row.state==='inactive'"
              :disabled="isSuperUser(row)"
              type="success"
              size="mini"
              plain
              @click="handleActive(row.id)"
            >{{ $t("Enable") }}</el-button>
            <el-button
              v-if="row.state==='active'"
              :disabled="isSuperUser(row)"
              type="warning"
              size="mini"
              plain
              @click="handleInactive(row.id)"
            >{{ $t("Disable") }}</el-button>
            <el-button
              :disabled="isSuperUser(row)"
              type="primary"
              size="mini"
              plain
              @click="handleEdit(row)"
            >{{ $t("Edit") }}</el-button>
            <el-button
              :disabled="isSuperUser(row)"
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
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <user-form ref="form" :data="data" :mode="mode" />
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

<i18n>
{
  "zh-CN": {
    "tipAccount": "请输入账号",
    "tipName": "请输入姓名"
  },
  "en-US": {
    "tipAccount": "Input account",
    "tipName": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import UserForm from "@/views/dashboard/form/UserForm";
import listMixin from "@/views/dashboard/listMixin";
import { GetUsersList, CreateUser, UpdateUser, DeleteUser } from "@/api/users";
export default {
  name: "Users",
  components: {
    UserForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        username: "",
        name: "",
        userType: "",
        state: ""
      }
    };
  },
  computed: {
    ...mapGetters(["userTypes", "userState"])
  },
  methods: {
    getList() {
      this.loading = true;
      GetUsersList(this.params)
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
      CreateUser(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateUser(id, updateData)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm("此操作将永久删除, 是否继续?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "error"
      })
        .then(() => {
          DeleteUser(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handleExport() {
      const params = JSON.parse(JSON.stringify(this.params));
      params.limit = 0;
      GetUsersList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          return {
            创建时间: new Date(item.createdAt),
            上次更新: new Date(item.updatedAt),
            姓名: item.name,
            账号: item.username,
            状态: this.userState[item.state].name,
            用户类型: this.userTypes[item.userType].name
          };
        });
        this.handleDownloadSheet(sheetData, "user");
      });
    },
    handleActive(id) {
      const updateData = { state: "active" };
      this.handleUpdate(id, updateData);
    },
    handleInactive(id) {
      const updateData = { state: "inactive" };
      this.handleUpdate(id, updateData);
    },
    isSuperUser(row) {
      return row.userType === "super";
    }
  }
};
</script>
