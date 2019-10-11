<template>
  <div>
    <el-input
      v-model="keyword"
      class="options"
      clearable
      placeholder="请输入内容"
      @keyup.enter.native="search"
    >
      <el-select
        v-if="showAdvance"
        slot="prepend"
        v-model="keywordType"
        placeholder="请选择"
        style="width:110px"
      >
        <el-option label="中文" value="chinese" />
        <el-option label="English" value="english" />
      </el-select>
      <i
        v-if="!advancedSearch && showAdvance"
        slot="prefix"
        class="el-input__icon el-icon-setting"
        style="cursor:pointer;"
        @click="toggleAdvance"
      />
      <i
        v-if="advancedSearch && showAdvance"
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
      <el-card v-if="advancedSearch && showAdvance" shadow="never" class="advance">
        <div class="search-options">
          <gender-selector v-model="searchParams.gender" />
          <sign-selector v-model="searchParams.leftSignID" orientation="left" />
          <sign-selector v-model="searchParams.rightSignID" orientation="right" />
          <word-pos-selector v-model="searchParams.pos" />
          <city-selector v-model="searchParams.regionID" />
          <word-construct-selector v-model="searchParams.constructType" />
          <el-input
            v-model="searchParams.constructWords"
            clearable
            style="width:200px"
            placeholder="请输入复合词构成词语"
          />
        </div>
      </el-card>
    </el-collapse-transition>
  </div>
</template>

<script>
import SignSelector from "@/components/form/SignSelector.vue";
import CitySelector from "@/components/form/CitySelector.vue";
import GenderSelector from "@/components/form/GenderSelector.vue";
import WordPosSelector from "@/components/form/WordPosSelector.vue";
import WordConstructSelector from "@/components/form/WordConstructSelector.vue";

export default {
  name: "VideoSearchInput",
  components: {
    CitySelector,
    SignSelector,
    GenderSelector,
    WordPosSelector,
    WordConstructSelector
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
      keywordType: "chinese"
    };
  },
  computed: {
    searchParams: {
      get() {
        return this.params;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  methods: {
    search() {
      // Reset side panel search parameters
      this.searchParams.word = "";
      this.searchParams.sign = "";
      this.searchParams.chinese = "";
      this.searchParams.english = "";
      if (this.keywordType === "english") {
        this.searchParams.english = this.keyword;
      } else {
        this.searchParams.chinese = this.keyword;
      }
      this.$emit("search-clicked");
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
      min-width: 250px;
    }
  }
}
</style>