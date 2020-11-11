import proxy from 'api/proxy';
import TokenUtil from 'utils/token';

export const my = (client) => client.request({
  url: '/users',
  method: 'get',
});

export const profile = (client, id = null) => client.request({
  url: `/users/${id}`,
  method: 'get',
  params: { id },
});

export const tags = (client) => client.request({
  url: '/users/tags',
  method: 'get',
});

export const archives = (client, params) => client.request({
  url: '/users/archives',
  method: 'get',
  params: { ...params },
});

const login = (client, email, idToken) => client.request({
  url: '/users/login',
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
  url: '/user/logout',
  method: 'delete',
})
  .finally(() => {
    TokenUtil.clearAccessToken();
  });

export default (client) => proxy({ profile, my, tags, archives, login, logout }, client);
