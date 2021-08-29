import { createStore } from 'vuex'
import allProduct from "./modules/allProduct"
import userStatus from "./modules/userStatus"

export default createStore({
 
  modules: {
    allProduct,
    userStatus
  }
})
