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
  wordInitial: [
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "W",
    "X",
    "Y",
    "Z"
  ],
  wordFormations: {
    compound: { name: "CompoundWord" },
    simple: { name: "SimpleWord" }
  },
  userTypes: {
    super: { name: "SuperAdmin", color: "danger" },
    admin: { name: "Admin", color: "warning" },
    learner: { name: "Learner", color: "info" },
    user: { name: "User", color: "" }
  },
  userState: {
    active: { name: "Active", color: "success" },
    inactive: { name: "Inactive", color: "danger" }
  },
  genderTypes: {
    M: { name: "Male" },
    F: { name: "Female" }
  },
  newsColumns: {
    news: { icon: "news", name: "NewsColumn" },
    activity: { icon: "activity", name: "ActivityColumn" },
    notice: { icon: "notice", name: "NoticeColumn" },
    research: { icon: "research", name: "ResearchColumn" }
  },
  newsState: {
    draft: { color: "warning", name: "Draft" },
    published: { color: "success", name: "Published" }
  },
  newsTypes: {
    link: { name: "Link" },
    document: { name: "Document" }
  },
  memberTypes: {
    consultant: {
      name: "SeniorConsultant"
    },
    researcher: {
      name: "Researcher"
    },
    assistantResearcher: {
      name: "AssistantResearcher"
    },
    signLanguageTranslator: {
      name: "SignLanguageTranslator"
    },
    deafResearchAssistant: {
      name: "DeafResearchAssistant"
    }
  },
  languageTypes: {
    "zh-CN": { name: "Chinese" },
    "en-US": { name: "English" }
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
