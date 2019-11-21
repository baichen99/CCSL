import Vue from "vue";
import ElementUI from "element-ui";
import App from "./App";
import i18n from "@/locale";
import store from "@/store";
import router from "@/router";

import "@/directives";
import "@/utils/error";
import "normalize.css/normalize.css"; // CSS resets
import "@/styles/index.scss";
import "@/icons"; // icon

Vue.use(ElementUI, {
  zIndex: 1000,
  i18n: (key, value) => i18n.t(key, value)
});

Vue.filter("addNumberSup", word => {
  const reg = new RegExp(/[\d]+/g);
  return word.replace(reg, "<sup>$&</sup>");
});

Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  store,
  i18n,
  render: h => h(App)
});
