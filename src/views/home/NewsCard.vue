<template>

      <el-card shadow="hover">
        <span>{{ key }}</span>
        <el-button
          class="news-button-more"
          type="text"
          @click="getMoreNews(key)"
        >
          更多
        </el-button>
      </el-card>

</template>


<style lang="scss" scoped>
.news {
  .el-card {
    margin: 1rem;

    .news-button-more {
      float: right;
      padding: 3px 0;
    }

    .content {
      padding-top: 0.2rem;
    }

    .desc {
      cursor: pointer;
      display: inline-block;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      width: 20vw;

      @media only screen and (max-device-width: 480px) {
        width: 50vw;
      }
    }

    .time {
      float: right;
      color: #a0a0a0;
    }
  }
}
</style>
