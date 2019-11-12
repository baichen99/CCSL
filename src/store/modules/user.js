import { Login, RefreshToken, GetUser } from "@/api/users";
import router from "@/router";

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
  SET_ROLES: (state, role) => {
    state.roles = role;
  },
  SET_USERNAME: (state, username) => {
    state.username = username;
  }
};

const actions = {
  login({ commit }, user) {
    const { username, password } = user;
    return new Promise((resolve, reject) => {
      Login({ username: username.trim(), password: password })
        .then(response => {
          const token = response.data;
          commit("SET_TOKEN", token);
          const user = JSON.parse(atob(token.split(".")[1]));
          const userID = user.user;
          GetUser(userID)
            .then(res => {
              const { data } = res;
              if (!data) {
                reject("身份校验失败，请重新登录");
              }
              commit("SET_ID", data.id);
              commit("SET_NAME", data.name);
              commit("SET_ROLES", [data.userType]);
              commit("SET_USERNAME", data.username);
              commit("SET_AVATAR", data.avatar);
              resolve(data);
            })
            .catch(err => {
              reject(err);
            });
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  logout({ commit }) {
    return new Promise(resolve => {
      commit("SET_ID", "");
      commit("SET_NAME", "");
      commit("SET_ROLES", []);
      commit("SET_USERNAME", "");
      commit("SET_AVATAR", "");
      commit("SET_TOKEN", "");
      router.push("/login");
      resolve();
    });
  },

  getUserInfo({ dispatch, commit, state }) {
    return new Promise((resolve, reject) => {
      const userID = state.id;
      GetUser(userID)
        .then(res => {
          const { data } = res;
          commit("SET_ID", data.id);
          commit("SET_NAME", data.name);
          commit("SET_ROLES", [data.userType]);
          commit("SET_USERNAME", data.username);
          commit("SET_AVATAR", data.avatar);
          resolve(res.data);
        })
        .catch(err => {
          dispatch("logout");
          reject(err);
        });
    });
  },

  refreshToken({ dispatch, commit }) {
    return new Promise((resolve, reject) => {
      if (state.token) {
        RefreshToken()
          .then(res => {
            commit("SET_TOKEN", res.data);
            resolve();
          })
          .catch(err => {
            dispatch("logout");
            reject(err);
          });
      }
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
