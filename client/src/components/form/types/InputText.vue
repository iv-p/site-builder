<template>
  <div class="field">
    <label class="label">{{ field.label }}</label>
    <div class="control">
      <input
        class="input"
        type="text"
        placeholder="Normal input"
        @change="handleChange"
        v-model="value"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: "FormTypeText",
  props: ["field"],
  methods: {
    handleChange() {
      this.$emit("change", this.value);
    }
  },
  data() {
    return {
      value: ""
    };
  }
};
</script>
