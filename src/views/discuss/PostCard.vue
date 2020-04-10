<template>
  <el-card v-if="JSON.stringify(post) != '{}'" class="PostCard">
    <div class="topbar">
      <div>
        <h2 class="title" @click="showDetail(post.id)">{{ post.title }}</h2>
        <div class="user-info">{{ $t('Publisher') }}: {{ post.creator.name }}</div>
      </div>
      <div>
        <el-button v-show="showDelete" type="danger" plain @click="deletePost">{{ $t('Delete') }}</el-button>
      </div>
    </div>

    <div class="content" v-html="post.content"></div>
  </el-card>
</template>

<i18n>
{
  "zh-CN": {
    "reply": "回复",
    "DeletePost": "删除帖子"
  },
  "en-US": {
    "reply": "Reply",
    "DeletePost": "Deleted Post"
  }
}
</i18n>

<script>
import { DeletePost } from '@/api/posts'
export default {
  name:'PostCard',
  props: {
    post: {
      type: Object,
      default: ()=>({}),
    }
  },
  data() {
    return {
      comment: ""
    }
  },
  computed: {
    IsSuper() {
      let roles = this.$store.getters.roles;
      if (roles.indexOf('super') > -1) {
        return true;
      }
      return false;
    },
    showDelete() {
      let id = this.$store.getters.id;
      return this.post.creator.id == id || this.IsSuper;
    }
  },
  methods: {
    showDetail() {
      this.$router.push(`/learning-plartform/posts/${this.post.id}`).catch(err => {err})
    },
    deletePost() {
      DeletePost(this.post.id).
        then(() => {
          if(this.$route.path != '/learning-plartform/posts') {
            this.$router.push('/learning-plartform/posts')
          } else {
            this.$emit('delete');
          }
          this.$notify.success({
            title: this.$t("DeletePost"),
            duration: 2000
          })
        }).
        catch(err => {
          console.error(err);
        })
    }
  }
}
</script>

<style lang="scss" scoped>
  .PostCard {
    width: 80%;
    border-bottom: 1px solid #e6e6e6;
    .topbar {
      display: flex;
      justify-content: space-between;
      .el-button {
        height: 40px;
      }
      div {
        .user-info {
          font-size: 15px;
          font-weight: 600;
          color: #444;
        }
        .title {
          cursor: pointer;
          &:hover {
            text-decoration: underline;
            color: royalblue;
          }
        }
      }
    }
    .content {
      padding: 10px;
    }
    .actions {
      display: flex;
      justify-content: flex-end;
    }
  }
</style>