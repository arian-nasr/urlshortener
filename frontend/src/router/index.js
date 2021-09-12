import Vue from 'vue';
import VueRouter from 'vue-router';

const routerOptions = [
  { path: '/', component: 'HelloWorld' },
  { path: '/index.html', component: 'HelloWorld' }
]

const routes = routerOptions.map(route => {
  return {
    ...route,
    component: () => import(`@/components/${route.component}.vue`)
  }
})

Vue.use(VueRouter);

const router = new VueRouter({
  routes,
  mode: 'history'
});

export default router;
