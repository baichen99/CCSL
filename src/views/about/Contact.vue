<template>
  <div class="contact flex-column">
    <h2 class="title">{{ $t("ContactMenu") }}</h2>
    <el-row :gutter="20" class="container">
      <el-col :span="12">
        <el-card shadow="hover">
          <el-table :data="tableData">
            <el-table-column prop="key" width="100" />
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
  computed: {
    tableData() {
      return [
        { key: this.$t("Website"), value: "ccsl.shu.edu.cn" },
        {
          key: this.$t("Address"),
          value: this.$t("AddressCCSL")
        },
        { key: this.$t("Postcode"), value: "200443" },
        { key: this.$t("Email"), value: "ccsl@oa.shu.edu.cn" },
        { key: this.$t("Tel"), value: "021-66133312" }
      ];
    }
  }
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

  .el-card {
    width: 450px;
    height: 380px;
    margin: 1rem auto;
    text-align: center;
  }
}
</style>
