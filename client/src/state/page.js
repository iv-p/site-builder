const pages = {
  namespaced: true,
  state: {
    pages: {
      page1: {
        id: "page1",
        title: "",
        url: "",
        slug: "",
        data: {
          content: "w/e"
        },
        meta: {
          title: "title",
          something: "other"
        },
        fragments: ["instance1", "instance2"]
      }
    }
  },
  getters: {
    GET_PAGE: state => id => state.pages[id],
    GET_PAGE_SECTIONS: state => id => {
      state.pages[id] ? state.pages[id].fragments : null;
    },
    GET_RENDERED: (state, getters, rootState, rootGetters) => (
      id,
      siteDefaults
    ) => {
      const page = getters["GET_PAGE"](id);
      if (!page) {
        return;
      }

      const defaults = {
        ...siteDefaults,
        ...page.content
      };

      const renderedFragments = page.fragments
        .map(fragmentId => {
          return rootGetters["fragment/instance/GET_RENDERED"](
            fragmentId,
            defaults
          );
        })
        .filter(fragment => fragment);
      const mergedFragments = renderedFragments.reduce(
        (result, fragment) => {
          result.html += fragment.html;
          result.css += fragment.css;
          result.js = [...result.js, ...fragment.js];
          return result;
        },
        { html: "", css: "", js: [] }
      );

      return mergedFragments;
    }
  },
  actions: {},
  mutations: {}
};

export default pages;
