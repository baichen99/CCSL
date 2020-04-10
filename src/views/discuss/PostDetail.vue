<template>
  <div class="container">
    <PostCard :post="post" />
    <CommentInput v-model="comment" @commentClick="handleComment" />
    <ReplyCard v-for="reply in post.replies" :key="reply.id" :reply="reply" @delete="deleteComment" />
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "CreateComment": "增加评论",
    "DeleteComment": "删除评论"
  },
  "en-US": {
    "CreateComment": "Created comment",
    "DeleteComment": "Deleted comment"
  }
}
</i18n>

<script>
import PostCard from '@/views/discuss/PostCard'
import ReplyCard from '@/views/discuss/ReplyCard'
import CommentInput from '@/components/form/CommentInput'
import { GetPost } from '@/api/posts'
import { CreateReply, DeleteReply } from '@/api/replies'
export default {
  name: 'PostDetail',
  components: {
    PostCard,
    ReplyCard,
    CommentInput
  },
  data: () => ({
    comment:"",
    post: {}
  }),
  mounted() {
    this.$nextTick(() => {
      this.getPost();
    })
  },
  methods: {
    getPost() {
      let id = this.$route.params.id;
      GetPost(id).then(res => {
        this.post = res.data;        
      })
    },
    handleComment() {
      CreateReply({
        postID: this.post.id,
        content: this.comment
      }).then(() => {
        this.$notify.success({
          title: this.$t("CreateComment"),
          duration: 2000
        })
      }).catch(err => {
        this.$notify.error({
        title: this.$t(err),
        duration: 2000
        });
      }).finally(() => {
        this.comment = "";
        this.getPost();
      })
    },
    deleteComment(id) {
      DeleteReply(id).then(() => {
        this.$notify.info({
          title: this.$t("DeleteComment"),
          duration: 2000
        });
      }).catch(err => {
        console.error(err);
      }).finally(() => {
        this.getPost()
      })
    }

  },
}
</script>

<style lang="scss" scoped>
  .container {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    .ReplyCard {
      margin-top: 10px;
    }
    .CommentInput {
      margin-top: 20px;
    }
  }
</style>