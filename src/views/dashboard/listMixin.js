import { compareArray } from "@/utils";
const listMixin = {
  data() {
    return {
      show: false,
      mode: "",
      originalData: {},
      data: {},
      params: {
        limit: 8,
        page: 1
      },
      total: 0,
      loading: false,
      list: []
    };
  },
  created() {
    this.getList();
  },
  watch: {
    "params.page"() {
      this.getList();
    }
  },
  computed: {
    checkDiff() {
      if (this.mode === "create") {
        return true;
      } else {
        let diffFound = false;
        for (let key in this.originalData) {
          if (typeof this.data[key] !== "object") {
            if (this.data[key] !== this.originalData[key]) {
              diffFound = true;
            }
          } else if (Array.isArray(this.data[key])) {
            diffFound = !compareArray(this.data[key], this.originalData[key]);
          }
        }
        return diffFound;
      }
    }
  },
  methods: {
    clearData() {
      this.data = {};
      this.originalData = {};
      this.mode = "";
    },
    handleSearch() {
      this.params.page = 1;
      this.getList();
    },
    handleNew() {
      this.show = true;
      this.mode = "create";
    },
    handleEdit(data) {
      this.show = true;
      this.mode = "edit";
      // Use JSON.parse(JSON.stringify(data)) to deep copy a object
      this.originalData = JSON.parse(JSON.stringify(data));
      this.data = JSON.parse(JSON.stringify(data));
    },
    handleClose() {
      this.clearData();
      if (this.$refs.form) {
        this.$refs.form.resetFields();
      }
      this.show = false;
    },
    handleSave() {
      this.loading = true;
      if (this.mode === "edit") {
        const updateData = this.makeUpdatedData();
        this.handleUpdate(this.data.id, updateData);
      } else if (this.mode === "create") {
        this.$refs.form.checkFields(validateResult => {
          this.loading = false;
          if (validateResult) {
            this.handleCreate(this.data);
          } else {
            return;
          }
        });
      } else {
        this.loading = false;
        console.error("NO AVALIABLE OPERATION");
        this.handleClose();
      }
    },
    handleModify() {
      this.showSuccess();
      this.handleClose();
      this.handleSearch();
    },
    makeUpdatedData() {
      let updateData = {};
      for (let key in this.originalData) {
        if (this.data[key] !== this.originalData[key]) {
          updateData[key] = this.data[key];
          if (typeof updateData[key] == "string") {
            updateData[key] = updateData[key].trim();
          }
        }
      }
      return updateData;
    },
    showSuccess() {
      this.$message({
        type: "success",
        message: "操作成功"
      });
    },
    showCancel() {
      this.$message({
        type: "info",
        message: "已取消操作"
      });
      this.loading = false;
    }
  }
};

export default listMixin;
