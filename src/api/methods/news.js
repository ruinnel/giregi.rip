import proxy from 'api/proxy';

const preview = (client, url) => client.request({
  url: '/Preview',
  method: 'get',
  params: { url },
});

const get = (client, id) => client.request({
  url: `/News/${id}`,
  method: 'get',
});

const search = (client, { reporterId, agencyId, offset, count }) => client.request({
  url: '/News',
  method: 'get',
  params: { reporterId, agencyId, offset, count },
});

const create = (client, { url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment }) => client.request({
  url: '/News',
  method: 'post',
  data: { url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment },
});

const update = (client, { id, url }) => client.request({
  url: '/News',
  method: 'put',
  data: { id, url },
});

const remove = (client, id) => client.request({
  url: '/News',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ preview, get, search, create, update, remove }, client);
