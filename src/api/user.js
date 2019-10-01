import client from './client';

export const get = ({ id }) => client.request({
  url: '/User',
  method: 'post',
  params: { id },
});

export default {
  get,
};
