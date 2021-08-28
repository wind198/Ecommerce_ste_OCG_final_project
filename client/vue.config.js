module.exports = {
  css: {
    loaderOptions: {
      scss: {
        prependData: `
          @import "@/scss/_color.scss";
          @import "@/scss/_mixin.scss";
        `
      }
    }
  }
}