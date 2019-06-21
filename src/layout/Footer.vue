<template>
  <footer>
    <div class="iconfont icon-ccsl" />
    <div>沪ICP备 09014157号</div>
    <div>中国手语及聋人研究中心</div>
    <div>邮箱：ccsl@oa.shu.edu.cn &nbsp; 邮编：200444</div>
    <div>Copyright &copy; 2018-2019 Shanghai University</div>
    <div>地址：上海市宝山区南陈路333号上海大学东区文学院105室</div>
  </footer>
</template>


<style lang="scss" scoped>
footer {
  padding: 20px;
  font-size: 14px;
  height: 240px;
  line-height: 24px;
  background-image: linear-gradient(
    to top,
    darken(#2363c3, 20%) 0%,
    #2363c3 100%
  );
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;

  div {
    width: 100%;
    text-align: center;
  }

  .icon-ccsl {
    font-size: 50px;
    padding: 20px;
  }
}
</style>
