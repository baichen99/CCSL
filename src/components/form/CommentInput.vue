<template>
  <div class="CommentInput">
    <el-input v-model="commentText" :placeholder="$t('commentPlaceholder')" type="textarea" autosize required /> 
    <el-button type="primary" size="small" @click="$emit('commentClick')">评论</el-button>
  </div>
</template>


<i18n>
{
  "zh-CN": {
    "commentPlaceholder": "增加评论"
  },
  "en-US": {
    "commentPlaceholder": "Add a comment"
  }
}
</i18n>

<script>
export default {
  model: {
    prop: "comment",
    event: "update:comment"
  },
  props: {
    comment: {
      type: String,
      default: "",
    }
  },
  data() {
    return {
    }
  },
  computed: {
    commentText: {
      get() {
        return this.comment;
      },
      set(val) {
        this.$emit("update:comment", val);
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .CommentInput {
    width: 70%;
    display: flex;
    .el-button {
      margin-left: 20px;
      align-self: flex-end;
    }
  }
</style>