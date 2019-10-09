const state = {
  wordTypes: [
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
  newsTypes: {
    news: { icon: "news", name: "新闻动态" },
    activity: { icon: "activity", name: "学术活动" },
    notice: { icon: "notice", name: "通知公告" },
    research: { icon: "research", name: "科研动态" }
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
