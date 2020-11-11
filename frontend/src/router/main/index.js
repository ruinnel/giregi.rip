import Layout from 'layout/Main';
import Home from 'views/home/index';
import MyArchiveList from 'views/archive/my';

const router = {
  path: '',
  component: Layout,
  children: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/archives/my',
      name: 'archiveList',
      component: MyArchiveList,
    },
  ],
};

export default router;
