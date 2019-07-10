<template>
  <div class="list-card">
    <el-card shadow="hover">
      <div slot="header">
        <span :class="['iconfont','icon-'+(icon+1)] " />
        <span class="list-title">{{ list.title }}</span>
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
        <i class="dot">&bull;</i>
        <el-link class="list-desc">
          {{ item.title }}
        </el-link>
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
    },
    icon: {
      type: Number,
      default: 0
    }
  }
};
</script>



<style lang="scss" scoped>
.list-card {
  .el-card {
    margin: 1rem 0;
    width: 430px;
    .list-title {
      font-weight: bolder;
      color: #2363c3;
    }
    .list-more {
      float: right;
      padding: 3px 0;
    }
    .iconfont {
      color: #2363c3;
      padding-right: 1.2rem;
    }

    .list-content {
      padding: 0.3rem;
    }
    .dot {
      color: #2363c3;
      padding: 0.2rem;
    }
    .list-desc {
      display: inline-block;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      max-width: 260px;
      font-size: 1rem;
      // font-weight: normal;

      // @media only screen and (max-device-width: 480px) {
      //   width: 50vw;
      // }
    }

    .list-time {
      float: right;
      color: #a0a0a0;
      height: 100%;
      padding-top: 0.2rem;
    }
  }
}
</style>
