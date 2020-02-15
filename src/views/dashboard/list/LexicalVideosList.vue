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
    <template #toolbar-search="{params,handleSearch}">
      <search-input v-model="params.name" :placeholder="$t('tipZh')" @update="handleSearch" />
      <city-selector
        v-model="params.regionID"
        class="city-selector"
        size="mini"
        @update="handleSearch"
      />
      <handshape-selector v-model="params.handshapeID" size="mini" @update="handleSearch" />
      <handshape-selector
        v-model="params.leftHandshapeID"
        orientation="left"
        size="mini"
        @update="handleSearch"
      />
      <handshape-selector
        v-model="params.rightHandshapeID"
        orientation="right"
        size="mini"
        @update="handleSearch"
      />
      <search-input
        v-model="params.morpheme"
        :placeholder="$t('tipMorpheme')"
        @update="handleSearch"
      />
    </template>
    <template #chinese="{row}">
      <span
        v-if="lexicons[row.lexiconID]"
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
  </list-view>
</template>

<i18n>
{
  "zh-CN": {
    "tipZh": "请输入中文",
    "tipMorpheme": "请输入语素"
  },
  "en-US": {
    "tipZh": "Input Chinese",
    "tipMorpheme": "Input Morpheme"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import ListView from "@/components/ListView";
import SearchInput from "@/components/form/SearchInput";
import {
  GetLexicalVideosList,
  CreateLexicalVideo,
  UpdateLexicalVideo,
  DeleteLexicalVideo
} from "@/api/videos";
import LexicalVideoForm from "@/views/dashboard/form/LexicalVideoForm";
import CitySelector from "@/components/form/CitySelector";
import HandshapeSelector from "@/components/form/HandshapeSelector";
export default {
  name: "LexicalVideoList",
  components: {
    ListView,
    CitySelector,
    HandshapeSelector,
    SearchInput
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
          formatter: row =>
            this.lexicons[row.lexiconID]
              ? this.lexicons[row.lexiconID].initial
              : ""
        },
        {
          prop: "region",
          label: this.$t("Region"),
          width: "200px",
          formatter: row =>
            this.performers[row.performerID]
              ? this.$options.filters.getRegionName(
                  this.performers[row.performerID].regionID
                )
              : ""
        },
        {
          prop: "gender",
          label: this.$t("Gender"),
          width: "100px",
          filters: [
            {
              text: this.$t("Male"),
              value: "M"
            },
            {
              text: this.$t("Female"),
              value: "F"
            }
          ],
          formatter: row =>
            this.performers[row.performerID]
              ? this.$t(
                  this.genderTypes[this.performers[row.performerID].gender].name
                )
              : ""
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "120px",
          formatter: row =>
            this.performers[row.performerID]
              ? this.performers[row.performerID].name
              : ""
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
          filters: [
            {
              text: this.$t("CompoundWord"),
              value: "compound"
            },
            {
              text: this.$t("SimpleWord"),
              value: "simple"
            }
          ],
          formatter: row =>
            row.wordFormation
              ? this.$t(this.wordFormations[row.wordFormation].name)
              : this.$t("NoData")
        },
        {
          slot: "morpheme",
          label: this.$t("Morpheme"),
          width: "150px",
          hideOverflow: true
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

.el-select,
.el-input,
.city-selector {
  margin-right: 5px;
  margin-top: 2px;
}

.city-selector {
  width: 200px;
}
</style>