<template>
  <div>
    <el-carousel
      class="img-card"
      height="100%"
    >
      <el-carousel-item
        v-for="item in carousels"
        :key="item.id"
      >
        <img
          :src="'img/'+item.cover"
          :alt="item.id"
          width="100%"
          height="100%"
        >
        <div class="img-title">{{ item.title }}
        </div>

      </el-carousel-item>
    </el-carousel>
    <el-rows
      class="news"
      :gutter="12"
    >
      <el-col
        v-for="(value, key) in news"
        :key="key"
        :md="12"
        :sm="24"
      >
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
      </el-col>
    </el-rows>
  </div>
</template>


<style lang="scss" scoped>
.index {
  margin-bottom: 2rem;

  .img-card {
    height: 450px;
    border-radius: 10px;
    margin: 1rem;

    @media only screen and (max-device-width: 480px) {
      height: 200px;
    }

    .img-title {
      border-radius: 0 0 10px 10px;
      position: relative;
      background-color: rgba(0, 0, 0, 0.7);
      color: #ffffff;
      top: -3.4rem;
      height: 3rem;
      text-align: center;
      padding-top: 0.7rem;

      @media only screen and (max-device-width: 480px) {
        font-size: 14px;
        padding-top: 0.5rem;
        height: 2.2rem;
        top: -3rem;
      }
    }
  }
}
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
