<template>
  <article>
    <div class="picture-holder">
      <img :src="productImageSource" :alt="productName" />
    </div>

    <div class="product-info">
      <p class="product-name">{{ productName }}</p>
      <p class="product-price">OurPrice: {{ productPrice }} $</p>
      <div class="buy-decision">
        <div class="console">
          <button class="minus">-</button>
          <span class="number">0</span>
          <button class="buy-more">+</button>
        </div>
        <div class="buy"><button>Buy</button></div>
      </div>
    </div>
  </article>
</template>

<script>
export default {
  name: "Article",
  props: {
    productName: String,
    productPrice: Number,
    productImageSource: String,
  },
  data() {
    return {
      articleListPosition: 0,
    };
  },
  
};
</script>

<style lang="scss" scoped>
@media only screen and (max-width: 768px) {
  article {
    display: inline-block;
    white-space: normal;
    width: 302px;
    height: 302px;
    margin: 0 10px;
    position: relative;
    border: 1px solid rgba(0, 0, 0, 0.233);
    .picture-holder {
      width: 100%;
      height: 100%;
      img {
        display: block;
        height: 100%;
        max-height: 300px;
        width: auto;
        max-width: 300px;
        margin: auto;
      }
    }

    .product-info {
      background: rgba(0, 0, 0, 0.4);
      border-radius: 10px 15px 0 0;
      color: white;
      position: absolute;
      display: none;
      bottom: 0px;
      left: 50%;
      transform: translateX(-50%);
      width: 100%;
      padding: 0.2rem 0.5rem;
      p.product-name {
        // text-decoration: underline;
        font-size: 1rem;
        // font-weight: bolder;
      }
      p.product-price {
        font-size: 0.9rem;
      }
      .buy-decision {
        display: grid;
        grid-template-columns: auto auto;
        width: 100%;
        button {
          border: none;
          border-radius: 5px;
        }
        .console {
          border-radius: 7.5px;
          padding: 4px;
          height: 30px;
          box-sizing: border-box;
          background: rgba(0, 0, 0, 0.452);
          text-align: center;
          display: grid;
          grid-template-columns: 35% 30% 35%;
          
          span {
            font-size: 0.8rem;
            line-height: 22px;
          }
        }
        .buy {
          button {
            background-color: rgb(221, 178, 86);
            width: 100%;
            margin: 0 2px;
            color: rgba(0, 0, 0, 0.6);
            font-weight: bold;
            height: 30px;
          }
        }
      }
    }
    &:hover {
      .product-info {
        display: block;
      }
    }
  }
}
</style>