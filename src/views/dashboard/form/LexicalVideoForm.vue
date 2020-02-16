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
      <simple-selector
        v-model="formData.wordFormation"
        :options="wordFormations"
        @update="resetMorpheme"
      />
    </el-form-item>

    <el-form-item :label="$t('Morpheme')" prop="morpheme">
      <morphemes-picker v-model="formData.morpheme" />
    </el-form-item>

    <el-form-item :label="$t('LeftHandshape')" prop="leftHandshapesID">
      <multi-handshape-selector v-model="formData.leftHandshapesID" />
    </el-form-item>

    <el-form-item :label="$t('RightHandshape')" prop="rightHandshapesID">
      <multi-handshape-selector v-model="formData.rightHandshapesID" />
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from "vuex";
import VideoUploader from "@/components/video/VideoUploader";
import SimpleSelector from "@/components/form/SimpleSelector.vue";
import MorphemesPicker from "@/components/form/MorphemesPicker";
import PerformerSelector from "@/components/form/PerformerSelector";
import WordSelector from "@/components/form/WordSelector";
import MultiHandshapeSelector from "@/components/form/MultiHandshapeSelector";
import formMixin from "./formMixin";
export default {
  name: "LexicalVideoForm",
  components: {
    SimpleSelector,
    MorphemesPicker,
    PerformerSelector,
    WordSelector,
    MultiHandshapeSelector,
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
    ...mapGetters(["settings", "wordFormations"])
  },

  methods: {
    resetMorpheme() {
      this.formData.morpheme = [];
    }
  }
};
</script>