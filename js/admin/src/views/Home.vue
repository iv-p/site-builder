<template>
  <div class="home">
    <Fragments />
  </div>
</template>

<script>
// @ is an alias to /src
import Fragments from '@/components/fragments/Fragments.vue';

export default {
  name: 'home',
  components: {
    Fragments,
  },
};
</script>
