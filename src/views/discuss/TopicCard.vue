<template>
  <el-card class="topicItem">
    <div>
      <div class="user-info">{{ topic.creator }}</div>
      <h2 class="title" @click="showDetail(topic.id)">{{ topic.title }}</h2>
    </div>
    <div class="content"> {{ topic.content }} </div>
    <div class="actions">
      <el-button type="primary" plain>回复</el-button>
    </div>
  </el-card>
</template>

<script>
export default {
  props: {
    topic: {
      type: Object,
      default: ()=>({}),
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
  .topicItem {
    width: 80%;
    border-bottom: 1px solid #e6e6e6;
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
    .content {
      padding: 10px;
    }
    .actions {
      display: flex;
      flex-direction: column;
      align-items: flex-end;
    }
  }
</style>