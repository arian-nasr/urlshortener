import Vue from 'vue'
import App from './App'
import router from './router'
import 'bootstrap/dist/css/bootstrap.css'
import VueGtag from 'vue-gtag'

Vue.config.productionTip = false
/* eslint-disable no-new */
Vue.use(VueGtag, {
  config: { id: 'UA-207682383-1' },
  bootstrap: false
}, router)

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
