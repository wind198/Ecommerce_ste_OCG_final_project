import axios from 'axios'

const state = {
  allProduct: {}
}
const getters = {
  allProduct: (state) => state.allProduct
}
const mutations = {
  muta_loadAllProduct: (state, allProduct) => {
    state.allProduct = allProduct
  }
}
const actions = {
  async act_loadAllProduct({ commit }) {
    console.log("hello");
    const response = await axios.get("/home")
    console.log(response.data);
    commit('muta_loadAllProduct',await response.data)
  }
}
export default {
  state,
  getters,
  actions,
  mutations
};