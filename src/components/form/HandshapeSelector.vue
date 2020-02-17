<template>
  <el-select
    v-model="data"
    filterable
    clearable
    :size="size"
    :placeholder="tip"
    :loading="loading"
    @clear="$emit('clear')"
  >
    <el-option v-for="item in options" :key="item.id" :value="item.id" :label="item.name">
      <span>{{ item.name }}</span>
      <img :src="settings.publicURL + item.image" :alt="item.name" />
    </el-option>
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "left": "请选择左手手形",
    "right": "请选择右手手形",
    "handshape": "请选择手形"
  },
  "en-US": {
    "left": "Select left handshape",
    "right": "Select right handshape",
    "handshape": "Select handshape"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "HandshapeSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: () => ""
    },
    orientation: {
      type: String,
      default: ""
    },
    size: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      options: []
    };
  },
  computed: {
    ...mapGetters(["settings"]),
    data: {
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
      else return this.$t("handshape");
    }
  },
  created() {
    this.$nextTick(async () => {
      this.loading = true;
      try {
        await this.$store.dispatch("data/getHandshapes");
        const data = this.$store.state.data.handshapes;
        this.options = Object.values(data);
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
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