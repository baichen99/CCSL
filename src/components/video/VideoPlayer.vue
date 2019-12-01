<template>
  <div>
    <video ref="player">
      <source :src="settings.publicURL + src" type="video/mp4" />
    </video>
  </div>
</template>



<script>
import { mapGetters } from "vuex";
import Plyr from "plyr";

export default {
  name: "VideoPlayer",
  props: {
    src: {
      type: String,
      default: () => ""
    }
  },
  data() {
    return {
      player: null
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  mounted() {
    this.player = new Plyr(this.$refs.player, {
      controls: ["play-large", "progress"]
    });
  }
};
</script>


<style lang="scss" scoped>
</style>