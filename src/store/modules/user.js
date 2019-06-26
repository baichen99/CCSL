import { login, getInfo } from "@/api/users";
import { getToken, setToken, removeToken } from "@/utils/tools";

const state = {
  token: getToken(),
  name: "",
  avatar: "https://axiom-public.axiomacademy.cn/static/default.png",
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
  login({ commit }, user) {
    const { username, password, remember } = user;
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password })
        .then(response => {
          const token = response.data;
          commit("SET_TOKEN", token);
          // If remember password, save JWT token to local storage
          if (remember) {
            setToken(token);
          }
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

  logout({ commit }) {
    return new Promise(resolve => {
      commit("SET_TOKEN", "");
      commit("SET_ID", "");
      commit("SET_ROLES", []);
      removeToken();
      resolve();
    });
  },

  getUserInfo({ commit, state }) {
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
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
