import client from './client';

export const get = ({ id, reporterId, agencyId, name, email, offset, limit }) => client.request({
  url: '/Reporter',
  method: 'post',
  params: { id, reporterId, agencyId, name, email, offset, limit },
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
  create,
  update,
  remove,
};
