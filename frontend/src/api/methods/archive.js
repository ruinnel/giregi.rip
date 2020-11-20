import proxy from 'api/proxy';

export const preview = (client, url) => client.request({
  url: '/archives/preview',
  method: 'get',
  params: { url },
});

export const archive = (client, url, title, memo, tags, isPublic) => client.request({
  url: '/archives',
  method: 'post',
  data: { url, title, memo, tags, public: isPublic },
});

export const getByUrl = (client, url) => client.request({
  url: '/archives/url',
  method: 'get',
  params: { url },
});

export default (client) => proxy({ preview, archive, getByUrl }, client);
