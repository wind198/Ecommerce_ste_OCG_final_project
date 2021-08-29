<template>
  <Header />
  <Main />
  <Footer />
</template>
<script>
import Header from "./components/Header.vue";
import Main from "./components/Main.vue";
import Footer from "./components/Footer.vue";
export default {
  name: "App",
  components: {
    Header,
    Main,
    Footer,
  },
};
</script >

<style lang="scss" >
* {
  box-sizing: border-box;
  font-family: "Poppins", "sans-serif";
}
html {
  a {
    @include linkReset;
  }
  body {
    margin: 0;
    padding: 0;
  }
  button {
    cursor: pointer;
  }
  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
  }
}
</style>