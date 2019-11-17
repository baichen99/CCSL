<template>
  <el-carousel class="carousel-card" height="100%" trigger="click">
    <el-carousel-item v-for="item in carousels" :key="item.id">
      <el-image style="width: 100%; height: 100%" :src="settings.publicURL + item.image" fit="fit" />
      <div class="carousel-title">{{ item.title }}</div>
    </el-carousel-item>
  </el-carousel>
</template>

<script>
import { mapGetters } from "vuex";
import { GetCarouselsList } from "@/api/carousel";
export default {
  name: "CarouselCard",
  data: () => ({
    carousels: []
  }),
  computed: {
    ...mapGetters(["settings"])
  },
  created() {
    this.getData();
  },
  methods: {
    getData() {
      GetCarouselsList({ order: "desc", state: "published" }).then(res => {
        this.carousels = res.data;
      });
    }
  }
};
</script>



<style lang="scss" scoped>
.carousel-card {
  width: 1000px;
  height: 562px;
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
