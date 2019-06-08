<template>
  <div class="field">
    <label class="label">{{ field.label }}</label>
    <div class="control">
      <input
        class="input"
        type="text"
        placeholder="Normal input"
        @change="handleChange"
        v-model="v"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: "FormTypeText",
  props: ["field", "value"],
  methods: {
    handleChange() {
      this.$emit("change", this.v);
    }
  },
  data() {
    return {
      v: this.value
    };
  }
};
</script>
