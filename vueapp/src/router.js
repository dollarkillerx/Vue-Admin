import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      redirect: '/index'
    },{
      path: '/index',
      name: 'index',
      component: () => import('./views/index.vue')
    },{
      path: '/register',
      name: 'register',
      component: () => import('./views/register.vue')
    },{
      path: '*',
      name: '404',
      component: () => import('./views/404.vue')
    }
  ]
})
