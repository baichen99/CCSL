<template>
  <div class="container">
    <topBar v-model="params" />
    <SearchResult 
      :topics="topics" 
      :total="100"
      :page="1"
      @update:page="changePage" 
    />
  </div>
</template>

<script>
import topBar from "./TopBar"
import SearchResult from "./SearchResult"
export default {
  components: {
    topBar,
    SearchResult,
  },
  data() {
    return {
      params: {
        page: 1,
        limit: 4,
      },
      topics: [
        { id: 1, creator: "用户名称", title: "这是标题", content: "asdadhaksdjasjdklajsdkjaksldjaksdjaksdjalksdjaklsdjalkdsadjaklsjdklasjdaksjdlaksjdklasjdlka" },
        { id: 2, creator: "用户名称", title: "这是标题", content: "asdadhaksdjasjdklajsdkjaksldjaksdjaksdjalksdjaklsdjalkdsadjaklsjdklasjdaksjdlaksjdklasjdlka" },
        { id: 3, creator: "用户名称", title: "这是标题", content: "asdadhaksdjasjdklajsdkjaksldjaksdjaksdjalksdjaklsdjalkdsadjaklsjdklasjdaksjdlaksjdklasjdlka" },
        { id: 4, creator: "用户名称", title: "这是标题", content: "asdadhaksdjasjdklajsdkjaksldjaksdjaksdjalksdjaklsdjalkdsadjaklsjdklasjdaksjdlaksjdklasjdlka" },
        { id: 5, creator: "用户名称", title: "这是标题", content: "asdadhaksdjasjdklajsdkjaksldjaksdjaksdjalksdjaklsdjalkdsadjaklsjdklasjdaksjdlaksjdklasjdlka" },
      ]
    }
  },
  methods: {
    changePage(page) {
      console.log(page);
    }
  }
}
</script>

<style lang="scss" scoped>
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    .searchResult {
      margin-top: 20px;
    }
  }
</style>
