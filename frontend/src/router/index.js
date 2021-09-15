import Vue from 'vue'
import Router from 'vue-router'
import VueAnalytics from 'vue-analytics'

const routerOptions = [
  { path: '/', component: 'Index' },
  { path: '/index.html', component: 'Index' },
  { path: '/login', component: 'Login' },
  { path: '/panel', component: 'Panel' }
]

const routes = routerOptions.map(route => {
  return {
    ...route,
    component: () => import(`@/components/${route.component}.vue`)
  }
})

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes
})

Vue.use(VueAnalytics, {
  id: 'UA-207682383-1',
  routes
})

export default router
