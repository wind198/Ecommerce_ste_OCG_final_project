import axios from 'axios'

const state = {
  allProduct: {}
}
const getters = {
  allProduct: (state) => state.allProduct
}
const mutations = {
  muta_loadAllProduct: (state, allProduct) => {
    console.log(allProduct);
    state.hompageProduct = allProduct
  }
}
const actions = {
  async act_loadAllProduct: ({ commit }) => {
    const response = await axios.get("/home")
    commit('muta_loadAllProduct', response.data)
  }
}