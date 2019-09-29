import { isEmpty, get } from 'lodash';
import axios from 'axios';
import config from 'config';
import TokenUtil from 'utils/token';

function initClient(vm) {
  const client = axios.create(config.axios);

  // request interceptor
  if (isEmpty(client.interceptors.request.handlers)) {
    client.interceptors.request.use((config) => {
      const accessToken = TokenUtil.getAccessToken();
      config.headers.Authorization = accessToken;
      return config;
    }, (error) => {
      return Promise.reject(error);
    });
  }

  // response interceptor
  if (isEmpty(client.interceptors.response.handlers)) {
    client.interceptors.response.use((response) => {
      return response.data;
    }, (error) => {
      const status = get(error, 'response.status');
      switch (status) {
        case 401: {
          const options = {
            color: 'warning',
            title: '로그인이 필요합니다.',
            text: '로그인 페이지로 이동합니다.',
            vm,
          };
          vm.$vs.notify(options);
          vm.$router.push('/login');
        }
          break;
        case 403: {
          const options = {
            color: 'warning',
            title: '요청이 실패하였습니다.',
            text: '권한이 존재하지 않습니다.',
            vm,
          };
          vm.$vs.notify(options);
        }
          break;
        default: {
          const options = {
            color: 'warning',
            title: `Unknown response. (status: ${status})`,
            text: 'Unknown response.',
            vm,
          };
          vm.$vs.notify(options);
        }
          break;
      }
    });

    return client;
  }
}

let Vue; // bind on install
let client;

function init() {
  client = initClient(this);
}

const install = (_Vue) => {
  if (Vue && _Vue === Vue) {
    console.error('[vuex] already installed. Vue.use(Vuex) should be called only once.');
    return;
  }

  Vue = _Vue;
  Vue.mixin({ created: init });
};
// export default initClient(config.axios);
export default {
  install, // call from Vue
  get request() {
    return client;
  },
};
