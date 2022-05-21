import Vue from 'vue';
import Vuex from 'vuex';
import userModule from '@/store/moudles/userModule';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    userModule,
  },
  strict: process.env.NODE_ENV !== 'production',
});
