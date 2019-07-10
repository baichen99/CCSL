<template>
  <div class="contact">
    <h1 class="title">联系方式</h1>
    <el-row :gutter="20" class="container">
      <el-col :span="12">
        <el-card class="card">
            <el-table :show-header="flase" :data="tableData" class="table">
              <el-table-column prop="key" width="60">
              </el-table-column>
              <el-table-column prop="value">
              </el-table-column>
              <br>
            </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="card">
          <h3>微信公众号</h3>
          <img src="../../assets/wechat.png">
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Contact",
  data: () => ({
    tableData:[
      {key:"网址",value:"www.ccsl.shu.edu.cn"},
      {key:"地址",value:"上海市宝山区南陈路333号上海大学东区文学院105室"},
      {key:"邮编",value:"200443"},
      {key:"邮箱",value:"ccsl@oa.shu.edu.cn"},
      {key:"电话",value:"021-66133312"},
    ]
  })
};
</script>


<style lang="scss" scoped>
.contact {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 5px;
  .container {
    margin: 1rem auto;
  }
  .card {
    width: 450px;
    height: 350px;
    text-align: center;
  }
  .table {
    margin-bottom: 2rem 
  }
}
.title {
   margin: 1rem;
}
</style>
