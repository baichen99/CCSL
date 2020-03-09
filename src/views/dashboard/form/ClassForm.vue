<template>
  <el-form ref="form" v-loading="loading" :model="formData" :rules="rules">
    <el-form-item :label="$t('ClassName')" prop="name">
      <el-input v-model="formData.name" :placeholder="$t('nameTip')" clearable />
    </el-form-item>
    <el-tabs type="border-card">
      <el-tab-pane :label="$t('ClassDetails')">
        <el-form-item>
          <rich-text-editor v-model="formData.details" height="600" dir="classes" />
        </el-form-item>
      </el-tab-pane>

      <el-tab-pane :label="$t('ClassResources')">
        <el-form-item>
          <rich-text-editor v-model="formData.resources" height="600" dir="classes" />
        </el-form-item>
      </el-tab-pane>

      <el-tab-pane v-if="mode==='edit'" label="学生管理">
        <student-form :cid="formData.id" :get-data-method="getData" :students="students" />
      </el-tab-pane>

      <el-tab-pane v-if="mode==='edit'" label="教师管理">
        <div class="table-toolbar">
          <el-popover v-model="showCreateTeacherPopover" placement="right" width="200">
            <remote-selector
              v-model="teacherCreateID"
              :params="{
                limit: 0,
                roles: 'teacher',
              }"
              size="mini"
              :get-list-method="GetUsersList"
              :formatter="(item)=>`${item.name} - ${item.username}`"
            />
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="showCreateTeacherPopover=false">取消</el-button>
              <el-button type="primary" size="mini" @click="handleCreateTeacher">确定</el-button>
            </div>
            <el-button slot="reference" size="mini" plain type="primary">{{ $t('New') }}</el-button>
          </el-popover>
        </div>
        <el-table stripe border :data="teachers" height="600px">
          <el-table-column prop="username" :label="$t('Account')" align="center"></el-table-column>
          <el-table-column prop="name" :label="$t('Name')" align="center"></el-table-column>
          <el-table-column prop="action" :label="$t('Action')" align="center">
            <template #default="{row}">
              <el-button
                type="warning"
                size="mini"
                plain
                @click="handleDeleteTeacher(row)"
              >{{ $t('Remove') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </el-form>
</template>

<i18n>
{
    "zh-CN": {
        "nameTip":"请输入班级名称"
    },
    "en-US": {
        "nameTip":"Input class name"
    }
}
</i18n>

<script>
import formMixin from "./formMixin";
import RichTextEditor from "@/components/form/RichTextEditor";
import RemoteSelector from "@/components/form/RemoteSelector";
import StudentForm from "./StudentForm";
import {
  GetClass,
  CreateClassTeacher,
  DeleteClassTeacher
} from "@/api/classes";
import { GetUsersList } from "@/api/users";
export default {
  name: "ClassForm",
  components: {
    RichTextEditor,
    RemoteSelector,
    StudentForm
  },
  mixins: [formMixin],
  props: {
    mode: {
      type: String,
      default: () => ""
    }
  },
  data() {
    return {
      GetUsersList,
      rules: {
        name: [{ required: true, message: "请输入班级名称" }]
      },
      loading: false,
      students: [],
      teachers: [],
      showCreateTeacherPopover: false,
      teacherCreateID: ""
    };
  },
  created() {
    this.$nextTick(() => {
      if (this.mode === "edit") {
        this.getData();
      }
    });
  },
  methods: {
    async getData() {
      const id = this.formData.id;
      this.loading = true;
      try {
        const res = await GetClass(id);
        this.teachers = res.data.teachers;
        this.students = res.data.students;
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },

    async handleCreateTeacher() {
      const id = this.formData.id;
      const uid = this.teacherCreateID;
      this.loading = true;
      try {
        await CreateClassTeacher(id, uid);
        this.showSuccessNotification();
        this.getData();
        this.teacherCreateID = "";
        this.showCreateTeacherPopover = false;
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    async handleDeleteTeacher(row) {
      this.loading = true;
      const id = this.formData.id;
      const uid = row.id;
      try {
        await DeleteClassTeacher(id, uid);
        this.showSuccessNotification();
        this.getData();
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    showSuccessNotification() {
      this.$notify({
        type: "success",
        title: this.$t("SuccessfulOperation")
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.table-toolbar {
  margin: 0 10px 10px 10px;
  display: flex;
  justify-content: center;
  .el-button {
    margin: 5px;
  }
}
.el-popover {
  .el-select {
    margin: 10px 0;
  }
}
</style>