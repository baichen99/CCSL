<template>
  <footer>
    <svg-icon icon-class="logo" class="logo" />
    <div>沪ICP备 09014157号</div>
    <div>中国手语及聋人研究中心</div>
    <div>邮箱：<a href="mailto:ccsl@oa.shu.edu.cn">ccsl@oa.shu.edu.cn</a> &nbsp; 邮编：200444</div>
    <div>Copyright &copy; 2018-2019 Shanghai University</div>
    <div>地址：上海市宝山区南陈路333号上海大学东区文学院105室</div>
  </footer>
</template>


<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
footer {
  padding: 20px;
  font-size: 14px;
  min-height: 240px;
  line-height: 24px;
  background-image: linear-gradient(
    to top,
    darken($--color-primary, 20%) 0%,
    $--color-primary 100%
  );
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;

  div {
    width: 100%;
    text-align: center;
  }

  .logo {
    font-size: 100px;
    fill: #fff !important;
  }
}
</style>
