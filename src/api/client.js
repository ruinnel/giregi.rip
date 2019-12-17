import { isEmpty, get, mapValues, invert } from 'lodash';
import axios from 'axios';
import config from 'config';
import TokenUtil from 'utils/token';
import auth from 'api/methods/auth';
import comment from 'api/methods/comment';
import memo from 'api/methods/memo';
import news from 'api/methods/news';
import reaction from 'api/methods/reaction';
import reporter from 'api/methods/reporter';
import user from 'api/methods/user';

function initClient(vm) {
  const client = axios.create(config.axios);

  // request interceptor
  if (isEmpty(client.interceptors.request.handlers)) {
    client.interceptors.request.use((config) => {
      const accessToken = TokenUtil.getAccessToken();
      config.headers.Authorization = accessToken;

      const cancelToken = get(vm, '$cancelToken');
      if (cancelToken) {
        console.log('cancelToken - ', cancelToken);
        config.cancelToken = cancelToken;
      }
      if (__DEV__) {
        console.log('request - ', config);
      }
      return config;
    }, (error) => {
      return Promise.reject(error);
    });
  }

  // response interceptor
  if (isEmpty(client.interceptors.response.handlers)) {
    client.interceptors.response.use((response) => {
      const result = get(response, 'data.result');
      if (result) {
        return response.data.data;
      }
    }, (error) => {
      const status = get(error, 'response.status');
      switch (status) {
        case 401: {
          const options = {
            color: 'warning',
            title: '로그인이 필요합니다.',
            text: '로그인 페이지로 이동합니다.',
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
          };
          vm.$vs.notify(options);
        }
          break;
        default:
          if (!__DEV__) {
            const options = {
              color: 'warning',
              title: `Unknown response. (status: ${status})`,
              text: 'Unknown response.',
            };
            vm.$vs.notify(options);
          }
          break;
      }
      return Promise.reject(error);
    });
  }
  return client;
}

export const API = Object.freeze({
  AUTH: 'auth',
  COMMENT: 'comment',
  MEMO: 'memo',
  NEWS: 'news',
  REACTION: 'reaction',
  REPORTER: 'reporter',
  USER: 'user',
});
export default {
  created() {
    const client = initClient(this);
    const source = axios.CancelToken.source();
    this.$cancelToken = source.token;
    this.cancelApi = source.cancel;
    const apis = mapValues(invert(API), (val, key) => {
      switch (key) {
        case API.AUTH:
          return auth(client);
        case API.COMMENT:
          return comment(client);
        case API.MEMO:
          return memo(client);
        case API.NEWS:
          return news(client);
        case API.REACTION:
          return reaction(client);
        case API.REPORTER:
          return reporter(client);
        case API.USER:
          return user(client);
        default:
          return null;
      }
    });
    this.getApi = (name) => apis[name];
  },
  beforeDestroy() {
    console.log('mixin - beforeDestroy');
    this.$cancelToken.cancel();
  },
};
