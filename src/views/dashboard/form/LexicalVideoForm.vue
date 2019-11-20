<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Video')" prop="videoPath">
      <div class="video-container">
        <video-player :src="formData.videoPath" />
      </div>
    </el-form-item>

    <el-form-item :label="$t('Performer')" prop="performerID" required>
      <performer-selector v-model="formData.performerID" />
    </el-form-item>

    <el-form-item :label="$t('Word')" prop="lexicalWordID" required>
      <word-selector v-model="formData.lexicalWordID" />
    </el-form-item>

    <el-form-item :label="$t('WordFormation')" prop="wordFormation" required>
      <word-formation-selector v-model="formData.wordFormation" />
    </el-form-item>

    <el-form-item v-if="formData.wordFormation==='compound'" :label="$t('Morpheme')" prop="morpheme">
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
import VideoPlayer from "@/components/video/VideoPlayer"
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
    VideoPlayer
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        performerID: [{ required: true, message: "请选择被试" }],
        lexicalWordID: [{ required: true, message: "请选择词语" }],
        videoPath: [{ required: true, message: "请上传视频" }],
        wordFormation: [{ required: true, message: "请选择构词方式" }]
      }
    };
  },
  computed: {
    ...mapGetters(["settings"])
  }
};
</script>