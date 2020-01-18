import Vue from "vue";
import ElementUI from "element-ui";
import App from "./App";
import i18n from "@/locale";
import store from "@/store";
import router from "@/router";
import cities from "@/assets/cities";

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

Vue.filter("getRegionName", regionID => {
  const provinceID = parseInt(regionID.toString().slice(0, 2));
  const cityID = parseInt(regionID.toString().slice(0, 4));
  const provinces = cities.filter(province => province.value === provinceID);
  if (provinces.length === 1) {
    const province = provinces[0];
    const cities = province.children.filter(city => city.value === cityID);
    if (cities.length == 1) {
      const city = cities[0];
      const districts = city.children.filter(
        district => district.value === regionID
      );
      if (districts.length == 1) {
        const district = districts[0];
        return [province.label, city.label, district.label].join(" / ");
      }
    }
  }
});

Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  store,
  i18n,
  render: h => h(App)
});
