import { Login, RefreshToken, GetUser } from "@/api/users";
import {
  getToken,
  setToken,
  removeToken,
  setUser,
  getUser,
  removeUser
} from "@/utils/tools";

const state = {
  token: getToken(),
  name: "",
  username: "",
  avatar: "https://ccsl.shu.edu.cn/public/assets/default.png",
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
  },
  SET_USERNAME: (state, username) => {
    state.username = username;
  }
};

const actions = {
  login({ commit }, user) {
    const { username, password, remember } = user;
    return new Promise((resolve, reject) => {
      Login({ username: username.trim(), password: password })
        .then(response => {
          const token = response.data;
          commit("SET_TOKEN", token);
          // If remember password, save JWT token to local storage
          const user = JSON.parse(atob(token.split(".")[1]));
          const userID = user.user;
          if (remember) {
            setToken(token);
            setUser(userID);
          }
          GetUser(userID)
            .then(response => {
              const { data } = response;
              if (!data) {
                reject("身份校验失败，请重新登录");
              }
              commit("SET_ID", data.id);
              commit("SET_NAME", data.name);
              commit("SET_ROLES", data.userType);
              commit("SET_USERNAME", data.username);
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
      commit("SET_TOKEN", "");
      commit("SET_ID", "");
      commit("SET_ROLES", []);
      removeToken();
      removeUser();
      resolve();
    });
  },

  getUserInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      const userID = getUser() || state.id;
      GetUser(userID)
        .then(res => {
          commit("SET_ROLES", res.data.userType);
          commit("SET_NAME", res.data.name);
          resolve(res.data);
        })
        .catch(err => {
          reject(err);
        });
    });
  },

  refreshToken({ commit }) {
    return new Promise((resolve, reject) => {
      if (getToken()) {
        RefreshToken()
          .then(res => {
            commit("SET_TOKEN", res.data);
            setToken(res.data);
            resolve();
          })
          .catch(err => {
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
