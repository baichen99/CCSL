<template>
  <div class="container">
    <div class="topBar">
      <el-input v-model="keyword" :placeholder="searchPlaceholder">
        <i
          slot="suffix"
          class="el-input__icon el-icon-search"
          style="cursor:pointer;"
          @click="search"
        />
      </el-input>
      <div class="btns">
        <el-button type="primary" round @click="addTopic">添加话题</el-button>
      </div>
    </div>
    <el-collapse-transition>
      <TopicCreate v-show="showForm" />
    </el-collapse-transition>
  </div>
</template>

<script>
import TopicCreate from "@/components/form/TopicCreate"
export default {
  components: {
    TopicCreate,
  },
  data() {
    return {
      searchPlaceholder: "搜索话题",
      showForm: false,
      keyword: "",
    }
  },
  methods: {
    addTopic() {
      this.showForm = !this.showForm;
    },
    search() {

    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  .topBar {
      width: 80%;
      height: 60px;
      border-bottom: 1px solid #000;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .el-input {
          width: 30%;
      }
  }
  .TopicCreate {
    margin-top: 20px;
    width: 80%;
  }
}

</style>
