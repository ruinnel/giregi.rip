import { join } from 'lodash';
import proxy from 'api/proxy';

const my = (client, reporterIds) => client.request({
  url: '/memo/my',
  method: 'get',
  params: { reporterIds: join(reporterIds, ',') },
});

const get = (client, id) => client.request({
  url: `/memo/${id}`,
  method: 'get',
});

const search = (client, { userId, reporterId, offset, count }) => client.request({
  url: '/memo',
  method: 'get',
  params: { userId, reporterId, offset, count },
});

const create = (client, { reporterId, memo }) => client.request({
  url: '/memo',
  method: 'post',
  data: { reporterId, memo },
});

const update = (client, { id, memo }) => client.request({
  url: '/memo',
  method: 'put',
  data: { id, memo },
});

const remove = (client, id) => client.request({
  url: '/memo',
  method: 'delete',
  params: { id },
});

export default (client) => proxy({ my, get, search, create, update, remove }, client);
