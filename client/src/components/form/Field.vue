<template>
  <component
    v-bind:is="typeComponentMap[field.type]"
    v-on:change="handleChange"
    :field="field"
    :value="value"
  ></component>
</template>

<script>
import InputText from "./types/InputText";

export default {
  name: "FormField",
  components: {
    InputText
  },
  props: ["field", "value"],
  methods: {
    handleChange(value) {
      this.$emit("change", value);
    }
  },
  data() {
    return {
      typeComponentMap: {
        text: "InputText"
      }
    };
  }
};
</script>
