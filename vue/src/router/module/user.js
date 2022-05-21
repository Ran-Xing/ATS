const userRoutes = [
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "login" */ '@/views/UserLogin/UserLogin.vue'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import(/* webpackChunkName: "about" */ '@/views/UserRegister/UserRegister.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    component: () => import(/* webpackChunkName: "about" */ '@/views/Profile/Profile.vue'),
  },
];

export default userRoutes;
