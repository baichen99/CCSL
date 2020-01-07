import { GetWordsList } from "@/api/lexicons";
import { GetPerformersList } from "@/api/performers";
import { GetSignsList } from "@/api/signs";

const SET_LEXICON = "SET_LEXICON";
const SET_PERFORMER = "SET_PERFORMER";
const SET_SIGN = "SET_SIGN";

const state = {
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
