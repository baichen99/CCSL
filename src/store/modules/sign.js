const state = {
  wordPosTypes: [
    "名词",
    "动词/名词",
    "助动词",
    "数词",
    "量词",
    "名词/形容词",
    "动词",
    "连词",
    "形容词/动词",
    "方位名词",
    "疑问词",
    "介词",
    "副词",
    "代词",
    "形容词/副词",
    "形容词",
    "疑问代词"
  ],
  constructTypes: ["单纯词", "复合词"],
  newsColumns: {
    news: { icon: "news", name: "新闻动态", value: "news" },
    activity: { icon: "activity", name: "学术活动", value: "activity" },
    notice: { icon: "notice", name: "通知公告", value: "notice" },
    research: { icon: "research", name: "科研动态", value: "research" }
  },
  userTypes: {
    super: { name: "超级管理员", color: "danger" },
    admin: { name: "普通管理员", color: "warning" },
    learner: { name: "学习平台用户", color: "info" },
    user: { name: "语料库用户", color: "" }
  },
  userState: {
    active: { name: "活跃", color: "success" },
    inactive: { name: "停用", color: "danger" }
  },
  genderTypes: {
    M: { name: "男", value: "M" },
    F: { name: "女", value: "F" }
  }
};

const mutations = {};

const actions = {};

export default {
  namespaced: true,
  state,
  actions,
  mutations
};
