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
        <div class="table-toolbar">
          <el-upload action :limit="1" :show-file-list="false" :before-upload="handleExcelUpload">
            <el-button size="mini" plain type="primary">上传Excel</el-button>
          </el-upload>
          <el-popover v-model="showCreateStudentPopover" placement="right" width="200">
            <el-form size="mini" :model="studentCreateForm" label-width="40px">
              <el-form-item :label="$t('Account')">
                <el-input v-model="studentCreateForm.username" />
              </el-form-item>
              <el-form-item :label="$t('Name')">
                <el-input v-model="studentCreateForm.name" />
              </el-form-item>
            </el-form>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="handleCancleCreateStudent">取消</el-button>
              <el-button type="primary" size="mini" @click="handleCreateStudent">确定</el-button>
            </div>
            <el-button slot="reference" size="mini" plain type="primary">{{ $t('New') }}</el-button>
          </el-popover>
        </div>

        <el-table
          stripe
          border
          :data="students"
          show-summary
          :summary-method="getSummaries"
          height="600px"
        >
          <el-table-column prop="username" :label="$t('Account')" align="center"></el-table-column>
          <el-table-column prop="name" :label="$t('Name')" align="center"></el-table-column>
          <el-table-column prop="action" :label="$t('Action')" align="center">
            <template #default="{row}">
              <el-button
                type="warning"
                size="mini"
                plain
                @click="handleDeleteStudent(row)"
              >{{ $t('Remove') }}</el-button>
              <el-button
                type="danger"
                size="mini"
                plain
                @click="handleDeleteUser(row)"
              >{{ $t('Delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog :visible.sync="showUploadDialog" :append-to-body="true">
          <el-table stripe border :data="studentsToCreate" height="400px">
            <el-table-column prop="username" :label="$t('Account')" align="center"></el-table-column>
            <el-table-column prop="name" :label="$t('Name')" align="center"></el-table-column>
            <el-table-column prop="state" :label="$t('State')" align="center">
              <template #default="{row}">
                <el-tag v-if="row.state==='unsave'" type="primary" size="mini">未保存</el-tag>
                <el-tag v-else-if="row.state==='saved'" type="success" size="mini">已保存</el-tag>
                <el-tag v-else-if="row.state==='error'" type="error" size="mini">保存出错</el-tag>
              </template>
            </el-table-column>
          </el-table>
          <span slot="footer">
            <el-button plain size="small" @click="handleCancleCreateStudents">{{ $t("Cancel") }}</el-button>
            <el-button
              plain
              size="small"
              type="primary"
              :loading="loading"
              @click="handleCreateStudents"
            >{{ $t("Confirm") }}</el-button>
          </span>
        </el-dialog>
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
import xlsx from "xlsx";
import formMixin from "./formMixin";
import RichTextEditor from "@/components/form/RichTextEditor";
import RemoteSelector from "@/components/form/RemoteSelector";
import {
  GetClass,
  CreateClassStudent,
  DeleteClassStudent,
  CreateClassTeacher,
  DeleteClassTeacher
} from "@/api/classes";
import { GetUsersList, DeleteUser } from "@/api/users";
export default {
  name: "ClassForm",
  components: {
    RichTextEditor,
    RemoteSelector
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
      teachers: [],
      students: [],
      studentsToCreate: [],
      showUploadDialog: false,
      showCreateStudentPopover: false,
      showCreateTeacherPopover: false,
      studentCreateForm: {
        username: "",
        name: ""
      },
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
    handleExcelUpload(file) {
      const reader = new FileReader();
      reader.onload = e => {
        const data = e.target.result;
        const workbook = xlsx.read(data, {
          type: "binary"
        });
        const worksheet = workbook.Sheets["平时成绩记分册"];
        const result = xlsx.utils.sheet_to_json(worksheet);
        this.studentsToCreate = result.map(item => ({
          name: item["姓名"],
          username: item["学号"],
          state: "unsave"
        }));
        this.showUploadDialog = true;
      };
      reader.readAsBinaryString(file);
      return false;
    },
    async handleCreateStudents() {
      const id = this.formData.id;
      this.loading = true;
      for (let item of this.studentsToCreate) {
        try {
          await CreateClassStudent(id, {
            username: item.username,
            name: item.name
          });
          item.state = "saved";
        } catch (err) {
          item.state = "error";
        }
      }
      this.loading = false;
    },
    handleCancleCreateStudents() {
      this.studentsToCreate = [];
      this.showUploadDialog = false;
    },
    async handleCreateStudent() {
      const id = this.formData.id;
      this.loading = true;
      try {
        await CreateClassStudent(id, this.studentCreateForm);
        this.showSuccessNotification();
        this.getData();
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
        this.handleCancleCreateStudent();
      }
    },
    handleCancleCreateStudent() {
      this.showCreateStudentPopover = false;
      this.studentCreateForm = {
        username: "",
        name: ""
      };
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
    async handleDeleteStudent(row) {
      this.loading = true;
      const id = this.formData.id;
      const uid = row.id;
      try {
        await DeleteClassStudent(id, uid);
        this.showSuccessNotification();
        this.getData();
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    async handleDeleteUser(row) {
      this.loading = true;
      const uid = row.id;
      try {
        await DeleteUser(uid);
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
    },
    getSummaries() {
      return ["合计", "", this.students.length + "人"];
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