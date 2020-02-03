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

        <el-table-column :label="$t('HandshapeName')" align="center" prop="name" />

        <el-table-column :label="$t('HandshapeGlyph')" align="center" prop="glyph" />

        <el-table-column :label="$t('HandshapeImage')" align="center">
          <template slot-scope="{row}">
            <img :src="settings.publicURL + row.image" alt="handshape" style="width:80px" />
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
        <handshape-form ref="form" :data="data" :mode="mode" />
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
import lodash from "lodash";
import { mapGetters } from "vuex";
import HandshapeForm from "@/views/dashboard/form/HandshapeForm";
import listMixin from "@/views/dashboard/listMixin";
import {
  GetHandshapesList,
  DeleteHandshape,
  CreateHandshape,
  UpdateHandshape
} from "@/api/handshapes";

export default {
  name: "Handshapes",
  components: {
    HandshapeForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        name: ""
      }
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  created() {
    this.$nextTick(() => {
      this.getList();
    });
  },
  methods: {
    getList() {
      this.loading = true;
      GetHandshapesList(this.params)
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
      CreateHandshape(data)
        .then(() => {
          this.handleModify("data/getHandshapes");
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateHandshape(id, updateData)
        .then(() => {
          this.handleModify("data/getHandshapes");
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
          DeleteHandshape(id).then(() => {
            this.handleModify("data/getHandshapes");
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handleExport() {
      const params = lodash.cloneDeep(this.params);
      params.limit = 0;
      GetHandshapesList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          return {
            [this.$t("CreatedAt")]: new Date(item.createdAt),
            [this.$t("UpdatedAt")]: new Date(item.updatedAt),
            [this.$t("HandshapeName")]: item.name,
            [this.$t("HandshapeGlyph")]: item.glyph
          };
        });
        this.handleDownloadSheet(sheetData, "handshape");
      });
    }
  }
};
</script>

