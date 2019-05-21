import Vue from 'vue';
import Vuex from 'vuex';
import fragments from './fragments';

import fragmentService from '../services/fragment';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    fragments: fragments(fragmentService),
  },
});
