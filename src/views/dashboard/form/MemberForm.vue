<template>
  <el-form ref="form" label-position="left" label-width="200px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Profile')" prop="profile">
      <image-uploader v-model="formData.profile" dir="members" />
    </el-form-item>

    <el-form-item :label="$t('NameZh')" prop="nameZh">
      <el-input v-model="formData.nameZh" :placeholder="$t('nameZhTip')" />
    </el-form-item>

    <el-form-item :label="$t('NameEn')" prop="nameEn">
      <el-input v-model="formData.nameEn" :placeholder="$t('nameEnTip')" />
    </el-form-item>

    <el-form-item :label="$t('Type')" prop="type">
      <simple-selector v-model="formData.type" :options="memberTypes" />
    </el-form-item>

    <el-form-item :label="$t('Degree')" prop="degree">
      <simple-selector v-model="formData.degree" :options="memberDegrees" />
    </el-form-item>

    <el-form-item :label="$t('EmployerZh')" prop="employerZh">
      <el-input v-model="formData.employerZh" :placeholder="$t('employerZhTip')" />
    </el-form-item>

    <el-form-item :label="$t('EmployerEn')" prop="employerEn">
      <el-input v-model="formData.employerEn" :placeholder="$t('employerEnTip')" />
    </el-form-item>

    <el-form-item :label="$t('DescriptionZh')" prop="descriptionZh">
      <rich-text-editor v-model="formData.descriptionZh" />
    </el-form-item>

    <el-form-item :label="$t('DescriptionEn')" prop="descriptionEn">
      <rich-text-editor v-if="formData.descriptionEn!==null" v-model="formData.descriptionEn" />
    </el-form-item>
  </el-form>
</template>

<i18n>
{
    "zh-CN": {
        "nameZhTip": "请输入中文名",
        "nameEnTip": "请输入英文名",
        "employerZhTip": "请输入单位（中文）",
        "employerEnTip": "请输入单位（英文）",
        "descriptionZhTip": "请输入介绍（中文）",
        "descriptionEnTip": "请输入介绍（英文）"
    },
    "en-US": {
        "nameZhTip": "Input name(Chinese)",
        "nameEnTip": "Input name(English)",
        "employerZhTip": "Input employer(Chinese)",
        "employerEnTip": "Input employer(English)",
        "descriptionZhTip": "Input description(Chinese)",
        "descriptionEnTip": "Input desciprtion(English)"
    }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import ImageUploader from "@/components/form/ImageUploader";
import RichTextEditor from "@/components/form/RichTextEditor";
import SimpleSelector from "@/components/form/SimpleSelector";
import formMixin from "./formMixin";
export default {
  name: "MemberForm",
  components: {
    ImageUploader,
    RichTextEditor,
    SimpleSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        type: [{ required: true, message: "请选择类型", trigger: "blur" }],
        nameZh: [
          { required: true, message: "请输入中文姓名", trigger: "blur" }
        ],
        nameEn: [{ required: true, message: "请输入英文姓名", trigger: "blur" }]
      }
    };
  },
  computed: {
    ...mapGetters(["memberTypes", "memberDegrees"])
  }
};
</script>