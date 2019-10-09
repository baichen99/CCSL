const formMixin = {
  model: {
    prop: "data",
    event: "update"
  },
  props: {
    data: {
      type: Object,
      default: () => ({})
    }
  },
  computed: {
    formData: {
      get() {
        return this.data;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  methods: {
    checkFields(callback) {
      for (let key in this.formData) {
        if (typeof this.formData[key] == "string") {
          this.formData[key] = this.formData[key].trim();
        }
      }
      return this.$refs.form.validate(valid => {
        if (valid) {
          return callback(true);
        } else {
          return callback(false);
        }
      });
    },
    resetFields() {
      this.$refs.form.resetFields();
    }
  }
};

export default formMixin;
