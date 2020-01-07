const TOGGLE_SIDEBAR = "TOGGLE_SIDEBAR";
const CLOSE_SIDEBAR = "CLOSE_SIDEBAR";
const TOGGLE_DEVICE = "TOGGLE_DEVICE";
const SET_LOCALE = "SET_LOCALE";

const state = {
  sidebar: {
    opened: true,
    withoutAnimation: false
  },
  device: "desktop",
  locale: ""
};

const mutations = {
  TOGGLE_SIDEBAR: state => {
    state.sidebar.opened = !state.sidebar.opened;
    state.sidebar.withoutAnimation = false;
  },
  CLOSE_SIDEBAR: (state, withoutAnimation) => {
    state.sidebar.opened = false;
    state.sidebar.withoutAnimation = withoutAnimation;
  },
  TOGGLE_DEVICE: (state, device) => {
    state.device = device;
  },
  SET_LOCALE: (state, locale) => {
    state.locale = locale;
  }
};

const actions = {
  toggleSideBar({ commit }) {
    commit(TOGGLE_SIDEBAR);
  },
  closeSideBar({ commit }, { withoutAnimation }) {
    commit(CLOSE_SIDEBAR, withoutAnimation);
  },
  toggleDevice({ commit }, device) {
    commit(TOGGLE_DEVICE, device);
  },
  setLocale({ commit }, locale) {
    commit(SET_LOCALE, locale);
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
