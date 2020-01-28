import Vue from "vue";
import cities from "@/assets/cities";

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
