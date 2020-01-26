import lodash from "lodash";
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
    "疑问词",
    "介词",
    "副词",
    "代词",
    "形容词/副词",
    "形容词"
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
  partOfSpeech: {
    noun: { name: "Noun" },
    verb: { name: "Verb" },
    adjective: { name: "Adjective" },
    auxiliary: { name: "Auxiliary" },
    numeral: { name: "Numeral" },
    classifier: { name: "Classifier" },
    pronoun: { name: "Pronoun" },
    conjunction: { name: "Conjunction" },
    interrogative: { name: "Interrogative" },
    adverb: { name: "Adverb" },
    prepositions: { name: "Prepositions" }
  },
  signs: {},
  lexicons: {},
  performers: {}
};

const mutations = {
  SET_LEXICON: (state, data) => {
    data.map(item => {
      state.lexicons[item.id] = item;
    });
  },
  SET_SIGN: (state, data) => {
    data.map(item => {
      state.signs[item.id] = item;
    });
  },
  SET_PERFORMER: (state, data) => {
    data.map(item => {
      state.performers[item.id] = item;
    });
  }
};

const actions = {
  getLexicons({ commit, state }, force) {
    return new Promise((resolve, reject) => {
      if (force || lodash.isEmpty(state.lexicons)) {
        GetWordsList({ limit: 0 })
          .then(res => {
            const { data } = res;
            commit(SET_LEXICON, data);
            resolve();
          })
          .catch(err => {
            reject(err);
          });
      } else {
        resolve();
      }
    });
  },

  getPerformers({ commit, state }, force) {
    return new Promise((resolve, reject) => {
      if (force || lodash.isEmpty(state.performers)) {
        GetPerformersList({ limit: 0 })
          .then(res => {
            const { data } = res;
            commit(SET_PERFORMER, data);
            resolve();
          })
          .catch(err => {
            reject(err);
          });
      } else {
        resolve();
      }
    });
  },

  getSigns({ commit, state }, force) {
    return new Promise((resolve, reject) => {
      if (force || lodash.isEmpty(state.signs)) {
        GetSignsList({ limit: 0 })
          .then(res => {
            const { data } = res;
            commit(SET_SIGN, data);
            resolve();
          })
          .catch(err => {
            reject(err);
          });
      } else {
        resolve();
      }
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
