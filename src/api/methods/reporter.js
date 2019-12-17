import proxy from 'api/proxy';

export const get = (client, id) => client.request({
  url: `/Reporter/${id}`,
  method: 'get',
});

export const search = (client, { agencyId, name, email, offset = 0, count = 10 }) => client.request({
  url: '/Reporter',
  method: 'get',
  params: { agencyId, name, email, offset, count },
});

export const create = (client, { agencyId, name, email, photoUrl, memo }) => client.request({
  url: '/Reporter',
  method: 'post',
  data: { agencyId, name, email, photoUrl, memo },
});

export const update = (client, { id, name, photoUrl, agencies }) => client.request({
  url: '/Reporter',
  method: 'put',
  data: { id, name, photoUrl, agencies },
});

export const remove = (client, id) => client.request({
  url: '/Reporter',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ get, search, create, update, remove }, client);
