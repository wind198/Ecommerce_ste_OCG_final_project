<template>
  <section>
    <h2><a href="#">Some title</a></h2>
    <div
      class="article-list"
      v-for="product in productArray"
      :key="product.productID"
    >
      <Product
        :productName="product.productName"
        :productPrice="product.productPrice"
        :productImageSource="product.productImageSource"
      />
    </div>
  </section>
</template>
<script>
import Product from "./Product.vue"
export default {
  name: "Section",
  components: {
    Product
  },
  props: {
    productArray: Array
  }
}
</script>

<style lang="scss" scoped>
@media screen and (max-width: 768px) {
  .article-list {
    width: 50%;
    margin: auto;
  }
  h2 a {
    font-size: 1.2rem;
    display: block;
    text-align: center;
    color: black;
  }
}
</style>