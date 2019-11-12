import lodash from "lodash";
import xlsx from "xlsx";
const listMixin = {
  data() {
    return {
      removeProperties: [],
      show: false,
      mode: "",
      originalData: {},
      data: {},
      params: {
        limit: 10,
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
    },
    "params.limit"() {
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
          if (!this.removeProperties.includes(key)) {
            diffFound = !lodash.isEqual(this.data[key], this.originalData[key]);
            if (diffFound) {
              return true;
            }
          }
        }
        return false;
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
      this.getList();
    },
    handleDownloadSheet(sheetData, filename) {
      const workbook = xlsx.utils.book_new();
      const worksheet = xlsx.utils.json_to_sheet(sheetData);
      xlsx.utils.book_append_sheet(workbook, worksheet, "Sheet1");
      xlsx.writeFile(workbook, filename + ".xlsx");
    },
    makeUpdatedData() {
      let updateData = {};
      for (let key in this.originalData) {
        if (!this.removeProperties.includes(key)) {
          if (!lodash.isEqual(this.data[key], this.originalData[key])) {
            updateData[key] = this.data[key];
            if (typeof updateData[key] == "string") {
              updateData[key] = updateData[key].trim();
            }
          }
        }
      }
      return updateData;
    },
    showSuccess() {
      this.$notify({
        type: "success",
        title: "操作成功"
      });
    },
    showCancel() {
      this.$notify({
        type: "info",
        title: "已取消操作"
      });
      this.loading = false;
    }
  }
};

export default listMixin;
