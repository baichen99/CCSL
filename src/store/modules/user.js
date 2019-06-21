import { login, getInfo } from "@/api/users";
import { getToken, setToken, removeToken } from "@/utils/tools";
// import { resetRouter } from "@/router";

const state = {
  token: getToken(),
  name: "",
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
  SET_ROLES: (state, role) => {
    state.roles = [role];
  }
};

const actions = {
  login({ commit }, userInfo) {
    console.log("BEFORE_LOGIN");
    const { username, password } = userInfo;
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password })
        .then(response => {
          const token = response.data;
          commit("SET_TOKEN", token);
          setToken(token);
          const user = JSON.parse(atob(token.split(".")[1]));
          commit("SET_ID", user.user);
          resolve();
        })
        .catch(error => {
          console.log("LOGIN_ERROR");
          reject(error);
        });
    });
  },
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.id)
        .then(response => {
          const { data } = response;
          console.log(data);
          if (!data) {
            reject("验证失败，请重新登录");
          }
          commit("SET_ROLES", data.userType);
          commit("SET_NAME", data.name);
          resolve(data);
        })
        .catch(error => {
          reject(error);
        });
    });
  },
  logout({ commit }) {
    return new Promise(resolve => {
      commit("SET_TOKEN", "");
      commit("SET_ID", "");
      commit("SET_ROLES", []);
      removeToken();
      resolve();
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
