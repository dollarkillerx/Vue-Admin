import Vue from 'vue'
import Router from 'vue-router'
import ro from "element-ui/src/locale/lang/ro";

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      redirect: '/index'
    },{
      path: '/index',
      name: 'index',
      component: () => import('./views/index.vue'),
      children: [
        {path: '',component: () => import('./views/home.vue')},
        {path: '/home',name: 'home',component: () => import('./views/home.vue')},
        {path: '/infoshow',name: 'infoshow',component: () => import('./views/infoShow.vue')},
        {path: '/sortlist',name: 'sortlist',component: () => import('./views/sortList.vue')}
      ]
    },{
      path: '/register',
      name: 'register',
      component: () => import('./views/register.vue')
    },{
      path: '/login',
      name: 'login',
      component: () => import('./views/login.vue')
    },{
      path: '*',
      name: '404',
      component: () => import('./views/404.vue')
    }
  ]
})

// 路由守卫
router.beforeEach((to,from,next) => {
  const isLogin = localStorage.DKToken ? true : false;
  if (to.path == "/login" || to.path == "/register") {
    next();
  }else {
    isLogin ? next() : next("/login");
  }
})

export default router;