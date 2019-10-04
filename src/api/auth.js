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
  .then((data) => {
    console.log(data);
    const { id, token } = data;
    if (id && token) {
      TokenUtil.saveTokenId(id);
      TokenUtil.saveAccessToken(token);
    }
    return data;
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
