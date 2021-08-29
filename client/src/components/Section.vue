<template>
  <section>
    <h2>
      <a href="#">{{ categoryName }}</a>
    </h2>
    <div class="section-window">
      <div
        class="article-list"
        :left="leftPosition"
        :style="{ left: left + 'px' }"
      >
        <Product
          v-for="product in productArray"
          :key="product.ProductID"
          :productName="product.productName"
          :productPrice="product.productPrice"
          :productImageSource="product.productImageSource"
        />
      </div>
    </div>
    <div class="wipe-handler">
      <i class="fas fa-chevron-left left" @click="wipe"></i>
      <i class="fas fa-chevron-right right" @click="wipe"></i>
    </div>
  </section>
</template>
<script>
import Product from "./Product.vue";
export default {
  name: "Section",
  data() {
    return {
      left: 0,
    };
  },
  props: {
    categoryName: String,
    productArray: Array,
  },
  components: {
    Product,
  },
  computed: {
    leftPosition() {
      return this.left;
    },
  },
  methods: {
    wipe(e) {
      let classList = e.target.classList;
      let direction;
      for (let i = 0; i < classList.length; i++) {
        direction =
          classList[i] == "left"
            ? "left"
            : classList[i] == "right"
            ? "right"
            : "";
      }
      if (direction != "") {
        this.left += direction == "left" ? 322 : -322;
      }
      if (this.left > 0) {
        this.left = 0;
      }
      if (this.left < -6118) {
        this.left = -6118;
      }
      
      console.log(this.left);
    },
  },
};
</script>

<style lang="scss" scoped>
@media screen and (max-width: 768px) {
  section {
    position: relative;
  }
  .section-window {
    width: 322px;
    margin: auto;
    overflow: hidden;
    .article-list {
      white-space: nowrap;
      position: relative;
      transition: all 0.2s linear;
    }
  }
  h2 a {
    font-size: 1.2rem;
    display: block;
    text-align: center;
    color: black;
  }
  .wipe-handler {
    position: absolute;
    top: 60%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 60%;
    display: flex;
    justify-content: space-between;
    font-size: 120%;
    color: rgba(0, 0, 0, 0.459);
    i {
      cursor: pointer;
      padding: 1rem;
    }
  }
}
</style>