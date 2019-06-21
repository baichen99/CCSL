<template>
  <div class="list-card">
    <el-card shadow="hover">
      <div slot="header">
        <span>{{ list.title }}</span>
        <el-link
          class="list-more"
          type="primary"
        >
          更多
        </el-link>
      </div>
      <div
        v-for="item in list.content"
        :key="item.title"
        class="list-content"
      >
        <div class="list-desc"><i class="dot">&bull;</i> {{ item.title }}</div>
        <div class="list-time">{{ item.time }}</div>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "ListCard",
  props: {
    list: {
      type: Object,
      default: () => ({
        title: "",
        content: []
      })
    }
  }
};
</script>



<style lang="scss" scoped>
.list-card {
  .el-card {
    margin: 10px 0;
    width: 430px;
    .list-more {
      float: right;
      padding: 3px 0;
    }

    .list-content {
      padding-top: 0.2rem;
    }

    .list-desc {
      cursor: pointer;
      display: inline-block;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      width: 280px;
      .dot {
        color: #2363c3;
        padding: 0 5px;
      }

      // @media only screen and (max-device-width: 480px) {
      //   width: 50vw;
      // }
    }

    .list-time {
      float: right;
      color: #a0a0a0;
    }
  }
}
</style>
