<template>
  <div>
    <el-input v-model="keyword" class="options" clearable @keyup.enter.native="search">
      <el-select v-if="showAdvance" slot="prepend" v-model="keywordType" style="width:110px">
        <el-option label="中文" value="zh-CN" />
        <el-option label="English" value="en-US" />
      </el-select>
      <i
        v-show="!advancedSearch && showAdvance"
        slot="prefix"
        class="el-input__icon el-icon-setting"
        style="cursor:pointer;"
        @click="toggleAdvance"
      />
      <i
        v-show="advancedSearch && showAdvance"
        slot="prefix"
        class="el-input__icon el-icon-s-tools"
        style="cursor:pointer;"
        @click="toggleAdvance"
      />

      <i
        slot="suffix"
        class="el-input__icon el-icon-search"
        style="cursor:pointer;"
        @click="search"
      />
    </el-input>

    <el-collapse-transition>
      <el-card v-show="advancedSearch && showAdvance" shadow="never" class="advance">
        <div class="search-options">
          <handshape-selector
            v-model="searchParams.leftHandshapeID"
            size="small"
            orientation="left"
          />
          <handshape-selector
            v-model="searchParams.rightHandshapeID"
            size="small"
            orientation="right"
          />
          <simple-selector v-model="searchParams.gender" size="small" :options="genderTypes" />
          <simple-selector
            v-model="searchParams.wordFormation"
            size="small"
            :options="wordFormations"
          />
          <simple-selector v-model="searchParams.pos" size="small" :options="partOfSpeech" />
          <search-input
            v-model="searchParams.morpheme"
            size="small"
            :placeholder="$t('morphemesTip')"
          />
          <city-selector v-model="searchParams.regionID" size="small" />
        </div>
      </el-card>
    </el-collapse-transition>
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "morphemesTip": "请输入构词语素"
  },
  "en-US": {
    "morphemesTip": "Input morphemes(Chinese)"
  }
}
</i18n>


<script>
import { mapGetters } from "vuex";
import SimpleSelector from "@/components/form/SimpleSelector.vue";
import HandshapeSelector from "@/components/form/HandshapeSelector.vue";
import CitySelector from "@/components/form/CitySelector.vue";
import SearchInput from "@/components/form/SearchInput.vue";

export default {
  name: "VideoSearchInput",
  components: {
    CitySelector,
    HandshapeSelector,
    SimpleSelector,
    SearchInput
  },
  model: {
    prop: "params",
    event: "update"
  },
  props: {
    params: {
      type: Object,
      default: () => ({})
    },
    showAdvance: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      advancedSearch: false,
      keyword: "",
      keywordType: "zh-CN"
    };
  },
  computed: {
    ...mapGetters(["wordFormations", "genderTypes", "partOfSpeech"]),
    searchParams: {
      get() {
        return this.params;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  created() {
    this.keywordType = this.$i18n.locale;
  },
  methods: {
    search() {
      // Reset side panel search parameters
      this.searchParams.word = "";
      this.searchParams.handshape = "";
      this.searchParams.chinese = "";
      this.searchParams.english = "";
      if (this.keywordType === "en-US") {
        this.searchParams.english = this.keyword;
      } else {
        this.searchParams.chinese = this.keyword;
      }
      this.$emit("search");
    },
    toggleAdvance() {
      this.advancedSearch = !this.advancedSearch;
    }
  }
};
</script>


<style lang="scss" scoped>
.advance {
  .search-options {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-content: center;
    .el-input,
    .el-cascader,
    .el-select {
      margin: 5px 5px;
    }
    .el-cascader {
      width: 200px;
    }
    .el-input,
    .el-select {
      width: 150px;
    }
  }
}
</style>