const TOKEN_ID_KEY = 'token_id';
const ACCESS_TOKEN_KEY = 'access_token';

const getAccessToken = () => localStorage.getItem(ACCESS_TOKEN_KEY);
const saveAccessToken = (accessToken) => localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
const clearAccessToken = () => localStorage.removeItem(ACCESS_TOKEN_KEY);

const getTokenId = () => localStorage.getItem(TOKEN_ID_KEY);
const saveTokenId = (tokenId) => localStorage.setItem(TOKEN_ID_KEY, tokenId);
const clearTokenId = () => localStorage.removeItem(TOKEN_ID_KEY);

export default {
  getAccessToken,
  saveAccessToken,
  clearAccessToken,
  getTokenId,
  saveTokenId,
  clearTokenId,
};
