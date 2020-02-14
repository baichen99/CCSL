<template>
  <list-view
    :get-list-method="GetLexicalVideosList"
    :create-item-method="CreateLexicalVideo"
    :update-item-method="UpdateLexicalVideo"
    :delete-item-method="DeleteLexicalVideo"
    :list-form-component="LexicalVideoForm"
    :export-list-config="exportConfig"
    :columns="columns"
    delete-warning="此操作将永久删除该视频以及其相关的标注信息，是否继续？"
    entity="lexical_database"
  >
    <template #toolbar="{params,handleSearch}">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        :placeholder="$t('chineseTip')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <city-selector v-model="params.regionID" @update="handleSearch" />
      <handshape-selector v-model="params.handshapeID" @update="handleSearch" />
      <handshape-selector
        v-model="params.leftHandshapeID"
        orientation="left"
        @update="handleSearch"
      />
      <handshape-selector
        v-model="params.rightHandshapeID"
        orientation="right"
        @update="handleSearch"
      />
      <morphemes-input v-model="params.morpheme" @update="handleSearch" @enter="handleSearch" />
    </template>
    <template #chinese="{row}">
      <span
        class="word-sup"
        v-html="$options.filters.addNumberSup(lexicons[row.lexiconID].chinese)"
      />
    </template>

    <template #morpheme="{row}">
      <span v-if="row.morpheme.length > 0">
        <el-tag
          v-for="value in row.morpheme"
          :key="value"
          :disable-transitions="true"
          class="tags word-sup"
          size="small"
          v-html="$options.filters.addNumberSup(value)"
        ></el-tag>
      </span>
      <el-tag v-else :disable-transitions="true" size="small" type="info">{{ $t("NoData") }}</el-tag>
    </template>

    <template #action="{row,handleDeleteItem}">
      <el-button
        type="danger"
        size="mini"
        plain
        @click.stop="handleDeleteItem(row.id)"
      >{{ $t("Delete") }}</el-button>
    </template>
  </list-view>
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
import { mapGetters } from "vuex";
import ListView from "@/components/ListView";
import {
  GetLexicalVideosList,
  CreateLexicalVideo,
  UpdateLexicalVideo,
  DeleteLexicalVideo
} from "@/api/videos";
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import CitySelector from "@/components/form/CitySelector";
import HandshapeSelector from "@/components/form/HandshapeSelector";
import MorphemesInput from "@/components/form/MorphemesInput";
export default {
  name: "LexicalVideoList",
  components: {
    ListView,
    CitySelector,
    HandshapeSelector,
    MorphemesInput
  },
  data() {
    return {
      LexicalVideoForm,
      GetLexicalVideosList,
      CreateLexicalVideo,
      UpdateLexicalVideo,
      DeleteLexicalVideo,
      columns: [
        {
          prop: "initial",
          label: this.$t("Initial"),
          width: "80px",
          filters: [],
          formatter: row => this.lexicons[row.lexiconID].initial
        },
        {
          prop: "region",
          label: this.$t("Region"),
          width: "200px",
          formatter: row =>
            this.$options.filters.getRegionName(
              this.performers[row.performerID].regionID
            )
        },
        {
          prop: "gender",
          label: this.$t("Gender"),
          width: "100px",
          formatter: row =>
            this.$t(
              this.genderTypes[this.performers[row.performerID].gender].name
            )
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "120px",
          formatter: row => this.performers[row.performerID].name
        },
        {
          slot: "chinese",
          label: this.$t("Chinese"),
          width: "150px",
          hideOverflow: true
        },
        {
          prop: "wordFormation",
          label: this.$t("WordFormation"),
          width: "150px",
          formatter: row =>
            row.wordFormation
              ? this.$t(this.wordFormations[row.wordFormation].name)
              : this.$t("NoData")
        },
        {
          slot: "morpheme",
          label: this.$t("Morpheme"),
          width: "200px",
          hideOverflow: true
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "100px"
        }
      ]
    };
  },
  computed: {
    ...mapGetters([
      "genderTypes",
      "wordFormations",
      "wordInitial",
      "partOfSpeech",
      "performers",
      "lexicons",
      "handshapes"
    ])
  },
  async created() {
    await this.$store.dispatch("data/getPerformers");
    await this.$store.dispatch("data/getLexicons");
    await this.$store.dispatch("data/getHandshapes");
    this.wordInitial.map(item => {
      this.columns[0].filters.push({
        text: item,
        value: item
      });
    });
  },
  methods: {
    exportConfig(item) {
      const leftHandshapes = [];
      const rightHandshapes = [];
      const partOfSpeech = [];
      this.lexicons[item.lexiconID].pos.map(k => {
        const v = this.$t(this.partOfSpeech[k].name);
        partOfSpeech.push(v);
      });
      item.leftHandshapesID.map(id => {
        const handshape = this.handshapes[id];
        leftHandshapes.push(handshape.name);
      });
      item.rightHandshapesID.map(id => {
        const handshape = this.handshapes[id];
        rightHandshapes.push(handshape.name);
      });
      return {
        [this.$t("CreatedAt")]: new Date(item.createdAt),
        [this.$t("UpdatedAt")]: new Date(item.updatedAt),
        [this.$t("Initial")]: this.lexicons[item.lexiconID].initial,
        [this.$t("Chinese")]: this.lexicons[item.lexiconID].chinese,
        [this.$t("English")]: this.lexicons[item.lexiconID].english,
        [this.$t("PoS")]: partOfSpeech.join("/"),
        [this.$t("WordFormation")]: item.wordFormation
          ? this.$t(this.wordFormations[item.wordFormation].name)
          : this.$t("NoData"),
        [this.$t("Morpheme")]: item.morpheme.join(","),
        [this.$t("LeftHandshape")]: leftHandshapes.join(","),
        [this.$t("RightHandshape")]: rightHandshapes.join(","),
        [this.$t("Name")]: this.performers[item.performerID].name,
        [this.$t("Gender")]: this.$t(
          this.genderTypes[this.performers[item.performerID].gender].name
        ),
        [this.$t("Region")]: this.$options.filters.getRegionName(
          this.performers[item.performerID].regionID
        )
      };
    }
  }
};
</script>

<style lang="scss" scoped>
.tags {
  margin: 3px;
}
</style>