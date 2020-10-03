import proxy from 'api/proxy';

export const preview = (client, url) => client.request({
  url: '/archive/preview',
  method: 'get',
  params: { url },
});

export const archive = (client, url) => client.request({
  url: '/archive',
  method: 'post',
  params: { url },
});

export const getByUrl = (client, url) => client.request({
  url: '/archive/url',
  method: 'get',
  params: { url },
});

export default (client) => proxy({ preview, archive, getByUrl }, client);
