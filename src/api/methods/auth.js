import TokenUtil from 'utils/token';
import proxy from 'api/proxy';

const check = (client) => client.request({
  url: '/Token',
  method: 'get',
});

const login = (client, email, idToken) => client.request({
  url: '/Token',
  method: 'post',
  data: { email, idToken, tokenId: TokenUtil.getTokenId() },
})
  .then((data) => {
    const { id, token } = data;
    if (id && token) {
      TokenUtil.saveTokenId(id);
      TokenUtil.saveAccessToken(token);
    }
    return data;
  });

const logout = (client) => client.request({
  url: '/Token',
  method: 'delete',
})
  .finally(() => {
    TokenUtil.clearAccessToken();
  });

export default (client) => proxy({ check, login, logout }, client);
