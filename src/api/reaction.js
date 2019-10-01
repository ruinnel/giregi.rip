import client from './client';

export const get = ({ mode, id, isLike, isCancel }) => client.request({
  url: '/User',
  method: 'put',
  params: { mode, id, reaction: (isLike ? 'like' : 'unlike'), isCancel },
});

export default {
  get,
};
