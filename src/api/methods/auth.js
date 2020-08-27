import TokenUtil from 'utils/token';
import proxy from 'api/proxy';

const check = (client) => client.request({
  url: '/token',
  method: 'get',
});

const login = (client, email, idToken) => client.request({
  url: '/token',
  method: 'post',
  data: { email, idToken, tokenId: TokenUtil.getTokenId() },
})
  .then((data) => {
    const { tokenId, token } = data;
    if (tokenId && token) {
      TokenUtil.saveTokenId(tokenId);
      TokenUtil.saveAccessToken(token);
    }
    return data;
  });

const logout = (client) => client.request({
  url: '/token',
  method: 'delete',
})
  .finally(() => {
    TokenUtil.clearAccessToken();
  });

export default (client) => proxy({ check, login, logout }, client);
