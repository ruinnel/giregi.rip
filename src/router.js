import Vue from 'vue';
import Router from 'vue-router';
import Home from 'views/home';
import LoginPopup from 'views/login/Popup';
import Reporter from 'views/reporter';
// import ReporterDetail from 'views/reporter/module/Detail';
import News from 'views/news';
import Mypage from 'views/mypage';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginPopup,
    },
    {
      path: '/reporter',
      name: 'reporter',
      component: Reporter,
    },
    {
      path: '/news',
      name: 'news',
      component: News,
    },
    {
      path: '/mypage',
      name: 'mypage',
      component: Mypage,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
    },
  ],
});
