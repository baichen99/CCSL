<template>
  <div
    ref="rightPanel"
    :class="{show:show}"
    class="rightPanel-container"
  >
    <div class="rightPanel-background" />
    <div class="rightPanel">
      <!-- <div
        class="handle-button"
        :style="{'top':buttonTop+'px','background-color':theme}"
        @click="show=!show"
      >
        <i :class="show?'el-icon-close':'el-icon-setting'" />
      </div> -->
      <div class="rightPanel-items">
        <slot />
        <el-button @click="show=false">取消</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { addClass, removeClass } from "@/utils";

export default {
  name: "RightPanel",
  model: {
    prop: "showPanel",
    event: "update"
  },
  props: {
    showPanel: {
      default: () => false,
      type: Boolean
    }
  },
  computed: {
    // theme() {
    //   return this.$store.state.settings.theme;
    // },
    show: {
      get() {
        return this.showPanel;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  watch: {
    show(value) {
      if (value) {
        addClass(document.body, "showRightPanel");
      } else {
        removeClass(document.body, "showRightPanel");
      }
    }
  },
  mounted() {
    this.insertToBody();
  },
  beforeDestroy() {
    this.removeFromBody();
  },
  methods: {
    addEventClick() {
      window.addEventListener("click", this.closeSidebar);
    },
    insertToBody() {
      const elx = this.$refs.rightPanel;
      const body = document.querySelector("body");
      body.insertBefore(elx, body.firstChild);
    },
    removeFromBody() {
      const elx = this.$refs.rightPanel;
      elx.remove();
    }
  }
};
</script>

<style>
.showRightPanel {
  overflow: hidden;
  position: relative;
  width: calc(100% - 100px);
}
</style>

<style lang="scss" scoped>
.rightPanel-background {
  position: fixed;
  top: 0;
  left: 0;
  opacity: 0;
  transition: opacity 0.3s cubic-bezier(0.7, 0.3, 0.1, 1);
  background: rgba(0, 0, 0, 0.2);
  z-index: -1;
}

.rightPanel {
  width: 100%;
  max-width: 700px;
  height: 100vh;
  position: fixed;
  top: 0;
  right: 0;
  box-shadow: 0px 0px 15px 0px rgba(0, 0, 0, 0.05);
  transition: all 0.25s cubic-bezier(0.7, 0.3, 0.1, 1);
  transform: translate(100%);
  background: #fff;
  z-index: 40000;
}

.rightPanel-items {
  margin: 20px;
}

.show {
  transition: all 0.3s cubic-bezier(0.7, 0.3, 0.1, 1);

  .rightPanel-background {
    z-index: 20000;
    opacity: 1;
    width: 100%;
    height: 100%;
  }

  .rightPanel {
    transform: translate(0);
  }
}

// .handle-button {
//   width: 48px;
//   height: 48px;
//   position: absolute;
//   left: -48px;
//   text-align: center;
//   font-size: 24px;
//   border-radius: 6px 0 0 6px !important;
//   z-index: 0;
//   pointer-events: auto;
//   cursor: pointer;
//   color: #fff;
//   line-height: 48px;
//   i {
//     font-size: 24px;
//     line-height: 48px;
//   }
// }
</style>
