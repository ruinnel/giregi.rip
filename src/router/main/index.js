import Layout from 'layout/Main';
import Home from 'views/home/index';

const router = {
  path: '',
  component: Layout,
  children: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
  ],
};

export default router;
