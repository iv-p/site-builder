<template>
  <section class="section">
    <CreateFragment />
  </section>
</template>

<script>
import CreateFragment from "@/components/fragment/Create.vue";

export default {
  name: "home",
  components: {
    CreateFragment
  },
  methods: {
    handleSelected(id) {
      id;
    }
  }
};
</script>
