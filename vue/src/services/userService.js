import requests from '@/utils/requests';

// 用户注册

const register = ({ email, password }) => requests.post('/api/auth/register', { email, password });

// 用户登录
const login = ({ email, password }) => requests.post('/api/auth/login', { email, password });

// 用户信息
const info = () => requests.get('/api/auth/info');

export default {
  register,
  login,
  info,
};
