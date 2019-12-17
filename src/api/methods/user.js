import proxy from 'api/proxy';

export const get = (client, { id }) => client.request({
  url: '/User',
  method: 'post',
  params: { id },
});

export default (client) => proxy({ get }, client);
