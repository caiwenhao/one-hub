import { showError } from './common';
import axios from 'axios';
import { store } from '../store';
import { LOGIN } from 'store/actions';

const serializeParams = (params = {}) => {
  const parts = [];
  const encode = (value) => encodeURIComponent(value);

  Object.keys(params).forEach((key) => {
    const value = params[key];
    if (value === undefined || value === null) {
      return;
    }

    if (Array.isArray(value)) {
      if (value.length === 0) {
        return;
      }
      value.forEach((item) => {
        if (item === undefined || item === null) {
          return;
        }
        parts.push(`${encode(key)}=${encode(item)}`);
      });
      return;
    }

    if (typeof value === 'string' && value.trim() === '') {
      return;
    }

    parts.push(`${encode(key)}=${encode(value)}`);
  });

  return parts.join('&');
};

export const API = axios.create({
  // ... 其他代码 ...

  baseURL: import.meta.env.VITE_APP_SERVER || '/'
});

API.defaults.paramsSerializer = {
  serialize: serializeParams
};

API.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('user');
      store.dispatch({ type: LOGIN, payload: null });
      // window.location.href = '/login';
    }

    if (error.response?.data?.message) {
      error.message = error.response.data.message;
    }

    showError(error);
  }
);

export const LoginCheckAPI = axios.create({
  // ... 其他代码 ...

  baseURL: import.meta.env.VITE_APP_SERVER || '/'
});

LoginCheckAPI.defaults.paramsSerializer = {
  serialize: serializeParams
};
