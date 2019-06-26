import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    totalPrice: 0,
    tocken: '',
    foot: []
  },
  getters: {
    getTotal (state) {
      return state.totalPrice * 2
    }
  },
  mutations: {
    addTocken (state, tocken) {
      state.tocken = tocken
    },
    addFoot (state, page) {
      for (var i = 0; i < state.foot.length; i++) {
        if (state.foot[i] === page) {
          var newFoot = state.foot.slice(0, i + 1)
        }
      }
      if (newFoot != null) {
        state.foot = newFoot
      } else {
        state.foot.push(page)
      }
    },
    increment (state, price) {
      state.totalPrice += price
    },
    decrement (state, price) {
      state.totalPrice -= price
    }
  },
  actions: {
    increase (context, price) {
      context.commit('increment', price)
    }
  }
})
