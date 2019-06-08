<template>
  <div>
    <div class="field">
      <label class="label">Template</label>
      <div class="control">
        <div class="select">
          <select v-model="template">
            <option disabled value="">Select template</option>
            <option v-for="t in templateIds" :key="t">
              {{ t }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <FormField
      v-for="prop in templateProps(template)"
      :key="prop.name"
      :field="prop"
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
  name: "CreateFragment",
  components: {
    FormField
  },
  computed: {
    ...mapGetters({
      templateIds: "fragment/template/GET_TEMPLATE_IDS",
      templateProps: "fragment/template/GET_TEMPLATE_PROPS"
    })
  },
  methods: {
    ...mapActions({
      createInstance: "fragment/instance/CREATE_INSTANCE"
    }),
    handleChange(value, prop) {
      this.data[prop.name] = value;
    },
    async handleSubmit() {
      const res = await this.createInstance({
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
      template: "",
      data: {}
    };
  }
};
</script>
