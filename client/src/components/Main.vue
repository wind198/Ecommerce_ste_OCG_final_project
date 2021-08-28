<template>
  <main>
    <template v-for="category in allProduct" :key="category.id">
      <Section :productArray="allProduct[category]" />
    </template>
  </main>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

import Section from "./Section.vue"
export default {
  name: "Main",
  components: {
    Section,
  },
  methods: {
    ...mapActions(["act_loadAllProduct"])
  },
  computed: {
    ...mapGetters(["allProduct"])
  },
  created(){
      this.act_loadAllProduct()
  }
}
</script>

<style>
</style>