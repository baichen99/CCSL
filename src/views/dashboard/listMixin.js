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
        for (let key in this.data) {
          if (this.originalData[key] !== this.data[key]) {
            diffFound = true;
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
    handleCreate() {
      this.mode = "create";
      this.show = true;
    },
    handleEdit(data) {
      this.mode = "edit";
      this.show = true;
      // Use JSON.parse(JSON.stringify(data)) to deep copy a object
      this.originalData = JSON.parse(JSON.stringify(data));
      this.data = JSON.parse(JSON.stringify(data));
    },
    handleClose() {
      this.show = false;
      this.clearData();
    },
    makeUpdatedData() {
      let updateData = {};
      for (let key in this.data) {
        if (this.originalData[key] !== this.data[key]) {
          updateData[key] = this.originalData[key];
        }
      }
      return updateData;
    }
  }
};

export default listMixin;
