import Vue from "vue";
import Vuex from "vuex";
import getters from "./getters";
import app from "./modules/app";
import permission from "./modules/permission";
import settings from "./modules/settings";
import signlang from "./modules/signlang";
import user from "./modules/user";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    app,
    permission,
    settings,
    user,
    signlang
  },
  getters
});

export default store;
