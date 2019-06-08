// import uuid from "uuid/v4";
import mustache from "mustache";
import uuid from "uuid/v4";

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
    CREATE_INSTANCE: (state, { template, data }) => {
      const id = uuid();
      state.instances[id] = {
        id,
        template,
        data
      };
    },
    UPDATE_INSTANCE: (state, { id, template, data }) => {
      state.instances = {
        ...state.instances,
        [id]: {
          id,
          template,
          data
        }
      };
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
    }
  },
  actions: {
    CREATE_INSTANCE: ({ commit }, { template, data }) => {
      commit("CREATE_INSTANCE", { template, data });
      return true;
    },
    UPDATE_INSTANCE: ({ commit, state }, { id, template, data }) => {
      if (!state.instances[id]) {
        return false;
      }
      commit("UPDATE_INSTANCE", { id, template, data });
      return true;
    }
  }
};

export default instances;
