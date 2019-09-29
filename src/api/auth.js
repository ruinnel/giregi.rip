import client from './client';
import TokenUtil from 'utils/token';

export const check = () => client.request({
  url: '/Token',
  method: 'get',
});

export const login = (email, idToken) => client.request({
  url: '/Token',
  method: 'post',
  data: { email, idToken, tokenId: TokenUtil.getTokenId() },
})
  .then((res) => {
    const { result, data } = res;
    if (result && data) {
      TokenUtil.saveTokenId(data.id);
      TokenUtil.saveAccessToken(data.token);
    }
    return res;
  });

export const logout = () => client.request({
  url: '/Token',
  method: 'delete',
})
  .finally(() => {
    TokenUtil.clearAccessToken();
  });

export default {
  check,
  login,
  logout,
};
