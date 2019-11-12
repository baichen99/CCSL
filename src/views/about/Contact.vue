<template>
  <div class="contact flex-column">
    <h2 class="title">{{ $t("ContactMenu") }}</h2>
    <el-row :gutter="20" class="container">
      <el-col :span="12">
        <el-card shadow="hover">
          <el-table :data="tableData">
            <el-table-column prop="key" width="60" />
            <el-table-column prop="value" />
            <br />
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <h3>{{ $t("wechat") }}</h3>
          <img src="@/assets/wechat.png" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "wechat": "微信公众号"
  },
  "en-US": {
    "wechat": "WeChat Official Accounts"
  }
}
</i18n>

<script>
export default {
  name: "Contact",
  data: () => ({
    tableData: [
      { key: "网址", value: "ccsl.shu.edu.cn" },
      { key: "地址", value: "上海市宝山区南陈路333号上海大学东区文学院105室" },
      { key: "邮编", value: "200443" },
      { key: "邮箱", value: "ccsl@oa.shu.edu.cn" },
      { key: "电话", value: "021-66133312" }
    ]
  })
};
</script>


<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.contact {
  align-items: center;
  .title {
    color: $--color-primary;
    margin: 20px;
  }

  .container {
    margin: 1rem auto;
  }
  .el-card {
    width: 450px;
    height: 350px;
    text-align: center;
  }
  
  .el-table {
    margin-bottom: 2rem;
  }
}
</style>
