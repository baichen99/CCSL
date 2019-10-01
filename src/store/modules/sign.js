import { getWords } from "@/api/words";

const state = {
  sign: [
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "10",
    "10号变体",
    "11",
    "12",
    "13",
    "14",
    "15",
    "16",
    "17",
    "18",
    "19",
    "20",
    "21",
    "21号变体",
    "22",
    "23",
    "24",
    "25",
    "26",
    "27",
    "28",
    "29",
    "30",
    "31",
    "32",
    "33",
    "34",
    "35",
    "36",
    "37",
    "38",
    "39",
    "40",
    "41",
    "42",
    "43",
    "44",
    "45",
    "46",
    "47",
    "48",
    "49",
    "50",
    "51",
    "52",
    "53",
    "54",
    "55",
    "56",
    "57",
    "58",
    "59",
    "60",
    "61",
    "62",
    "63",
    "64",
    "65",
    "66",
    "67",
    "68",
    "69",
    "L17",
    "L37",
    "Y14",
    "Y16",
    "Y42",
    "Y43",
    "Y62",
    "Y72",
    "Y76",
    "Y82",
    "Y86",
    "Y90",
    "Z5",
    "Z10",
    "Z33",
    "Z50",
    "Z51",
    "Z54",
    "ZH"
  ],
  regions: [
    "通用手语",
    "徐汇区",
    "嘉定区",
    "杨浦区",
    "松江区",
    "青浦区",
    "奉贤区",
    "普陀区",
    "长宁区",
    "浦东新区",
    "金山区",
    "黄浦区",
    "宝山区",
    "静安区",
    "虹口区",
    "闵行区",
    "崇明区"
  ],
  wordTypes: [
    "名词",
    "动词/名词",
    "助动词",
    "数词",
    "量词",
    "名词/形容词",
    "动词",
    "连词",
    "形容词/动词",
    "方位名词",
    "疑问词",
    "介词",
    "副词",
    "代词",
    "形容词/副词",
    "形容词",
    "疑问代词"
  ],
  words: {}
};

const mutations = {
  SET_WORDS(state, { initial, item }) {
    if (state.words[initial]) {
      let itemFound = false;
      state.words[initial].map(word => {
        if (word.id === item.id) {
          itemFound = true;
        }
      });
      if (!itemFound) {
        state.words[initial].push(item);
      }
    } else {
      state.words[initial] = [item];
    }
  }
};

const actions = {
  getWordsDict({ commit }) {
    getWords({ limit: 0 }).then(res => {
      res.data.map(item => {
        const initial = item.initial;
        commit("SET_WORDS", { initial, item });
      });
    });
  }
};

export default {
  namespaced: true,
  state,
  actions,
  mutations
};
