import { resize, toBlob, getSize } from "@/services/image";
import uuid from "uuid/v4";

const images = {
  namespaced: true,
  state: {
    sizes: {
      small: {
        width: 256,
        height: 256
      },
      medium: {
        width: 512,
        height: 512
      },
      large: {
        width: 1024,
        height: 1024
      }
    },
    original: {},
    resized: {
      // imageId: {
      //   meta: {
      //     name: "image_one.png",
      //     description: "desc",
      //     width: 960,
      //     height: 960
      //   },
      //   blob: "image contents"
      // }
    }
  },
  getters: {
    GET_ORIGINAL_IMAGE_META: state => id =>
      state.original[id] ? state.original[id].meta : null,
    GET_ORIGINAL_IMAGE: state => id =>
      state.original[id] ? state.original[id].image : null,
    GET_RESIZED_IMAGE: state => id =>
      state.resized[id] ? state.resized[id].image : null,
    GET_ORIGINAL_IMAGES_IDS: state => Object.keys(state.original),
    GET_ORIGINAL_IMAGE_URL: state => id =>
      state.original[id] ? URL.createObjectURL(state.original[id].image) : null
  },
  actions: {
    SAVE_IMAGE: async ({ commit }, { dataUrl, meta }) => {
      const id = uuid();
      const blob = toBlob(dataUrl);
      const size = await getSize(dataUrl);
      commit("SAVE_ORIGINAL_IMAGE", {
        id,
        image: blob,
        meta: {
          ...meta,
          ...size
        }
      });
    },
    RESIZE_IMAGE: async ({ commit, state }, { originalId, width, height }) => {
      const resizedId = uuid();
      const original = state.original[originalId];
      if (!originalId) {
        return;
      }

      const resized = await resize(original.blob, {
        width,
        height
      });
      commit("SAVE_RESIZED_IMAGE", {
        id: resizedId,
        image: resized,
        meta: {
          width,
          height
        }
      });
      return resizedId;
    }
  },
  mutations: {
    SAVE_ORIGINAL_IMAGE: (state, { id, image, meta }) => {
      state.original = {
        ...state.original,
        [id]: {
          id,
          meta,
          image
        }
      };
    },
    SAVE_RESIZED_IMAGE: (state, { id, image, meta }) => {
      state.resized[id] = {
        id,
        meta,
        image
      };
    }
  }
};

export default images;
