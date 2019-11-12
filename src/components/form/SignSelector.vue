<template>
  <el-select v-model="sign" clearable :placeholder="tip" @clear="$emit('clear')">
    <el-option v-for="item in signs" :key="item.id" :label="item.name" :value="item.id">
      <span>{{ item.name }}</span>
      <img :src="'https://ccsl.shu.edu.cn/public/'+item.image" :alt="item.name" />
    </el-option>
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "left": "请选择左手手形",
    "right": "请选择右手手形",
    "sign": "请选择手形"
  },
  "en-US": {
    "left": "Select left hand sign",
    "right": "Select right hand sign",
    "sign": "Select sign"
  }
}
</i18n>

<script>
import { GetSignsList } from "@/api/signs";
export default {
  name: "SignSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: ""
    },
    orientation: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      signs: []
    };
  },
  computed: {
    sign: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit("update", val);
      }
    },
    tip() {
      if (this.orientation === "left") return this.$t("left");
      else if (this.orientation === "right") return this.$t("right");
      else return this.$t("sign");
    }
  },
  created() {
    GetSignsList({ limit: 0 }).then(res => {
      this.signs = res.data;
    });
  }
};
</script>

<style lang="scss" scoped>
.el-select-dropdown__item {
  display: flex;
  justify-content: space-around;
  align-items: center;

  height: 60px;
  img {
    max-height: 100%;
    max-width: 70px;
    padding: 5px;
  }
}
</style>