import proxy from 'api/proxy';

const preview = (client, url) => client.request({
  url: '/preview',
  method: 'get',
  params: { url },
});

const get = (client, id) => client.request({
  url: `/news/${id}`,
  method: 'get',
});

const search = (client, { reporterId, agencyId, offset, count }) => client.request({
  url: '/news',
  method: 'get',
  params: { reporterId, agencyId, offset, count },
});

const report = (client, { url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment }) => client.request({
  url: '/news',
  method: 'post',
  data: { url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment },
});

const update = (client, { id, url }) => client.request({
  url: '/news',
  method: 'put',
  data: { id, url },
});

const remove = (client, id) => client.request({
  url: '/news',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ preview, get, search, report, update, remove }, client);
