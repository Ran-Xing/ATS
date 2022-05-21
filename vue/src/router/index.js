import Vue from 'vue';
import VueRouter from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import userRoutes from '@/router/module/user';
import store from '@/store/index';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  // 判断 是否需要登录
  if (to.meta.auth) {
    // 判断是否登录
    if (store.state.userModule.token) {
      // 判断时效性
      next();
    } else {
      router.push({ name: 'home' });
    }
  } else {
    next();
  }
});

export default router;
