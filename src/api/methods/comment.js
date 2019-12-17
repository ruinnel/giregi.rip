import proxy from 'api/proxy';

const my = (client, newsId) => client.request({
  url: '/Comment/my',
  method: 'get',
  params: { newsId },
});

const get = (client, id) => client.request({
  url: `/Comment/${id}`,
  method: 'get',
});

const search = (client, { userId, newsId, offset, count }) => client.request({
  url: '/Comment',
  method: 'get',
  params: { userId, newsId, offset, count },
});

const create = (client, { newsId, comment }) => client.request({
  url: '/Comment',
  method: 'post',
  data: { newsId, comment },
});

const update = (client, { id, comment }) => client.request({
  url: '/Comment',
  method: 'put',
  data: { id, comment },
});

const remove = (client, id) => client.request({
  url: '/Comment',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ my, get, search, create, update, remove }, client);
