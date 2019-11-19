<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        :placeholder="$t('tip')"
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
      <el-table v-loading="loading" :data="list" stripe border>
        <el-table-column :label="$t('CreatedAt')" align="center" width="200px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.createdAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('UpdatedAt')" align="center" width="200px">
          <template slot-scope="{row}">
            <span>{{ $d(new Date(row.updatedAt), 'long') }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('SignName')" align="center">
          <template slot-scope="{row}">
            <el-tag>{{ row.name }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column :label="$t('SignImage')" align="center">
          <template slot-scope="{row}">
            <img :src="settings.publicURL + row.image" alt="sign" style="width:80px" />
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="200px" fixed="right">
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
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <sign-form ref="form" :data="data" :mode="mode" />
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
    "tip": "请输入名称"
  },
  "en-US": {
    "tip": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import SignForm from "@/views/dashboard/form/SignForm";
import listMixin from "./listMixin";
import { GetSignsList, DeleteSign, CreateSign, UpdateSign } from "@/api/signs";

export default {
  name: "Signs",
  components: {
    SignForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        title: "",
        state: ""
      }
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  methods: {
    getList() {
      this.loading = true;
      GetSignsList(this.params)
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
      CreateSign(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateSign(id, updateData)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm(
        "此操作会删除所有视频中含有该手形的标注，此操作将永久删除, 是否继续?",
        this.$t("Warning"),
        {
          confirmButtonText: this.$t("Confirm"),
          cancelButtonText: this.$t("Cancel"),
          type: "error"
        }
      )
        .then(() => {
          DeleteSign(id).then(() => {
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
      GetSignsList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          return {
            创建时间: new Date(item.createdAt),
            上次更新: new Date(item.updatedAt),
            手形名称: item.name,
            手形图片: item.image ? this.settings.publicURL + item.image : ""
          };
        });
        this.handleDownloadSheet(sheetData, "sign");
      });
    }
  }
};
</script>

