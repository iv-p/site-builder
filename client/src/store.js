import Vue from "vue";
import Vuex from "vuex";

import site from "./state/site";
import page from "./state/page";
import image from "./state/image";
import fragment from "./state/fragment";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    site,
    page,
    image,
    fragment
  }
});
