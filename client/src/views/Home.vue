<template>
  <section class="section">
    <ImageUpload />
    <ImageSelector @selected="handleSelected" />
    <Preview />
  </section>
</template>

<script>
// @ is an alias to /src
import Preview from "@/components/Preview.vue";
import ImageUpload from "@/components/file/ImageUpload.vue";
import ImageSelector from "@/components/file/ImageSelector.vue";

export default {
  name: "home",
  components: {
    Preview,
    ImageUpload,
    ImageSelector
  },
  methods: {
    handleSelected(id) {
      id;
    }
  }
};
</script>
