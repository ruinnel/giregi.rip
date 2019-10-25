import client from './client';

export const get = (id) => client.request({
  url: `/Reporter/${id}`,
  method: 'get',
});

export const search = ({ agencyId, name, email, offset = 0, limit = 10 }) => client.request({
  url: '/Reporter',
  method: 'get',
  params: { agencyId, name, email, offset, limit },
});

export const create = ({ agencyId, name, email, photoUrl, memo }) => client.request({
  url: '/Reporter',
  method: 'post',
  data: { agencyId, name, email, photoUrl, memo },
});

export const update = ({ id, name, photoUrl, agencies }) => client.request({
  url: '/Reporter',
  method: 'put',
  data: { id, name, photoUrl, agencies },
});

export const remove = (id) => client.request({
  url: '/Reporter',
  method: 'delete',
  params: { id },
});

export default {
  get,
  search,
  create,
  update,
  remove,
};
