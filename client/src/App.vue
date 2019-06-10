<template>
  <div id="app">
    <Nav />
    <router-view />
  </div>
</template>

<script>
import Nav from "./components/nav/Nav";

export default {
  name: "app",
  components: {
    Nav
  }
};
</script>

<style lang="scss"></style>
