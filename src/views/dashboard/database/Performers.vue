<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.name"
        prefix-icon="el-icon-search"
        :placeholder="$t('tipName')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
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

        <el-table-column :label="$t('Name')" align="center" prop="name" />

        <el-table-column
          column-key="gender"
          :filters="[
            { text: $t('Male'), value: 'M'}, 
            { text: $t('Female'), value: 'F'},
          ]"
          :filter-multiple="false"
          :label="$t('Gender')"
          width="100px"
          align="center"
        >
          <template slot-scope="{row}">
            <span>{{ $t(genderTypes[row.gender].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Region')" align="center" min-width="200px">
          <template slot-scope="{row}">
            <span class="hide-overflow">{{ row.regionID | getRegionName }}</span>
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
      ref="drawer"
      size="60%"
      :before-close="handleClose"
      :show-close="false"
      :destroy-on-close="true"
      :visible.sync="show"
      direction="rtl"
    >
      <div class="form-drawer__content">
        <performer-form ref="form" :data="data" :mode="mode" />
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
    "tipName": "请输入姓名"
  },
  "en-US": {
    "tipName": "Input name"
  }
}
</i18n>

<script>
import lodash from "lodash";
import { mapGetters } from "vuex";
import PerformerForm from "@/views/dashboard/form/PerformerForm";
import CitySelector from "@/components/form/CitySelector";
import listMixin from "@/views/dashboard/listMixin";
import {
  GetPerformersList,
  CreatePerformer,
  UpdatePerformer,
  DeletePerformer
} from "@/api/performers";
export default {
  name: "Performers",
  components: {
    PerformerForm,
    CitySelector
  },
  mixins: [listMixin],
  data() {
    return {
      removeProperties: [],
      params: {
        regionID: undefined,
        name: "",
        gender: ""
      }
    };
  },
  computed: {
    ...mapGetters(["genderTypes"])
  },
  created() {
    this.$nextTick(() => {
      this.getList();
    });
  },
  methods: {
    getList() {
      this.loading = true;
      GetPerformersList(this.params)
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
      CreatePerformer(data)
        .then(() => {
          this.handleModify("data/getPerformers");
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdatePerformer(id, updateData)
        .then(() => {
          this.handleModify("data/getPerformers");
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleDelete(id) {
      this.loading = true;
      this.$confirm(
        "此操作会删除所有视频中含有该被试信息的标注，此操作将永久删除, 是否继续?",
        this.$t("Warning"),
        {
          confirmButtonText: this.$t("Confirm"),
          cancelButtonText: this.$t("Cancel"),
          type: "error"
        }
      )
        .then(() => {
          DeletePerformer(id).then(() => {
            this.handleModify("data/getPerformers");
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handleExport() {
      const params = lodash.cloneDeep(this.params);
      params.limit = 0;
      GetPerformersList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          return {
            [this.$t("CreatedAt")]: new Date(item.createdAt),
            [this.$t("UpdatedAt")]: new Date(item.updatedAt),
            [this.$t("Name")]: item.name,
            [this.$t("Gender")]: this.$t(this.genderTypes[item.gender].name),
            [this.$t("Region")]: this.$options.filters.getRegionName(
              item.regionID
            )
          };
        });
        this.handleDownloadSheet(sheetData, "performer");
      });
    }
  }
};
</script>