export default fragmentService => ({
  namespaced: true,
  state: {
    fragments: {},
    selectedFragmentId: null,
  },
  getters: {
    fragmentsMeta: store => Object.keys(store.fragments).map(key => ({
      name: store.fragments[key].name,
      id: store.fragments[key].id,
    })),
    selectedFragment: store => store.fragments[store.selectedFragmentId],
  },
  mutations: {
    setFragments(state, fragments) {
      state.fragments = fragments;
    },
    selectFragment(state, fragmentId) {
      state.selectedFragmentId = fragmentId;
    },
  },
  actions: {
    populateFragments: async (context) => {
      const fragments = await fragmentService.getFragments();
      context.commit('setFragments', fragments);
    },
  },
});
