import storangeService from '@/services/storangeService';
import userService from '@/services/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storangeService.get(storangeService.USER_TOKEN),
    // eslint-disable-next-line max-len
    UserInfo: storangeService.get(storangeService.USER_INFO) ? JSON.parse(storangeService.get(storangeService.USER_INFO)) : null,
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storangeService.set(storangeService.USER_TOKEN, token);
      // 更新vuex状态
      state.token = token;
    },
    SET_USER_INFO(state, UserInfo) {
      // 更新本地缓存
      storangeService.set(storangeService.USER_INFO, JSON.stringify(UserInfo));
      // 更新vuex状态
      state.UserInfo = UserInfo;
    },
  },
  actions: {
    register(context, { email, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ email, password })
          .then((respone) => {
            // 保存token
            context.commit('SET_TOKEN', respone.data.data.token);
            document.cookie = `token=${respone.data.data.token}`;
            return userService.info();
          })
          .then((respone) => {
            // 保存用户信息
            context.commit('SET_USER_INFO', respone.data.data.UserInfo);
            resolve(respone);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    login(context, { email, password }) {
      return new Promise((resolve, reject) => {
        userService.login({ email, password })
          .then((respone) => {
            // 保存token
            context.commit('SET_TOKEN', respone.data.data.token);
            document.cookie = `token=${respone.data.data.token}`;
            return userService.info();
          })
          .then((respone) => {
            // 保存用户信息
            context.commit('SET_USER_INFO', respone.data.data.UserInfo);
            resolve(respone);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    info(context, { email, password }) {
      return new Promise((resolve, reject) => {
        userService.info({ email, password })
          .then((respone) => {
            // 保存token
            context.commit('SET_TOKEN', respone.data.data.token);
            document.cookie = `token=${respone.data.data.token}`;
            return userService.info();
          })
          .then((respone) => {
            // 保存用户信息
            context.commit('SET_USER_INFO', respone.data.data.UserInfo);
            resolve(respone);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    logout({ commit }) {
      // 清除本地缓存
      storangeService.set(storangeService.USER_TOKEN, '');
      storangeService.set(storangeService.USER_INFO, '');
      // TODO 清楚 cookie
      // document.cookie = '';
      // 清除vuex状态
      commit('SET_TOKEN', '');
      commit('SET_USER_INFO', '');
      console.log('Logout success!');
      window.location.reload();
    },
  },
};

export default userModule;
