// import uuid from "uuid/v4";
import mustache from "mustache";

const instances = {
  namespaced: true,
  state: {
    instances: {
      instance1: {
        id: "id",
        templateId: "template1",
        data: {
          content: {
            id: "uuid",
            value: "data/image/link/w/e",
            dynamic: false,
            name: "human friendly name"
          }
        }
      },
      instance2: {
        id: "id",
        templateId: "template1",
        data: {
          content: {
            id: "uuid",
            value: "data/image/link/w/e",
            dynamic: false,
            name: "human friendly name"
          }
        }
      },
      instance3: {
        id: "id",
        templateId: "template2",
        data: {}
      },
      instance4: {
        id: "id",
        templateId: "template2",
        data: {}
      },
      instance5: {
        id: "id",
        templateId: "template2",
        data: {}
      }
    }
  },
  getters: {
    GET_INSTANCE: state => id => state.instances[id],
    GET_RENDERED: (state, getters, rootState, rootGetters) => (
      id,
      defaults
    ) => {
      const fragment = getters["GET_INSTANCE"](id);
      if (!fragment) {
        return;
      }
      const template = rootGetters["fragment/template/GET_TEMPLATE"](
        fragment.templateId
      );
      if (!template) {
        return;
      }

      const content = {};
      Object.keys(fragment.data).map(key => {
        const value = fragment.data[key];
        if (!value.dynamic) {
          content[key] = value.value;
          return;
        }
        content[key] = defaults.content[value.id];
        return;
      });
      const data = { ...defaults, ...content };
      return {
        html: mustache.render(template.html, data),
        css: mustache.render(template.css, data),
        js: template.js
      };
    }
  },
  mutations: {
    SET_INSTANCE: (state, { id, instance }) => {
      state.instances[id] = instance;
    },
    SET_DYNAMIC_KEY: (state, { id, key_id, key, dynamic = true }) => {
      const instance = state.instances[id];
      if (!instance) {
        return;
      }
      instance.data[key] = {
        ...instance.data[key],
        id: key_id,
        dynamic
      };
      state.instances[id] = instance;
    },
    SET_INSTANCE_DATA: (state, { id, key, value }) => {
      const instance = state.instances[id];
      if (!instance) {
        console.log("fragment not found");
        return;
      }
      instance.data[key] = {
        ...instance.data[key],
        dynamic: false,
        value: value
      };
      state.instances[id] = instance;
    }
  },
  actions: {}
};

export default instances;
