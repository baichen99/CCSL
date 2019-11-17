<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item label="视频" prop="videoPath">
      <div class="video-container">
        <video-player
          v-if="formData.videoPath"
          class="vjs-big-play-centered"
          :options="{ 
            sources: [{ 
              type: 'video/mp4', 
              src: settings.publicURL + formData.videoPath 
            }]
          }"
          @contextmenu.native.prevent
        />
      </div>
    </el-form-item>

    <el-form-item label="被试" prop="performerID" required>
      <performer-selector v-model="formData.performerID" />
    </el-form-item>

    <el-form-item label="词语" prop="lexicalWordID" required>
      <word-selector v-model="formData.lexicalWordID" />
    </el-form-item>

    <el-form-item label="构词方式" prop="constructType" required>
      <word-construct-selector v-model="formData.constructType" />
    </el-form-item>

    <el-form-item v-if="formData.constructType==='复合词'" label="复合词构词词语" prop="constructWords">
      <construct-words-picker v-model="formData.constructWords" />
    </el-form-item>

    <el-form-item label="左手手形" prop="leftSigns">
      <multi-sign-selector v-model="formData.leftSignsID" />
    </el-form-item>

    <el-form-item label="右手手形" prop="rightSigns">
      <multi-sign-selector v-model="formData.rightSignsID" />
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from "vuex";
import WordConstructSelector from "@/components/form/WordConstructSelector";
import ConstructWordsPicker from "@/components/form/ConstructWordsPicker";
import PerformerSelector from "@/components/form/PerformerSelector.vue";
import WordSelector from "@/components/form/WordSelector.vue";
import MultiSignSelector from "@/components/form/MultiSignSelector.vue";
import formMixin from "./formMixin";
export default {
  name: "LexicalVideoForm",
  components: {
    WordConstructSelector,
    ConstructWordsPicker,
    PerformerSelector,
    WordSelector,
    MultiSignSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        performerID: [{ required: true, message: "请选择被试" }],
        lexicalWordID: [{ required: true, message: "请选择词语" }],
        // videoPath: [{ required: true, message: "请上传视频" }],
        constructType: [{ required: true, message: "请选择构词方式" }]
      }
    };
  },
  computed: {
    ...mapGetters(["settings"])
  }
};
</script>