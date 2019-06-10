<template>
  <aside class="menu">
    <p class="menu-label">
      SITE
    </p>
    <ul class="menu-list">
      <li><a>General</a></li>
      <li v-for="(section, sectionName) in siteSections" :key="sectionName">
        <a>{{ sectionName | capitalize }}</a>
      </li>
    </ul>
    <p class="menu-label">
      PAGE
    </p>
    <ul class="menu-list">
      <li v-for="section in pageSections(selectedPage)" :key="section">
        <a>{{ section }}</a>
      </li>
      <li>
        <a>Add New</a>
      </li>
    </ul>
  </aside>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "FragmentList",
  computed: {
    ...mapGetters({
      siteSections: "site/GET_SITE_SECTIONS",
      siteLayout: "site/GET_SITE_LAYOUT",
      pageSections: "page/GET_PAGE_SECTIONS"
    })
  },
  data() {
    return {
      selectedPage: "page1"
    };
  }
};
</script>
