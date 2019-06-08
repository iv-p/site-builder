<template>
  <section class="section">
    <CreateFragment />
    <input v-model="fragment" />
    <EditFragment :fragmentId="fragment" v-if="fragment" />
  </section>
</template>

<script>
import CreateFragment from "@/components/fragment/Create.vue";
import EditFragment from "@/components/fragment/Edit.vue";

export default {
  name: "home",
  components: {
    CreateFragment,
    EditFragment
  },
  methods: {
    handleSelected(id) {
      id;
    }
  },
  data() {
    return {
      fragment: null
    };
  }
};
</script>
