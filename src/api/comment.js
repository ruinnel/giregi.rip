import client from './client';

export const get = ({ id, userId, newsId, offset, limit }) => client.request({
  url: '/Comment',
  method: 'post',
  params: { id, userId, newsId, offset, limit },
});

export const create = ({ newsId, comment }) => client.request({
  url: '/Comment',
  method: 'post',
  data: { newsId, comment },
});

export const update = ({ id, comment }) => client.request({
  url: '/Comment',
  method: 'put',
  data: { id, comment },
});

export const remove = (id) => client.request({
  url: '/Comment',
  method: 'delete',
  params: { id },
});

export default {
  get,
  create,
  update,
  remove,
};
