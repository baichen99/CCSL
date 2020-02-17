import { Login, RefreshToken, GetUser } from "@/api/users";
import router from "@/router";

const SET_TOKEN = "SET_TOKEN";
const SET_ID = "SET_ID";
const SET_NAME = "SET_NAME";
const SET_AVATAR = "SET_AVATAR";
const SET_ROLES = "SET_ROLES";
const SET_USERNAME = "SET_USERNAME";

const state = {
  token: "",
  name: "",
  username: "",
  avatar: "",
  id: "",
  roles: []
};

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_ID: (state, id) => {
    state.id = id;
  },
  SET_NAME: (state, name) => {
    state.name = name;
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar;
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles;
  },
  SET_USERNAME: (state, username) => {
    state.username = username;
  }
};

const actions = {
  async login({ commit }, user) {
    try {
      const { username, password } = user;
      const response = await Login({ username: username.trim(), password: password })
      const token = response.data;
      commit(SET_TOKEN, token);
      const userToken = JSON.parse(atob(token.split(".")[1]));
      const userID = userToken.user;
      const res = await GetUser(userID)
      const { data } = res;
      if (!data) {
        throw new Error("身份校验失败，请重新登录");
      }
      commit(SET_ID, data.id);
      commit(SET_NAME, data.name);
      commit(SET_ROLES, data.roles);
      commit(SET_USERNAME, data.username);
      commit(SET_AVATAR, data.avatar);
    } catch (err) {
      console.error(err)
    }
  },

  logout({ commit }) {
    commit(SET_ID, "");
    commit(SET_NAME, "");
    commit(SET_ROLES, []);
    commit(SET_USERNAME, "");
    commit(SET_AVATAR, "");
    commit(SET_TOKEN, "");
    router.push("/login");
  },

  async getUserInfo({ dispatch, commit, state }) {
    const userID = state.id;
    try {
      const res = await GetUser(userID);
      const { data } = res;
      commit(SET_ID, data.id);
      commit(SET_NAME, data.name);
      commit(SET_ROLES, data.roles);
      commit(SET_USERNAME, data.username);
      commit(SET_AVATAR, data.avatar);
    }
    catch (err) {
      dispatch("logout");
    }
  },

  async refreshToken({ dispatch, commit, state }) {
    if (state.token) {
      try {
        const res = await RefreshToken();
        commit(SET_TOKEN, res.data);
      } catch (err) {
        dispatch("logout");
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
