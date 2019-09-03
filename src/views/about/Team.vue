<template>
  <div>
  <div v-show="!ifDetail" class="team">
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
          label="单位及学位"
          align="center"
        />
        <el-table-column
          label="详细信息"
          align="center">
          <template slot-scope="scope">
            <el-button @click="handleDetail(scope.row)" type="text" size="small" v-bind:disabled=!scope.row.detail>查看</el-button>
          </template>
        </el-table-column>

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
        <el-table-column
          label="详细信息"
          align="center">
          <template slot-scope="scope">
            <el-button @click="handleDetail(scope.row)" type="text" size="small" v-bind:disabled=!scope.row.detail>查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
  <div v-show="ifDetail" class="team">
    <h1 class="title">
      详细介绍
    </h1>
    <div class="list-container">
      <el-card>
        <div slot="header" class="clearfix">
          <el-button @click="backToTeam" style="float: left; padding: 3px 0" type="text">
            返回
          </el-button>
          <span style="float: right">
              成员详细信息
          </span>
        </div>
        <div>
          <el-row :gutter="20">
            <el-col :span="6">
              <img :src="imgSrc" alt="" width="160" height="200">
            </el-col>
            <el-col :span="18">
              <h1>{{emp.name}} {{emp.title}}</h1>
              <br>
              <p>{{emp.detail}}</p>
            </el-col>
          </el-row>
        </div>
      </el-card>
    </div>
  </div>
  </div>
</template>


<script>
// import { constants } from 'crypto';
export default {
  name: "Team",
  data: () => ({
    ifDetail: false,
    emp: "",
    imgSrc: "",
    EMPList: [
      { 
        cls: "高级顾问",
        name: "戴浩一", 
        unit: "台湾中正大学教授，博士",
        title: "教授",
        img: "daihaoyi.png",
        detail: "美国印地安那大学语言学博士，相继任教於美国南伊利诺大学(Southern Illinois University)、 美国俄亥俄州立大学(Ohio State University)、台湾中正大学。现为台湾中正大学人文与社会科学研究中心主任、手语语言学台湾研究中心教授，曾两次获得台湾国科会杰出研究奖励。从形式语法到功能语法，再到认知语言学，到现在的手语语言学的研究，戴先生的研究之路是中国语言学研究的探索之路《功能主义与汉语语法》(1994)，《时间顺序和汉语的语序》等著作成为语言学研究的必读书目，其相关著作被广泛引用。"
      },
      { cls: "高级顾问", 
        name: "龚群虎", 
        unit: "复旦大学中文系教授，博士",
        title: "教授" ,
        img: "gongqunhu.png",
        detail: "复旦大学中文系教授，博士生导师。新加坡国立大学博士后，龚群虎教授率先在中国大陆高校引进并开创手语语言学研究，开设手语语言学课程，培养了一批手语语言学硕士、博士研究人才，是大陆知名的手语语言学研究专家。研究兴趣涉及语言学理论、汉语及汉藏系语言的历史比较、中国手语语言学等，近些年尤其关注中国手语语言学的研究，目前承担国家社科基金重大项目《基于汉语和部分少数民族语言的手语语料库建设研究》。"
      },
      {
        cls: "高级顾问",
        name: "张吉生",
        unit: "华东师范大学外语学院教授/上海大学兼职教授，博士",
        title: "教授",
        img: "zhangjisheng.png",
        detail: "华东师范大学外语学院英语系教授，博士生导师。荷兰莱顿大学语言学博士，主要从事生成音系学的研究。他运用当代生成音系学理论‚阐述并揭示了汉语(包括方言)的一些内在音系机制和音系变化规则‚在国内率先进行了手语的音系学研究。近年来‚主持并完成省部级课题两项‚主持国家社科基金《上海手语音系研究》‚在国内外重要学术刊物发表论著40余篇，出版学术专著《上海手语音系》。"
      },
      {
        cls: "研究员",
        name: "倪  兰",
        unit: "上海大学文学院副院长、中心负责人（手语语言学），博士",
        title: "副教授",
        img: "nilan.png",
        detail: "上海大学文学院副院长，副教授，硕士生导师。毕业于复旦大学中文系，文学博士，现任上海大学文学院中国手语及聋人研究中心主任，中国残疾人事业发展研究会残疾人健康管理专业委员会常委委员。主持完成国家社会科学基金青年项目、教育部人文社科项目、中国残疾人联合会、国家语委、上海市语委、上海市残疾人联合会等部门多项委托课题。发表多篇学术论文，出版专著《中国手语动词研究》。"
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
        unit: "英国中央兰开夏大学高级讲师，博士",
        title: "博士",
        img: "yangjunhui.png",
        detail: "1994年毕业于首都师范大学；2000年毕业于美国罗切斯特理工学院Rochester Institute of Technology, Rochester，获教育学硕士学位；2006年毕业于美国加劳德特大学Gallaudet University，获聋人教育博士学位。现于英国中央兰开夏大学University of Central Lancashire任高级讲师。主要学术成果:Yang, J. H. (2008). Sign Language and Oral/Written Language in Deaf Education in China. In C. Plaza-Pust & E. Morales-López (eds), Sign Bilingualism: Language Development, Interaction, and Maintenance in Sign Language Contact Situations. Amsterdam: John Benjamins. p297-331."
      },
      { 
        cls: "研究员",
        name: "郑  璇（聋）",
        unit: "重庆师范大学教授，博士",
        title: "副教授",
        img: "zhenxuan.png",
        detail: "2006年毕业于武汉大学，获得文学硕士学位，2009年毕毕业于复旦大学，获博士学位，为我国大陆首位获得语言学博士学位的聋人，现为重庆师范大学特教系副教授，硕士生导师。目前在美国明尼苏达州圣克劳德州立大学孔子学院任教。出版专著《中国手语如何表达非视觉概念》，知识产权出版社，2011.8。主编教材《手语基础教程》，华东师范大学出版社，2015.9。主持国家社科基金一般项目《西藏地区手语使用状况调查及汉藏手语的语言比较研究》；主持完成国家社科基金青年项目《语言接触对聋人手语发展演变的影响研究》；主持重庆师范大学博士科研启动基金项目《聋人手语的语言学分析及其教育学意义》。"
      },
      {
        cls: "研究员",
        name: "仰国维（聋）",
        unit: "河南财经政法大学副教授，硕士",
        title: "副教授",
        img: "yangguowei.png",
        detail: "河南财经政法大学副教授，毕业于美国奥依华州立大学艺术设计专业，1999年获艺术设计学士学位。2000年在美国罗彻斯特理工学院任助教工作，并进修影视动画及多媒体设计专业。2011年在美国印第安纳波利斯大学，获文学硕士学位。现为中国聋协手语研究推广委员会副主任，长期从事聋教育和手语语言学的研究。"
      },
      {
        cls: "研究员",
        name: "赵伟时",
        unit: "上海市残疾人康复中心副主任，学士",
        img: "zhaoshiwei.png",
        detail: "上海市残疾人康复职业培训中心副主任，国家高级职业指导师，二级心理咨询师，上海市人社局手语翻译员考评员，毕业于华东师范大学大学特殊教育专业，曾任上海闵行启音学校党支部书记、副校长。"
      },
      { 
        cls: "研究员",
        name: "洪  泽", 
        unit: "上海市聋协主席，学士" ,
        img: "hongze.png",
        detail: "上海市聋协主席，毕业于上海大学美术学院，曾参与编写国际手语常用词汇，为迎接上海亚太地区聋人领袖大会，担任国际手语翻译，先后编写了上海地方手语词汇、卫生局医护工作人员手语培训教材、银行窗口服务手语教材等中国手语教学示范。"
      },
      {
        cls: "助理研究员",
        name: "陈雅清",
        unit: "海大学文学院博士后（手语语言学），博士",
        img: "chenyaqing.png",
        detail: "上海大学文学院汉语言文学博士后流动站博士后，毕业于复旦大学中文系手语语言学专业。"
      },
      {
        cls: "助理研究员",
        name: "杨丹阳",
        unit: "上海大学文学院助理研究员，硕士"
      }
    ],
    noUnitEMPList: [
      { 
        cls: "手语翻译员", 
        name: "唐文妍（硕士）", 
        img: "tangwenyan.png",
        detail: "资深手语翻译，电视台手语新闻主持，毕业于华东师范大学特殊教育系心理学专业。" 
      },
      { 
        cls: "手语翻译员", 
        name: "寇辰珠（学士）",
        img: "kouchenzhu.png",
        detail: "高级手语翻译员，长期从事手语教学及手语翻译工作。参与制作《上海市医院志愿者（窗口人员）手语培训手册》等多部行业手语教材；东方卫视《午间30分》及多个区有线台新闻节目手语翻译；《狮子王》《长靴皇后》等多部音乐剧的手语翻译。"
      },
      { 
        cls: "手语翻译员", 
        name: "程  莹（学士）",
        img: "chengying.png",
        detail: "高级手语翻译员，持有手语翻译高级证书，有五年从事手语翻译员的工作经历。"
      },
      { cls: "聋人研究助理", name: "何立峰" },
      { cls: "聋人研究助理", name: "赵明华" },
      { cls: "聋人研究助理", name: "倪颖杰" },
      { cls: "聋人研究助理", name: "朱莉娜" },
      { cls: "聋人研究助理", name: "刘梦琪" },
      { cls: "聋人研究助理", name: "沈  颖" },
      { cls: "聋人研究助理", name: "顾耀勇" },
      { cls: "聋人研究助理", name: "陈淡柳" },
      { cls: "聋人研究助理", name: "单仁冰" },
      // { cls: "研究生", name: "孙  玲" },
      // { cls: "研究生", name: "和子晴" },
      // { cls: "财务人员", name: "王玲玲" }
    ]
  }),
  methods: {
    handleDetail(row) {
      this.$data.emp = row;
      this.$data.ifDetail = true;
      var path = "emp/"+ row.img;
      this.$data.imgSrc = path;
    },
    backToTeam(){
      this.$data.ifDetail = false
    }
  }
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
.clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }
</style>
