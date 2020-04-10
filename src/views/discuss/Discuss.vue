<template>
  <div class="container">
    <topBar v-model="params" @search="searchByButton" @createPost="createPost" />
    <SearchResult 
      v-model="params.page"
      :posts="posts" 
      :total="total"
      @page-change="changePage" 
    />
  </div>
</template>

<script>
import topBar from "./TopBar"
import SearchResult from "./SearchResult"
import { GetPostsList } from '@/api/posts'
export default {
  components: {
    topBar,
    SearchResult,
  },
  data() {
    return {
      total: 0,
      loading: false,
      params: {
        title: "",
        page: 1,
        limit: 4,
        orderBy: "posts.created_at",
        order: "desc"
      },
      posts: []
    }
  },
  created() {
    this.$nextTick(() => {
      this.params.page = 1;
      this.getData();
    })
  },
  methods: {
    async getData() {
      this.loading = true;
      try {
        const res = await GetPostsList(this.params);
        if (res.total === 0) {
          this.$notify.info({
            title: this.$t("PostNotFound"),
            duration: 2000
          });
        }
        this.posts = res.data;
        this.params.page = res.page;
        this.params.limit = res.limit;
        this.total = res.total;
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
        this.params.title = ""
      }
    },
    createPost() {
      this.params.page = 1;
      this.getData();
    },
    changePage() {
      this.getData()
    },
    searchByButton() {
      this.params.page = 1;
      this.getData();
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
