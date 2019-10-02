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
        <el-option
          label="中文"
          value="chinese"
        />
        <el-option
          label="English"
          value="english"
        />
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
      <el-card
        v-if="advancedSearch && showAdvance"
        shadow="never"
        class="advance"
      >
        <div class="search-options">
          <el-select
            v-model="searchParams.gender"
            clearable
            placeholder="请选择性别"
          >
            <el-option
              label="男"
              value="男"
            />
            <el-option
              label="女"
              value="女"
            />
          </el-select>

          <sign-selector
            v-model="searchParams.leftSign"
            orientation="left"
          />
          <sign-selector
            v-model="searchParams.rightSign"
            orientation="right"
          />

          <el-select
            v-model="searchParams.pos"
            clearable
            placeholder="请选择词性"
          >
            <el-option
              v-for="(item,index) in $store.getters.wordTypes"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
          <city-selector v-model="searchParams.region" />
          <el-select
            v-model="searchParams.constructType"
            clearable
            placeholder="请选择构词方式"
          >
            <el-option
              label="单纯词"
              value="单纯词"
            />
            <el-option
              label="复合词"
              value="复合词"
            />
          </el-select>
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
export default {
  name: "VideoSearchInput",
  components: {
    CitySelector,
    SignSelector
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
  }
}
</style>