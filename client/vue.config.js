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
  },
  devServer: {
    proxy:  "http://localhost:5000/api"


}
}