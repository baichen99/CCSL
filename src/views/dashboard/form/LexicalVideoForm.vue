<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Video')" prop="videoPath">
      <div class="video-container">
        <video-uploader v-model="formData.videoPath" dir="lexical_videos" />
      </div>
    </el-form-item>

    <el-form-item :label="$t('Performer')" prop="performerID">
      <remote-selector
        v-model="formData.performerID"
        :get-list-method="GetPerformersList"
        :formatter="(item)=>item.name"
      >
        <template
          #default="{item}"
        >{{ item.name }} - {{ $t($options.filters.getObjectItem(genderTypes,item.gender).text) }} - {{ item.regionID | getRegionName }}</template>
      </remote-selector>
    </el-form-item>

    <el-form-item :label="$t('Word')" prop="lexiconID">
      <remote-selector
        v-model="formData.lexiconID"
        :get-list-method="GetWordsList"
        :formatter="(item)=>item.chinese"
      >
        <template #default="{item}">
          {{ item.initial }} -
          <span
            class="tag-value word-sup"
            v-html="$options.filters.addNumberSup(item.chinese)"
          ></span>
        </template>
      </remote-selector>
    </el-form-item>

    <el-form-item :label="$t('WordFormation')" prop="wordFormation">
      <simple-selector v-model="formData.wordFormation" :options="wordFormations" />
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
import { GetPerformersList } from "@/api/performers";
import { GetWordsList } from "@/api/lexicons";
import { mapGetters } from "vuex";
import VideoUploader from "@/components/video/VideoUploader";
import SimpleSelector from "@/components/form/SimpleSelector.vue";
import MorphemesPicker from "@/components/form/MorphemesPicker";
import RemoteSelector from "@/components/form/RemoteSelector";
import MultiHandshapeSelector from "@/components/form/MultiHandshapeSelector";
import formMixin from "./formMixin";
export default {
  name: "LexicalVideoForm",
  components: {
    SimpleSelector,
    MorphemesPicker,
    RemoteSelector,
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
      },
      GetPerformersList,
      GetWordsList
    };
  },
  computed: {
    ...mapGetters(["wordFormations", "genderTypes"])
  }
};
</script>