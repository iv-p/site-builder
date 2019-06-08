<template>
  <section class="section">
    Select image
    <figure
      class="image is-128x128"
      v-for="id in imageIds"
      :key="id"
      @click="select(id)"
    >
      <img :src="imageUrl(id)" />
    </figure>
  </section>
</template>

<script>
// @ is an alias to /src
import { mapGetters } from "vuex";

export default {
  name: "imageSelector",
  computed: {
    ...mapGetters({
      imageIds: "image/GET_ORIGINAL_IMAGES_IDS",
      imageUrl: "image/GET_ORIGINAL_IMAGE_URL"
    })
  },
  methods: {
    select(id) {
      this.$emit("selected", id);
    }
  }
};
</script>
