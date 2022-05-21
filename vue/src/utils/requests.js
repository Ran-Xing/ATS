// axios
import axios from 'axios';
import storangeService from '@/services/storangeService';

const service = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  timeout: 1000 * 5,
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  // Do something before request is sent
  Object.assign(config.headers, {
    Authorization: `Bearer ${storangeService.get(storangeService.USER_TOKEN)}`,
  });
  return config;
}, (error) => Promise.reject(error));

export default service;
