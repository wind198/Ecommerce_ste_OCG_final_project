import Vuex from 'vuex'
import Vue from 'vue'
import allProduct from './modules/allProduct';

//Load Vuex
Vue.useAttrs(Vuex)
//Create store
export default new Vuex.Store({
    modules:{
        allProduct
    }
})