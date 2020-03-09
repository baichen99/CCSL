<template>
  <div v-loading="loading">
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
      <el-button
        type="warning"
        size="mini"
        plain
        :disabled="selection.length<=0"
        @click="handleDeleteStudents"
      >{{ $t('Remove') }}</el-button>
      <el-button
        type="danger"
        size="mini"
        plain
        :disabled="selection.length<=0"
        @click="handleDeleteUsers"
      >{{ $t('Delete') }}</el-button>
    </div>

    <el-table
      stripe
      border
      :data="students"
      show-summary
      :summary-method="getSummaries"
      height="600px"
      @selection-change="handleSelectItems"
    >
      <el-table-column type="selection" width="45" align="center" fixed="left" />
      <el-table-column prop="username" :label="$t('Account')" align="center"></el-table-column>
      <el-table-column prop="name" :label="$t('Name')" align="center"></el-table-column>
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
  </div>
</template>

<script>
import xlsx from "xlsx";
import { CreateClassStudent, DeleteClassStudent } from "@/api/classes";
import { DeleteUser } from "@/api/users";
export default {
  name: "StudentsForm",
  props: {
    cid: {
      // Class ID
      type: String,
      required: true
    },
    getDataMethod: {
      type: Function,
      required: true
    },
    students: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      loading: false,
      selection: [],
      studentsToCreate: [],
      showUploadDialog: false,
      showCreateStudentPopover: false,
      studentCreateForm: {
        username: "",
        name: ""
      }
    };
  },
  methods: {
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
      this.loading = true;
      for (let item of this.studentsToCreate) {
        try {
          await CreateClassStudent(this.cid, {
            username: item.username,
            name: item.name
          });
          item.state = "saved";
        } catch (err) {
          item.state = "error";
        }
      }
      this.getDataMethod();
      this.loading = false;
    },
    handleCancleCreateStudents() {
      this.studentsToCreate = [];
      this.showUploadDialog = false;
    },
    async handleCreateStudent() {
      this.loading = true;
      try {
        await CreateClassStudent(this.cid, this.studentCreateForm);
        this.showSuccessNotification();
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
        this.getDataMethod();
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
    async handleDeleteStudents() {
      this.loading = true;
      try {
        const itemWarning = `此操作将会把${this.selection.length}名学生移出班级，不会删除其登录账号，是否继续`;
        await this.$confirm(itemWarning, "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        });
        for (const item of this.selection) {
          const uid = item.id;
          await DeleteClassStudent(this.cid, uid);
        }
      } catch (err) {
        console.error(err);
      } finally {
        this.getDataMethod();
        this.loading = false;
      }
    },
    async handleDeleteUsers() {
      this.loading = true;
      try {
        const itemWarning = `此操作将会删除${this.selection.length}名学生的登录账号，是否继续`;
        await this.$confirm(itemWarning, "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "error"
        });
        for (const item of this.selection) {
          await DeleteUser(item.id);
        }
      } catch (err) {
        console.error(err);
      } finally {
        this.getDataMethod();
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
      return ["", "合计", this.students.length + "人"];
    },
    handleSelectItems(value) {
      this.selection = value;
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