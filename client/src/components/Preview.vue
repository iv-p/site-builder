<template>
  <iframe width="1200" :srcdoc="GET_RENDERED('page1')"></iframe>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters("site", ["GET_RENDERED"])
  }
};
</script>
