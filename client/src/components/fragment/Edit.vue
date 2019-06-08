<template>
  <div>
    <FormField
      v-for="prop in templateProps(template)"
      :key="prop.name"
      :field="prop"
      :value="data[prop.name]"
      v-on:change="handleChange($event, prop)"
    />

    <div class="field is-grouped">
      <div class="control">
        <button class="button is-primary" @click="handleSubmit">Submit</button>
      </div>
      <div class="control">
        <button class="button is-text" @click="handleCancel">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import FormField from "@/components/form/Field";

import { mapGetters, mapActions } from "vuex";
export default {
  name: "EditFragment",
  components: {
    FormField
  },
  props: ["fragmentId"],
  mounted() {
    this.fragment = this.getInstance(this.fragmentId);
    if (!this.fragment) {
      return;
    }
    this.data = Object.assign({}, this.fragment.data);
    this.template = this.fragment.template;
  },
  computed: {
    ...mapGetters({
      getInstance: "fragment/instance/GET_INSTANCE",
      templateProps: "fragment/template/GET_TEMPLATE_PROPS"
    })
  },
  methods: {
    ...mapActions({
      updateInstance: "fragment/instance/UPDATE_INSTANCE"
    }),
    handleChange(value, prop) {
      this.data[prop.name] = value;
    },
    async handleSubmit() {
      const res = await this.updateInstance({
        id: this.fragmentId,
        template: this.template,
        data: this.data
      });
      if (!res) {
        return;
      }
      this.clearForm();
      this.$emit("submit");
    },
    handleCancel() {
      this.clearForm();
      this.$emit("cancel");
    },
    clearForm() {
      this.template = "";
      this.data = {};
    }
  },
  data() {
    return {
      fragment: {},
      template: "",
      data: {}
    };
  }
};
</script>
