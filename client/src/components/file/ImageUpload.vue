<template>
  <section class="section">
    <div class="columns">
      <div class="column">
        <div class="field">
          <label class="label">Description</label>
          <div class="control">
            <input type="text" class="input" v-model="form.description" />
          </div>
          <p class="help">
            This is a short summary of the image, used for SEO.
          </p>
        </div>

        <div class="field">
          <label class="label">Image</label>
          <div class="control">
            <input
              type="file"
              accept="image/*"
              class="input"
              @change="onFileChange"
            />
          </div>
          <p class="help">
            This is a short summary of the image, used for SEO.
          </p>
        </div>

        <div class="field">
          <div class="control">
            <button class="button is-primary" @click="upload">Upload</button>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="field">
          <label class="label">Preview</label>
          <div class="control">
            <img class="preview" :src="this.form.file" />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
// @ is an alias to /src
import { mapMutations, mapActions } from "vuex";

export default {
  name: "imageUpload",
  methods: {
    ...mapMutations({
      setData: "fragment/instance/SET_INSTANCE_DATA"
    }),
    ...mapActions({
      saveImage: "image/SAVE_IMAGE"
    }),
    onFileChange(e) {
      this.loading = true;
      var files = e.target.files || e.dataTransfer.files;
      var reader = new FileReader();
      reader.onload = async e => {
        this.form.file = e.target.result;
      };
      reader.readAsDataURL(files[0]);
    },
    upload() {
      this.saveImage({
        dataUrl: this.form.file,
        meta: {
          description: this.form.description
        }
      });
    }
  },
  data() {
    return {
      loading: false,
      form: {
        description: "",
        file: ""
      }
    };
  }
};
</script>
