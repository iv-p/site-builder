<template>
  <section class="section">
    <div class="columns">
      <div class="column">New Site</div>
      <div class="column">
        <div class="field is-expanded">
          <label class="label">Load Existing</label>
          <div class="control">
            <input
              class="input"
              type="file"
              accept="json"
              @change="handleChange"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  name: "Onboard",
  methods: {
    handleChange(e) {
      this.loading = true;
      var files = e.target.files || e.dataTransfer.files;
      var reader = new FileReader();
      reader.onload = async e => {
        this.site = JSON.parse(e.target.result);
        console.log(this.site);
      };
      reader.readAsText(files[0]);
    }
  },
  data() {
    return {
      site: {}
    };
  }
};
</script>
