<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        :placeholder="$t('chineseTip')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
      <sign-selector v-model="params.signID" @update="handleSearch" />
      <sign-selector v-model="params.leftSignID" orientation="left" @update="handleSearch" />
      <sign-selector v-model="params.rightSignID" orientation="right" @update="handleSearch" />
      <morphemes-input v-model="params.morpheme" @update="handleSearch" @enter="handleSearch" />
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
        <el-table-column
          column-key="initial"
          :filters="initialFilters"
          :filter-multiple="false"
          :label="$t('Initial')"
          align="center"
          width="80px"
        >
          <template slot-scope="{row}">
            <span>{{ lexicons[row.lexiconID].initial }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Region')" align="center" width="200px">
          <template slot-scope="{row}">
            <span>{{ performers[row.performerID].regionID | getRegionName }}</span>
          </template>
        </el-table-column>

        <el-table-column
          column-key="gender"
          :filters="[
            { text: $t('Male'), value: 'M'}, 
            { text: $t('Female'), value: 'F'},
          ]"
          :filter-multiple="false"
          :label="$t('Gender')"
          align="center"
          width="100px"
        >
          <template slot-scope="{row}">
            <span>{{ $t(genderTypes[performers[row.performerID].gender].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column :label="$t('Name')" align="center" width="120px">
          <template slot-scope="{row}">
            <span>{{ performers[row.performerID].name }}</span>
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('Chinese')"
          align="center"
          min-width="150px"
          show-overflow-tooltip
        >
          <template slot-scope="{row}">
            <span
              class="word-sup"
              v-html="$options.filters.addNumberSup(lexicons[row.lexiconID].chinese)"
            />
          </template>
        </el-table-column>

        <el-table-column
          column-key="wordFormation"
          :filters="[
            { text: $t('SimpleWord'), value: 'simple'}, 
            { text: $t('CompoundWord'), value: 'compound'},
          ]"
          :filter-multiple="false"
          :label="$t('WordFormation')"
          align="center"
          width="150px"
        >
          <template slot-scope="{row}">
            <span v-if="row.wordFormation">{{ $t(wordFormations[row.wordFormation].name) }}</span>
            <el-tag v-else :disable-transitions="true" size="small" type="info">{{ $t("NoData") }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('Morpheme')"
          align="center"
          min-width="200px"
          show-overflow-tooltip
        >
          <template slot-scope="{row}">
            <div v-if="row.morpheme.length > 0">
              <el-tag
                v-for="value in row.morpheme"
                :key="value"
                :disable-transitions="true"
                class="tags"
                size="small"
              >
                <span v-html="$options.filters.addNumberSup(value)"></span>
              </el-tag>
            </div>
            <div v-else>
              <el-tag :disable-transitions="true" size="small" type="info">{{ $t("NoData") }}</el-tag>
            </div>
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
        <lexical-video-form ref="form" :data="data" :mode="mode" />
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
    "chineseTip": "请输入中文"
  },
  "en-US": {
    "chineseTip": "Input Chinese"
  }
}
</i18n>

<script>
import lodash from "lodash";
import { mapGetters } from "vuex";
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import CitySelector from "@/components/form/CitySelector";
import SignSelector from "@/components/form/SignSelector";
import MorphemesInput from "@/components/form/MorphemesInput";
import listMixin from "@/views/dashboard/listMixin";
import {
  GetLexicalVideosList,
  CreateLexicalVideo,
  UpdateLexicalVideo,
  DeleteLexicalVideo
} from "@/api/videos";
export default {
  name: "LexicalVideos",
  components: {
    LexicalVideoForm,
    CitySelector,
    SignSelector,
    MorphemesInput
  },
  mixins: [listMixin],
  data() {
    return {
      removeProperties: [],
      params: {
        chinese: "",
        regionID: undefined,
        gender: "",
        signID: "",
        leftSignID: "",
        rightSignID: "",
        wordFormation: "",
        morpheme: ""
      },
      initialFilters: []
    };
  },
  computed: {
    ...mapGetters([
      "genderTypes",
      "wordFormations",
      "wordInitial",
      "performers",
      "lexicons",
      "signs"
    ])
  },
  created() {
    this.$nextTick(async () => {
      await this.$store.dispatch("data/getPerformers");
      await this.$store.dispatch("data/getLexicons");
      await this.$store.dispatch("data/getSigns");
      await this.getList();
      this.wordInitial.map(item => {
        this.initialFilters.push({
          text: item,
          value: item
        });
      });
    });
  },
  methods: {
    getList() {
      this.loading = true;
      GetLexicalVideosList(this.params)
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
      CreateLexicalVideo(data)
        .then(() => {
          this.handleModify();
        })
        .catch(() => {
          this.loading = false;
        });
    },
    handleUpdate(id, updateData) {
      UpdateLexicalVideo(id, updateData)
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
        "此操作将永久删除该视频以及其相关的标注信息，是否继续?",
        this.$t("Warning"),
        {
          confirmButtonText: this.$t("Confirm"),
          cancelButtonText: this.$t("Cancel"),
          type: "error"
        }
      )
        .then(() => {
          DeleteLexicalVideo(id).then(() => {
            this.handleModify();
          });
        })
        .catch(() => {
          this.showCancel();
        });
    },
    handleExport() {
      const params = lodash.cloneDeep(this.params);
      params.limit = 0;
      GetLexicalVideosList(params, true).then(res => {
        const sheetData = res.data.map(item => {
          const leftSigns = [];
          const rightSigns = [];
          item.leftSignsID.map(signID => {
            const sign = this.signs[signID];
            leftSigns.push(sign.name);
          });
          item.rightSignsID.map(signID => {
            const sign = this.signs[signID];
            rightSigns.push(sign.name);
          });
          return {
            [this.$t("CreatedAt")]: new Date(item.createdAt),
            [this.$t("UpdatedAt")]: new Date(item.updatedAt),
            [this.$t("Initial")]: this.lexicons[item.lexiconID].initial,
            [this.$t("Chinese")]: this.lexicons[item.lexiconID].chinese,
            [this.$t("English")]: this.lexicons[item.lexiconID].english,
            [this.$t("PoS")]: this.lexicons[item.lexiconID].pos,
            [this.$t("WordFormation")]: item.wordFormation
              ? this.$t(this.wordFormations[item.wordFormation].name)
              : this.$t("NoData"),
            [this.$t("Morpheme")]: item.morpheme.join(","),
            [this.$t("LeftSign")]: leftSigns.join(","),
            [this.$t("RightSign")]: rightSigns.join(","),
            [this.$t("Name")]: this.performers[item.performerID].name,
            [this.$t("Gender")]: this.$t(
              this.genderTypes[this.performers[item.performerID].gender].name
            ),
            [this.$t("Region")]: this.$options.filters.getRegionName(
              this.performers[item.performerID].regionID
            )
          };
        });
        this.handleDownloadSheet(sheetData, "video");
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.tags {
  margin: 4px;
}
</style>