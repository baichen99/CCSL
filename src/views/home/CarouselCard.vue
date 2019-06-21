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
        :src="item.image"
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
  data: () => ({
    carousels: [
      {
        id: 1,
        image: "carousels/carousel1.png",
        title: "国家通用手语比对语料库和监督测查项目启动研讨会在上海大学召开"
      },
      {
        id: 2,
        image: "carousels/carousel2.png",
        title:
          "上海大学中国手语及聋人研究中心成立仪式暨中国手语语言学研讨会在上海大学举行"
      },
      {
        id: 3,
        image: "carousels/carousel6.png",
        title: "中国手语研究论坛在我校召开"
      },
      {
        id: 4,
        image: "carousels/carousel5.png",
        title: "中国手语及聋人研究中心成功举办手语语料库标注研讨会"
      }
    ]
  })
};
</script>



<style lang="scss" scoped>
.carousel-card {
  width: 900px;
  height: 506px;
  border-radius: 10px;
  margin: 1rem;

  // @media only screen and (max-device-width: 480px) {
  //   height: 200px;
  // }

  .carousel-title {
    border-radius: 0 0 10px 10px;
    position: relative;
    background-color: rgba(0, 0, 0, 0.7);
    color: #ffffff;
    top: -3rem;
    height: 3rem;
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
