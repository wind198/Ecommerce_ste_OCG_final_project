<template>
  <article>
    <div class="picture-holder">
      <img
        :src="product.productImageSource"
        :alt="product.productName"
      />
    </div>
    <div class="wipe-handler">
      <i class="fas fa-chevron-left"></i>
      <i class="fas fa-chevron-right"></i>
    </div>
  </article>
</template>

<script>
export default {
  name: "Article",
  props:{
    productName:String,
    productPrice:Number,
    productImageSource:String,
  }
}
</script>

<style lang="scss" scoped>
@media only screen and (max-width: 768px) {
  article {
    width: 90%;
    margin: auto;
    position: relative;
    border: 6px inset rgb(59, 59, 59);
    .picture-holder {
      img {
        width: 100%;
        max-width: 300px;
        height: auto;
      }
    }
    .wipe-handler {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      width: 100%;
      display: flex;
      justify-content: space-between;
      font-size: 120%;
      color:rgba(0, 0, 0, 0.459);
      i{
          cursor: pointer;
      }
    }
  }
}
</style>