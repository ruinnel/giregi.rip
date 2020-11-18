// CONFIGS
const defaultConfig = {
  axios: {
    baseURL: __DEV__ ? 'http://127.0.0.1:8000/api' : '/api',
    timeout: 30 * 1000,
    crossDomain: true,
  },
  archivePrefix: 'https://web.archive.org',
};

export default defaultConfig;
