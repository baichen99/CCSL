import lodash from "lodash";
import { GetWordsList } from "@/api/lexicons";
import { GetPerformersList } from "@/api/performers";
import { GetHandshapesList } from "@/api/handshapes";

const SET_LEXICON = "SET_LEXICON";
const SET_PERFORMER = "SET_PERFORMER";
const SET_HANDSHAPE = "SET_HANDSHAPE";

const state = {
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
    data: [
      { text: "CompoundWord", value: "compound" },
      { text: "SimpleWord", value: "simple" }
    ],
    placeholder: {
      "zh-CN": "请选择构词方式",
      "en-US": "Select word formation"
    }
  },
  userRoles: {
    data: [
      { text: "SuperAdmin", value: "super", color: "danger" },
      { text: "Admin", value: "admin", color: "warning" },
      { text: "DatabaseUser", value: "dbuser", color: "" },
      { text: "Teacher", value: "teacher", color: "success" },
      { text: "Student", value: "student", color: "info" }
    ], placeholder: {
      "zh-CN": "请选择用户角色",
      "en-US": "Select user role"
    }
  },
  userState: {
    data: [
      { text: "Inactive", value: "inactive", color: "warning" },
      { text: "Active", value: "active", color: "success" }
    ],
    placeholder: {
      "zh-CN": "请选择账号状态",
      "en-US": "Select user state"
    }
  },
  genderTypes: {
    data: [
      { text: "Male", value: "M" },
      { text: "Female", value: "F" }
    ],
    placeholder: {
      "zh-CN": "请选择性别",
      "en-US": "Select gender"
    }
  },
  newsColumns: {
    data: [
      { text: "NewsColumn", value: "news", icon: "news" },
      { text: "ActivityColumn", value: "activity", icon: "activity" },
      { text: "NoticeColumn", value: "notice", icon: "notice" },
      { text: "ResearchColumn", value: "research", icon: "research" }
    ],
    placeholder: {
      "zh-CN": "请选择新闻栏目",
      "en-US": "Select news column"
    }
  },
  newsState: {
    data: [
      { text: "Draft", value: "draft", color: "warning" },
      { text: "Published", value: "published", color: "success" }
    ],
    placeholder: {
      "zh-CN": "请选择新闻状态",
      "en-US": "Select news state"
    }
  },
  newsTypes: {
    data: [
      { text: "Link", value: "link" },
      { text: "Document", value: "document" }
    ],
    placeholder: {
      "zh-CN": "请选择新闻类型",
      "en-US": "Select news type"
    }
  },
  memberTypes: {
    data: [
      { text: "SeniorConsultant", value: "consultant" },
      { text: "ResearchFellow", value: "researchFellow" },
      { text: "AssistantResearchFellow", value: "assistantResearchFellow" },
      { text: "SignLanguageTranslator", value: "signLanguageTranslator" },
      { text: "ResearchAssistantDeaf", value: "researchAssistantDeaf" },
      { text: "Postgraduate", value: "postgraduate" }
    ],
    placeholder: {
      "zh-CN": "请选择成员类型",
      "en-US": "Select member type"
    }
  },
  memberDegrees: {
    data: [
      { text: "Bachelor", value: "bachelor" },
      { text: "Master", value: "master" },
      { text: "Doctor", value: "doctor" }
    ],
    placeholder: {
      "zh-CN": "请选择成员学位",
      "en-US": "Select meber degree"
    }
  },
  languageTypes: {
    data: [
      { text: "Chinese", value: "zh-CN" },
      { text: "English", value: "en-US" }
    ],
    placeholder: {
      "zh-CN": "请选择语言",
      "en-US": "Select language"
    }
  },
  partOfSpeech: {
    data: [
      { text: "Noun", value: "noun" },
      { text: "Verb", value: "verb" },
      { text: "Adjective", value: "adjective" },
      { text: "Auxiliary", value: "auxiliary" },
      { text: "Numeral", value: "numeral" },
      { text: "Classifier", value: "classifier" },
      { text: "Pronoun", value: "pronoun" },
      { text: "Conjunction", value: "conjunction" },
      { text: "Interrogative", value: "interrogative" },
      { text: "Adverb", value: "adverb" },
      { text: "Prepositions", value: "prepositions" }
    ],
    placeholder: {
      "zh-CN": "请选择词性",
      "en-US": "Select part of speech"
    }
  },
  handshapes: {},
  lexicons: {},
  performers: {}
};

const mutations = {
  SET_LEXICON: (state, data) => {
    data.map(item => {
      state.lexicons[item.id] = item;
    });
  },
  SET_HANDSHAPE: (state, data) => {
    data.map(item => {
      state.handshapes[item.id] = item;
    });
  },
  SET_PERFORMER: (state, data) => {
    data.map(item => {
      state.performers[item.id] = item;
    });
  }
};

const actions = {
  async getLexicons({ commit, state }, force) {
    if (force || lodash.isEmpty(state.lexicons)) {
      try {
        const res = await GetWordsList({ limit: 0 })
        const { data } = res;
        commit(SET_LEXICON, data);
      } catch (err) {
        console.error(err)
      }
    }
  },

  async getPerformers({ commit, state }, force) {
    if (force || lodash.isEmpty(state.performers)) {
      try {
        const res = await GetPerformersList({ limit: 0 })
        const { data } = res;
        commit(SET_PERFORMER, data);
      } catch (err) {
        console.error(err)
      }
    }
  },

  async getHandshapes({ commit, state }, force) {
    if (force || lodash.isEmpty(state.handshapes)) {
      try {
        const res = await GetHandshapesList({ limit: 0 })
        const { data } = res;
        commit(SET_HANDSHAPE, data);
      } catch (err) {
        console.error(err)
      }
    }
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
