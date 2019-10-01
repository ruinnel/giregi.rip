import client from './client';

export const preview = (url) => client.request({
  url: '/Preview',
  method: 'get',
  params: { url },
});

export const get = ({ id, reporterId, agencyId, offset, limit }) => client.request({
  url: '/News',
  method: 'post',
  params: { id, reporterId, agencyId, offset, limit },
});

export const create = ({ url, agencyId, reporterId, memo, comment }) => client.request({
  url: '/News',
  method: 'post',
  data: { url, agencyId, reporterId, memo, comment },
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
  create,
  update,
  remove,
};
