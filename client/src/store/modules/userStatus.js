// import axios from 'axios'

const state = {
  logged: false,
  displayLogPanel: false
}
const getters = {
  userlogged: (state) => state.logged,
  displayLogPanel: (state) => state.displayLogPanel,
}
const mutations = {
  muta_login: (state) => {
    state.logged = true
  },
  muta_logout: (state) => {
    state.logged = false
  },
  muta_showLogPanel: (state) => {
    if (!state.displayLogPanel) {
      console.log("showing");
      
      state.displayLogPanel =true
    }
  },
  muta_hideLogPanel: (state) => {
    if (state.displayLogPanel) {
      console.log("hiding");
      state.displayLogPanel = false
    } 
  }
}
const actions = {
  act_login({ commit }) {
    console.log("loging in");
    commit('muta_login')
  },
  act_logout({ commit }) {
    console.log("loging out");
    commit('muta_logout')
  },
  act_showLogPanel({ commit }) {
    commit('muta_showLogPanel')
  },
  act_hideLogPanel({ commit }) {
    commit('muta_hideLogPanel')
  }
}
export default {
  state,
  getters,
  actions,
  mutations
};