import { get, invert, isEmpty, mapValues } from 'lodash';
import axios from 'axios';
import config from 'config';
import TokenUtil from 'utils/token';
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
      config.headers.Authorization = TokenUtil.getAccessToken();

      const cancelToken = get(vm, '$cancelToken');
      if (cancelToken) {
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
      if (axios.isCancel(error)) {
        return Promise.reject(Error('api canceled'));
      }
      const status = get(error, 'response.status');
      switch (status) {
        case 401: {
          const options = {
            title: '로그인이 필요합니다.',
            msg: '로그인 페이지로 이동합니다.',
          };
          vm.$toastr.w(options);
          vm.$router.push('/login');
        }
          break;
        case 403: {
          const options = {
            title: '요청이 실패하였습니다.',
            msg: '권한이 존재하지 않습니다.',
          };
          vm.$toastr.w(options);
        }
          break;
        default:
          if (!__DEV__) {
            const options = {
              title: `Unknown response. (status: ${status})`,
              msg: 'Unknown response.',
            };
            vm.$toastr.w(options);
          }
          break;
      }
      return Promise.reject(error);
    });
  }
  return client;
}

export const API = Object.freeze({
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
    this.cancelApi('page move');
  },
};
