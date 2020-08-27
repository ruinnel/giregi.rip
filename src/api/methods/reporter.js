import proxy from 'api/proxy';

export const get = (client, id) => client.request({
  url: `/reporter/${id}`,
  method: 'get',
});

export const search = (client, { agencyId, name, email, offset = 0, count = 10 }) => client.request({
  url: '/reporter',
  method: 'get',
  params: { agencyId, name, email, offset, count },
});

export const create = (client, { agencyId, name, email, photoUrl, memo }) => client.request({
  url: '/reporter',
  method: 'post',
  data: { agencyId, name, email, photoUrl, memo },
});

export const update = (client, { id, name, photoUrl, agencies }) => client.request({
  url: '/reporter',
  method: 'put',
  data: { id, name, photoUrl, agencies },
});

export const remove = (client, id) => client.request({
  url: '/reporter',
  method: 'delete',
  params: { id },
});

export const news = (client, id, { offset = 0, count = 10 }) => client.request({
  url: `/reporter/${id}/news`,
  method: 'get',
  params: { offset, count },
});

export default (client) => proxy({ get, search, create, update, remove, news }, client);
