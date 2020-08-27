import proxy from 'api/proxy';

const my = (client, newsId) => client.request({
  url: '/comment/my',
  method: 'get',
  params: { newsId },
});

const get = (client, id) => client.request({
  url: `/comment/${id}`,
  method: 'get',
});

const search = (client, { userId, newsId, offset, count }) => client.request({
  url: '/comment',
  method: 'get',
  params: { userId, newsId, offset, count },
});

const create = (client, { newsId, comment }) => client.request({
  url: '/comment',
  method: 'post',
  data: { newsId, comment },
});

const update = (client, { id, comment }) => client.request({
  url: '/comment',
  method: 'put',
  data: { id, comment },
});

const remove = (client, id) => client.request({
  url: '/comment',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ my, get, search, create, update, remove }, client);
