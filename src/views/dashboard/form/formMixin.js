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
  }
};

export default formMixin;
