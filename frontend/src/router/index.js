import Vue from 'vue'
import Router from 'vue-router'

const routerOptions = [
  { path: '/', component: 'HelloWorld' },
  { path: '/index.html', component: 'HelloWorld' },
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

export default new Router({
  routes,
  mode: 'history'
})
