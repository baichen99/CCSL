<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-input
        v-model="titleModel"
        prefix-icon="el-icon-search"
        :placeholder="$t('tipTitle')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
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

        <el-table-column :label="$t('Publisher')" align="center" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.creator.name }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Importance')" align="center" width="150px">
          <template slot-scope="{row}">
            <el-rate v-model="row.importance" disabled />
          </template>
        </el-table-column>

        <el-table-column
          column-key="state"
          :filters="[
            { text: $t('Draft'), value: 'draft'}, 
            { text: $t('Published'), value: 'published'},
          ]"
          :filter-multiple="false"
          :label="$t('State')"
          align="center"
          width="100px"
        >
          <template slot-scope="{row}">
            <el-tag
              size="small"
              :type="newsState[row.state].color"
            >{{ $t(newsState[row.state].name) }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          v-if="$i18n.locale==='en-US'"
          :label="$t('Title')"
          align="center"
          min-width="300px"
        >
          <template slot-scope="{row}">
            <span>{{ row.titleEn }}</span>
          </template>
        </el-table-column>

        <el-table-column v-else :label="$t('Title')" align="center" min-width="300px">
          <template slot-scope="{row}">
            <span>{{ row.titleZh }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Action')" align="center" width="250px" fixed="right">
          <template slot-scope="{row}">
            <el-button
              v-if="row.state==='draft'"
              type="success"
              size="mini"
              plain
              @click="handlePublish(row.id)"
            >{{ $t("Publish") }}</el-button>
            <el-button
              v-if="row.state==='published'"
              type="warning"
              size="mini"
              plain
              @click="handleDraft(row.id)"
            >{{ $t("Recall") }}</el-button>
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
      v-if="total>params.limit"
      background
      layout="total,prev, pager, next, jumper"
      :total="total"
      :page-size.sync="params.limit"
      :current-page.sync="params.page"
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
        <carousel-form ref="form" :data="data" :mode="mode" />
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
    "tipTitle": "请输入标题"
  },
  "en-US": {
    "tipTitle": "Input title"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import CarouselForm from "@/views/dashboard/form/CarouselForm";
import listMixin from "@/views/dashboard/listMixin";
import {
  GetCarouselsList,
  CreateCarousel,
  DeleteCarousel,
  UpdateCarousel
} from "@/api/carousel";

export default {
  name: "Carousels",
  components: {
    CarouselForm
  },
  mixins: [listMixin],
  data() {
    return {
      params: {
        order: "desc",
        titleZh: "",
        titleEn: "",
        state: ""
      }
    };
  },
  computed: {
    ...mapGetters(["newsState"]),
    titleModel: {
      get() {
        const lang = this.$i18n.locale;
        if (lang === "en-US") {
          return this.params.titleEn;
        } else {
          return this.params.titleZh;
        }
      },
      set(val) {
        const lang = this.$i18n.locale;
        if (lang === "en-US") {
          this.params.titleEn = val;
        } else {
          this.params.titleZh = val;
        }
      }
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
      GetCarouselsList(this.params)
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
      CreateCarousel(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateCarousel(id, updateData)
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
        "此操作将永久删除, 如果想暂时不显示请选择撤回, 是否继续?",
        this.$t("Warning"),
        {
          confirmButtonText: this.$t("Confirm"),
          cancelButtonText: this.$t("Cancel"),
          type: "error"
        }
      )
        .then(() => {
          DeleteCarousel(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handlePublish(id) {
      const updateData = { state: "published" };
      this.handleUpdate(id, updateData);
    },
    handleDraft(id) {
      const updateData = { state: "draft" };
      this.handleUpdate(id, updateData);
    }
  }
};
</script>
