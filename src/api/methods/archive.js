import proxy from 'api/proxy';

export const preview = (client, url) => client.request({
  url: '/archives/preview',
  method: 'get',
  params: { url },
});

export const archive = (client, url, memo) => client.request({
  url: '/archives',
  method: 'post',
  params: { url, memo },
});

export const getByUrl = (client, url) => client.request({
  url: '/archives/url',
  method: 'get',
  params: { url },
});

export default (client) => proxy({ preview, archive, getByUrl }, client);
