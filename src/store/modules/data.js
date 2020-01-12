import { GetWordsList } from "@/api/lexicons";
import { GetPerformersList } from "@/api/performers";
import { GetSignsList } from "@/api/signs";

const SET_LEXICON = "SET_LEXICON";
const SET_PERFORMER = "SET_PERFORMER";
const SET_SIGN = "SET_SIGN";

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
    researchFellow: {
      name: "ResearchFellow"
    },
    assistantResearchFellow: {
      name: "AssistantResearchFellow"
    },
    signLanguageTranslator: {
      name: "SignLanguageTranslator"
    },
    researchAssistantDeaf: {
      name: "ResearchAssistantDeaf"
    },
    postgraduate: {
      name: "Postgraduate"
    }
  },
  memberDegrees: {
    bachelor: { name: "Bachelor" },
    master: { name: "Master" },
    doctor: { name: "Doctor" }
  },
  languageTypes: {
    "zh-CN": { name: "Chinese" },
    "en-US": { name: "English" }
  },
  signs: {},
  lexicons: {},
  performers: {}
};

const mutations = {
  SET_LEXICON: (state, data) => {
    state.lexicons[data.id] = data;
  },
  SET_SIGN: (state, data) => {
    state.signs[data.id] = data;
  },
  SET_PERFORMER: (state, data) => {
    state.performers[data.id] = data;
  }
};

const actions = {
  getLexicons({ commit }) {
    return new Promise((resolve, reject) => {
      GetWordsList({ limit: 0 })
        .then(res => {
          const { data } = res;
          data.map(item => {
            commit(SET_LEXICON, item);
          });
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  },

  getPerformers({ commit }) {
    return new Promise((resolve, reject) => {
      GetPerformersList({ limit: 0 })
        .then(res => {
          const { data } = res;
          data.map(item => {
            commit(SET_PERFORMER, item);
          });
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  },

  getSigns({ commit }) {
    return new Promise((resolve, reject) => {
      GetSignsList({ limit: 0 })
        .then(res => {
          const { data } = res;
          data.map(item => {
            commit(SET_SIGN, item);
          });
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
