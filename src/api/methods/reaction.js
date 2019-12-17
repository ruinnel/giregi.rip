import proxy from 'api/proxy';

export const toggle = (client, { mode, id, isLike }) => client.request({
  url: '/Reaction',
  method: 'put',
  data: { mode, id, reaction: (isLike ? 'like' : 'unlike') },
});

export default (client) => proxy({ toggle }, client);
