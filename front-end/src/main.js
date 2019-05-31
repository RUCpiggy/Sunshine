// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import VueAxios from 'vue-axios'
import Axios from 'axios'
import vuex from 'vuex'

Vue.config.productionTip = false

/* eslint-disable no-new */
Vue.use(ElementUI)
Vue.use(VueAxios, Axios)
Vue.use(vuex)
Axios.defaults.baseURL = '/api'

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
