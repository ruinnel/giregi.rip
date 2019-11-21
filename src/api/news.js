import client from './client';

export const preview = (url) => client.request({
  url: '/Preview',
  method: 'get',
  params: { url },
});

export const get = (id) => client.request({
  url: `/News/${id}`,
  method: 'get',
});

export const search = ({ reporterId, agencyId, offset, count }) => client.request({
  url: '/News',
  method: 'get',
  params: { reporterId, agencyId, offset, count },
});

export const create = ({ url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment }) => client.request({
  url: '/News',
  method: 'post',
  data: { url, title, reportedAt, lastUpdatedAt, agencyId, reporterId, reporterName, memo, comment },
});

export const update = ({ id, url }) => client.request({
  url: '/News',
  method: 'put',
  data: { id, url },
});

export const remove = (id) => client.request({
  url: '/News',
  method: 'delete',
  params: { id },
});

export default {
  preview,
  get,
  search,
  create,
  update,
  remove,
};
