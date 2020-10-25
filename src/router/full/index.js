import Layout from 'layout/Full';
import Login from 'views/full/Login';

const router = {
  path: '',
  component: Layout,
  children: [
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
  ],
};

export default router;
