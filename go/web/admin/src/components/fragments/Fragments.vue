<template>
  <div>
    {{ fragments }}
    <ul>
      <li
        v-for="(fragment, index) in fragments"
        :key="index"
        @click="selectFragment(fragment.id)">
        {{ fragment.name }}
      </li>
    </ul>
    <Fragment :fragment="selectedFragment"/>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from 'vuex';
import Fragment from './Fragment.vue';

export default {
  name: 'Fragments',
  components: {
    Fragment,
  },
  created() {
    this.populateFragments();
  },
  computed: {
    ...mapGetters('fragments', {
      fragments: 'fragmentsMeta',
      selectedFragment: 'selectedFragment',
    }),
  },
  methods: {
    ...mapActions('fragments', [
      'populateFragments',
    ]),
    ...mapMutations('fragments', [
      'selectFragment',
    ]),
  },
};
</script>

<style>
</style>
