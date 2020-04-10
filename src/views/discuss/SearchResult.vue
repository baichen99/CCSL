<template>
  <div class="searchResult">
    <PostCard v-for="post in posts" :key="post.id" :post="post" />
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
import PostCard from "./PostCard"
export default {
  components: {
    PostCard,
  },
  model: {
    prop: 'page',
    event: 'update:page',
  },
  props: {
    posts: {
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
    .PostCard {
      margin-top: 10px;
    }
    .el-pagination {
      margin-top: 20px;
    }
  }
</style>