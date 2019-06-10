<template>
  <section class="section">
    <div class="columns">
      <div class="column is-one-quarter-tablet is-one-fifth-widescreen">
        <FragmentList />
      </div>
      <div class="column">
        <Preview />
      </div>
      <div class="column is-one-quarter-tablet is-one-fifth-widescreen">
        <CreateFragment />
        <input v-model="fragment" />
        <EditFragment :fragmentId="fragment" v-if="fragment" />
      </div>
    </div>
  </section>
</template>

<script>
import CreateFragment from "@/components/fragment/Create.vue";
import EditFragment from "@/components/fragment/Edit.vue";
import Preview from "@/components/Preview.vue";
import FragmentList from "@/components/editor/FragmentList.vue";

export default {
  name: "Editor",
  components: {
    CreateFragment,
    EditFragment,
    Preview,
    FragmentList
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
