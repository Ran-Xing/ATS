// 本地缓存服务

const PREFIX = 'grs_';

// user 模块

const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 储存

const set = (key, value) => localStorage.setItem(key, value);

// 获取
const get = (key) => localStorage.getItem(key);

export default {
  // user 模块
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};
