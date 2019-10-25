import client from './client';

export const toggle = ({ mode, id, isLike }) => client.request({
  url: '/Reaction',
  method: 'put',
  data: { mode, id, reaction: (isLike ? 'like' : 'unlike') },
});

export default {
  toggle,
};
