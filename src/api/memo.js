import client from './client';

export const get = ({ id, userId, reporterId, offset, limit }) => client.request({
  url: '/Memo',
  method: 'post',
  params: { id, userId, reporterId, offset, limit },
});

export const create = ({ reporterId, memo }) => client.request({
  url: '/Memo',
  method: 'post',
  data: { reporterId, memo },
});

export const update = ({ id, memo }) => client.request({
  url: '/Memo',
  method: 'put',
  data: { id, memo },
});

export const remove = (id) => client.request({
  url: '/Memo',
  method: 'delete',
  params: { id },
});

export default {
  get,
  create,
  update,
  remove,
};
