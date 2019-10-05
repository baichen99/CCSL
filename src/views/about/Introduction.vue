<template>
  <div class="introduction">
    <h1 class="title">中心介绍</h1>
    <el-card>
      <div>
        <p>
          2017年5月26日上海大学文学院成立中国手语及聋人研究中心。该中心致力于开展中国手语研究与推广工作，打造中国手语的跨地区合作研究、协同工作平台，建立中国手语研究、应用领域的合作机制，积极推动科学研究、应用推广、学科建设和人才培养等各项工作，以合作促发展，共同推进中国手语的调查、研究、应用等工作，为保障残疾人的语言文字权利和融合发展，为听障人群更好融入社会搭建信息沟通平台，努力探寻进一步改善听障人群平等参与社会生活的环境和提高残疾人发展水平的途径，建设信息无障碍的社会环境。
        </p>
        <p>
          该中心依托上海大学综合大学多学科的优势，在语言学、心理学、社会学、计算机、人工智能、脑科学等学科领域进行手语的研究和探索，并积极进行手语的应用研究。
        </p>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "Introduction"
};
</script>


<style lang="scss" scoped>
.introduction {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .el-card {
    width: 900px;
    margin: 1rem auto;
    padding: 3rem 2rem;
    p {
      text-indent: 2rem;
    }
  }
}
.title {
  margin: 1rem;
}
</style>
