import client from './client';

export const my = (id) => client.request({
  url: `/Comment/${id}/my`,
  method: 'get',
});

export const get = (id) => client.request({
  url: `/Comment/${id}`,
  method: 'get',
});

export const search = ({ userId, newsId, offset, limit }) => client.request({
  url: '/Comment',
  method: 'get',
  params: { userId, newsId, offset, limit },
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
  my,
  get,
  search,
  create,
  update,
  remove,
};
