<template>
  <div class="home">
    <carousel-card :carousels="carousels" />
    <div class="list-container">
      <el-row
        class="list"
        :gutter="40"
      >
        <el-col
          v-for="(item, index) in list"
          :key="index"
          :md="12"
          :sm="24"
        >
          <list-card :list="item" />
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import CarouselCard from "@/views/home/CarouselCard.vue";
import ListCard from "@/views/home/ListCard.vue";
export default {
  name: "Home",
  components: {
    CarouselCard,
    ListCard
  },
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
    ],
    list: [
      {
        title: "新闻动态",
        content: [
          { title: "中国手语研究论坛在我校召开", time: "2018-01-01" },
          { title: "上海大学中国手语及聋人研究中心成立仪式暨中国手语语言学研讨会在上海大学举行", time: "2018-01-01" },
          { title: "国家通用手语比对语料库和监督测查项目启动研讨会在上海大学召开", time: "2018-01-01" },
          { title: "中国手语及聋人研究中心成功举办手语语料库标注研讨会", time: "2018-01-01" },
          { title: "面向人工智能应用的手语语料库建设研讨会在上海大学召开", time: "2018-01-01" }
        ]
      },
      {
        title: "学术活动",
        content: [
          { title: "中国手语研究论坛在我校召开", time: "2018-01-01" },
          { title: "上海大学中国手语及聋人研究中心成立仪式暨中国手语语言学研讨会在上海大学举行", time: "2018-01-01" },
          { title: "国家通用手语比对语料库和监督测查项目启动研讨会在上海大学召开", time: "2018-01-01" },
          { title: "中国手语及聋人研究中心成功举办手语语料库标注研讨会", time: "2018-01-01" },
          { title: "面向人工智能应用的手语语料库建设研讨会在上海大学召开", time: "2018-01-01" }
        ]
      },
      {
        title: "通知公告",
        content: [
          { title: "中国手语研究论坛在我校召开", time: "2018-01-01" },
          { title: "上海大学中国手语及聋人研究中心成立仪式暨中国手语语言学研讨会在上海大学举行", time: "2018-01-01" },
          { title: "国家通用手语比对语料库和监督测查项目启动研讨会在上海大学召开", time: "2018-01-01" },
          { title: "中国手语及聋人研究中心成功举办手语语料库标注研讨会", time: "2018-01-01" },
          { title: "面向人工智能应用的手语语料库建设研讨会在上海大学召开", time: "2018-01-01" }
        ]
      },
      {
        title: "科研动态",
        content: [
          { title: "中国手语研究论坛在我校召开", time: "2018-01-01" },
          { title: "上海大学中国手语及聋人研究中心成立仪式暨中国手语语言学研讨会在上海大学举行", time: "2018-01-01" },
          { title: "国家通用手语比对语料库和监督测查项目启动研讨会在上海大学召开", time: "2018-01-01" },
          { title: "中国手语及聋人研究中心成功举办手语语料库标注研讨会", time: "2018-01-01" },
          { title: "面向人工智能应用的手语语料库建设研讨会在上海大学召开", time: "2018-01-01" }
        ]
      }
    ]
  })
};
</script>


<style lang="scss" scoped>
.home {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .list-container {
    width: 900px;
    margin: 1rem auto;
  }
}
</style>
