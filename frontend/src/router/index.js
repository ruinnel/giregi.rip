import Vue from 'vue';
import VueRouter from 'vue-router';
import mainRouter from './main';
import fullRouter from './full';

Vue.use(VueRouter);

const routes = [
  mainRouter,
  fullRouter,
];

const router = new VueRouter({
  mode: process.env.NODE_ENV === 'electron' ? '' : 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
