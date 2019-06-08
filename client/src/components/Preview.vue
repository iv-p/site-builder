<template>
  <iframe width="1200" height="500" :srcdoc="GET_RENDERED('page1')"></iframe>
</template>

<script>
import Mustache from "mustache";
import layout from "@/templates/layouts/one";
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters("site", ["GET_RENDERED"])
  },
  data() {
    return {
      html: Mustache.render(layout.html, {
        header: "HEADER",
        content: "CONTENT",
        footer: "FOOTER"
      })
    };
  }
};
</script>
