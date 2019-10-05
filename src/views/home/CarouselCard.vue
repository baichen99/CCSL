<template>
  <el-carousel
    class="carousel-card"
    height="100%"
    trigger="click"
  >
    <el-carousel-item
      v-for="item in carousels"
      :key="item.id"
    >
      <img
        :src="'https://ccsl.shu.edu.cn/public/'+item.image"
        :alt="item.id"
        width="100%"
        height="100%"
      >
      <div class="carousel-title">{{ item.title }}
      </div>
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
  name: "CarouselCard",
  props: {
    carousels: {
      type: Array,
      default: () => []
    }
  },
  data: () => ({})
};
</script>



<style lang="scss" scoped>
.carousel-card {
  width: 900px;
  height: 506px;
  border-radius: 5px;
  margin: 1rem;

  // @media only screen and (max-device-width: 480px) {
  //   height: 200px;
  // }

  .carousel-title {
    border-radius: 0 0 10px 10px;
    position: relative;
    background-color: rgba(0, 0, 0, 0.7);
    color: #ffffff;
    top: -3.5rem;
    height: 3.5rem;
    text-align: center;
    padding-top: 0.5rem;

    // @media only screen and (max-device-width: 480px) {
    //   font-size: 14px;
    //   padding-top: 0.5rem;
    //   height: 2.2rem;
    //   top: -3rem;
    // }
  }
}
</style>
