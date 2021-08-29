<template>
  <main>
    <Section v-for="category in allProduct" :key="category.categoryID" :categoryName="category.categoryName" :productArray="category.productArray" />
  </main>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

import Section from "./Section.vue";
export default {
  name: "Main",
  components: {
    Section,
  },
  methods: {
    ...mapActions(["act_loadAllProduct"]),
  },
  computed: {
    //allProduct
    //[{id:xxxx, product:[{},{},...]},...]
    ...mapGetters(["allProduct"]),
  },
  created() {
    console.log("getting data after created");
    this.act_loadAllProduct();
  },
  watch:{
    allProduct(){

    }
  }
  
};
</script>

<style>
</style>