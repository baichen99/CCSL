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
    checkFields() {
      return new Promise((resolve, reject) => {
        for (let key in this.formData) {
          if (typeof this.formData[key] == "string") {
            this.formData[key] = this.formData[key].trim();
          }
        }
        this.$refs.form.validate(valid => {
          if (valid) {
            resolve();
          } else {
            reject();
          }
        })
      })
    },
    resetFields() {
      this.$refs.form.resetFields();
    }
  }
};

export default formMixin;
