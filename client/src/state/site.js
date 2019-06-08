import _ from "lodash";
import mustache from "mustache";

const site = {
  namespaced: true,
  state: {
    name: "site name",
    title: {},
    layout: {
      layoutId: "ampLayout",
      components: {
        header: "instance3",
        navbar: "instance4",
        footer: "instance5"
      }
    },
    defaults: {
      colors: {
        primary: "base color",
        secondary: "secondary color",
        accent: "accent color",
        success: "additional color",
        danger: "danger color",
        warning: "warning color"
      }
    }
  },
  getters: {
    GET_SITE: state => state,
    GET_RENDERED: (state, getters, rootState, rootGetters) => pageId => {
      const site = getters["GET_SITE"];
      if (!site) {
        return;
      }

      const sections = {};
      Object.keys(site.layout.components).forEach(key => {
        const fragmentId = site.layout.components[key];
        sections[key] = rootGetters["fragment/instance/GET_RENDERED"](
          fragmentId
        );
      });

      sections["content"] = rootGetters["page/GET_RENDERED"](pageId);
      const layoutTemplate = rootGetters["fragment/template/GET_TEMPLATE"](
        site.layout.layoutId
      );
      if (!layoutTemplate) {
        return;
      }

      const cssAndJs = _.reduce(
        sections,
        (result, section) => {
          if (!section) {
            return result;
          }
          result.css += section.css;
          result.js = [...result.js, ...section.js];
          return result;
        },
        { css: "", js: [] }
      );

      const css = cssAndJs.css;
      const js = new Set(cssAndJs.js);

      Object.keys(sections).forEach(key => {
        if (!sections[key]) {
          return;
        }
        sections[key] = sections[key].html;
      });
      const data = {
        ...sections,
        css,
        js
      };

      return mustache.render(layoutTemplate.html, data);
    }
  },
  actions: {},
  mutations: {}
};

export default site;
