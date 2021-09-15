import Vue from 'vue'
import App from './App'
import router from './router'
import 'bootstrap/dist/css/bootstrap.css'
import VueAnalytics from 'vue-analytics'

Vue.config.productionTip = false
Vue.use(VueAnalytics, {
  id: 'UA-207682383-1'
})
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
