<template>
  <div class="team">
    <h1 class="title">研究团队</h1>
    <div class="list-container">
      <el-table
        class="table"
        :data="EMPList"
        stripe
      >
        <el-table-column
          prop="cls"
          label="类别"
          align="center"
          width="200"
        />
        <el-table-column
          prop="name"
          label="姓名"
          align="center"
          width="150"
        />
        <el-table-column
          prop="unit"
          label="单位"
          align="center"
        />

      </el-table>
    </div>
    <div class="list-container">
      <el-table
        :data="noUnitEMPList"
        stripe
      >
        <el-table-column
          prop="cls"
          label="类别"
          align="center"
        />
        <el-table-column
          prop="name"
          label="姓名"
          align="center"
        />
      </el-table>
    </div>
  </div>
</template>


<script>
export default {
  name: "Team",
  data: () => ({
    EMPList: [
      { cls: "高级顾问", name: "戴浩一", unit: "台湾中正大学教授，博士" },
      { cls: "高级顾问", name: "龚群虎", unit: "复旦大学中文系教授，博士" },
      {
        cls: "高级顾问",
        name: "张吉生",
        unit: "华东师范大学外语学院教授/上海大学兼职教授，博士"
      },
      {
        cls: "研究员",
        name: "倪  兰",
        unit: "上海大学文学院副院长、中心负责人（手语语言学），博士"
      },
      {
        cls: "研究员",
        name: "陈雅清",
        unit: "上海大学文学院博士后（手语语言学），博士"
      },
      {
        cls: "研究员",
        name: "方昱春",
        unit: "上海大学计算机工程与科学学院副教授（图像识别），博士"
      },
      {
        cls: "研究员",
        name: "李颖洁",
        unit: "上海大学通信学院教授（人工智能），博士"
      },
      {
        cls: "研究员",
        name: "华红琴",
        unit: "上海大学社会学院副教授（社工专业、心理学），博士"
      },
      {
        cls: "研究员",
        name: "雷红波",
        unit: "上海大学文学院讲师（社会语言学），博士"
      },
      {
        cls: "研究员",
        name: "杨军辉（聋）",
        unit: "英国中央兰开夏大学高级讲师，博士"
      },
      { cls: "研究员", name: "郑  璇（聋）", unit: "重庆师范大学教授，博士" },
      {
        cls: "研究员",
        name: "仰国维（聋）",
        unit: "河南财经政法大学副教授，硕士"
      },
      {
        cls: "研究员",
        name: "赵伟时",
        unit: "上海市残疾人康复中心副主任，学士"
      },
      { cls: "研究员", name: "洪  泽", unit: "上海市聋协主席，学士" },
      {
        cls: "助理研究员",
        name: "养丹阳",
        unit: "上海大学文学院助理研究员，硕士"
      }
    ],
    noUnitEMPList: [
      { cls: "手语翻译员", name: "唐文妍（硕士）" },
      { cls: "手语翻译员", name: "寇辰珠（学士）" },
      { cls: "聋人研究助理", name: "朱莉娜" },
      { cls: "聋人研究助理", name: "何立峰" },
      { cls: "聋人研究助理", name: "赵明华" },
      { cls: "聋人研究助理", name: "倪颖杰" },
      { cls: "聋人研究助理", name: "刘梦琪" },
      { cls: "聋人研究助理", name: "沈  颖" },
      { cls: "聋人研究助理", name: "顾耀勇" },
      { cls: "聋人研究助理", name: "陈淡柳" },
      { cls: "研究生", name: "孙  玲" },
      { cls: "研究生", name: "和子晴" },
      { cls: "财务人员", name: "王玲玲" }
    ]
  })
};
</script>


<style lang="scss" scoped>
.team {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .title {
    margin: 1rem;
  }
  .list-container {
    width: 900px;
    margin: 1rem auto;
    .el-table {
      border-radius: 5px;
    }
  }
}
</style>
