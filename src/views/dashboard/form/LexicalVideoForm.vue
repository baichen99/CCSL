<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Video')" prop="videoPath">
      <div class="video-container">
        <video-uploader v-model="formData.videoPath" dir="lexical_videos" />
      </div>
    </el-form-item>

    <el-form-item :label="$t('Performer')" prop="performerID">
      <performer-selector v-model="formData.performerID" />
    </el-form-item>

    <el-form-item :label="$t('Word')" prop="lexiconID">
      <word-selector v-model="formData.lexiconID" />
    </el-form-item>

    <el-form-item :label="$t('WordFormation')" prop="wordFormation">
      <word-formation-selector v-model="formData.wordFormation" @update="resetMorpheme" />
    </el-form-item>

    <el-form-item :label="$t('Morpheme')" prop="morpheme">
      <morphemes-picker v-model="formData.morpheme" />
    </el-form-item>

    <el-form-item :label="$t('LeftSign')" prop="leftSigns">
      <multi-sign-selector v-model="formData.leftSignsID" />
    </el-form-item>

    <el-form-item :label="$t('RightSign')" prop="rightSigns">
      <multi-sign-selector v-model="formData.rightSignsID" />
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from "vuex";
import VideoUploader from "@/components/video/VideoUploader";
import WordFormationSelector from "@/components/form/WordFormationSelector";
import MorphemesPicker from "@/components/form/MorphemesPicker";
import PerformerSelector from "@/components/form/PerformerSelector";
import WordSelector from "@/components/form/WordSelector";
import MultiSignSelector from "@/components/form/MultiSignSelector";
import formMixin from "./formMixin";
export default {
  name: "LexicalVideoForm",
  components: {
    WordFormationSelector,
    MorphemesPicker,
    PerformerSelector,
    WordSelector,
    MultiSignSelector,
    VideoUploader
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        performerID: [{ required: true, message: "请选择被试" }],
        lexiconID: [{ required: true, message: "请选择词语" }],
        videoPath: [{ required: true, message: "请上传视频" }],
        wordFormation: [{ required: true, message: "请选择构词方式" }]
      }
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },

  methods: {
    resetMorpheme() {
      this.formData.morpheme = [];
    }
  }
};
</script>