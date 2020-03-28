<template>
  <div class="searchResult">
    <TopicCard v-for="topic in topics" :key="topic.id" :topic="topic" />
    <el-pagination
      background
      layout="total, prev, pager, next"
      :current-page.sync="pageNumber"
      :total="total"
      :page-size.sync="limit"
      @current-change="$emit('page-change', page)"
    />
  </div>
</template>

<script>
import TopicCard from "./TopicCard"
export default {
  components: {
    TopicCard,
  },
  model: {
    prop: 'page',
    event: 'update',
  },
  props: {
    topics: {
      type: Array,
      default: ()=>([]),
    },
    limit: {
      type: Number,
      default: 5
    },
    page: {
      type: Number,
      default: 1
    },
    total: {
      type: Number,
      default: 1
    }
  },
  data() {
    return {

    }
  },
  computed: {
    pageNumber: {
      get() {
        return this.page;
      },
      set(page) {
        this.$emit("update:page", page);
      }
    }
  },
  created() {
    // get data
  },
  methods: {
    showDetail(id) {
      console.log(id);
    }
  }
}
</script>

<style lang="scss" scoped>
  .searchResult {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    .el-pagination {
      margin-top: 20px;
    }
  }
</style>