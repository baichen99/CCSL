<template>
  <footer>
    <svg-icon icon-class="logo" class="logo" />
    <div>{{ $t("icpLicense") }}</div>
    <div>{{ $t("CCSL") }}</div>
    <div>{{ $t("copyright") }} &copy; 2018-2019 {{ $t("SHU") }}</div>
    <div>
      {{ $t("Email") }}:
      <a href="mailto:ccsl@oa.shu.edu.cn">ccsl@oa.shu.edu.cn</a>
      &nbsp; {{ $t("Postcode") }}: 200443
    </div>
    <div>{{ $t("Address") }}: {{ $t("AddressCCSL") }}</div>
  </footer>
</template>

<i18n>
{
  "zh-CN": {
    "icpLicense": "沪ICP备 09014157号",
    "copyright": "版权所有"
  },
  "en-US": {
    "icpLicense": "ICP License: SH ICP 09014157",
    "copyright": "Copyright"
  }
}
</i18n>


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
