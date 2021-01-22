// CONFIGS
const makeBaseUrl = () => {
  if (process.env.NODE_ENV === 'electron') {
    return 'http://localhost:58899/api';
  } else {
    if (__DEV__) {
      return 'http://127.0.0.1:8000/api';
    } else {
      return '/api';
    }
  }
};

const defaultConfig = {
  axios: {
    baseURL: makeBaseUrl(),
    timeout: 30 * 1000,
    crossDomain: true,
  },
  archivePrefix: 'https://web.archive.org',
  archiveCheckPrefix: 'http://web.archive.org/save',
};

export default defaultConfig;
