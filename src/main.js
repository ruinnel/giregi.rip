import Vue from 'vue';
import Bluebird from 'bluebird';
import App from './App.vue';
import router from './router';
import store from './store';
import Vuesax from 'vuesax';
import 'assets/sass/main.scss';
import 'vuesax/dist/vuesax.css';
import 'material-icons/iconfont/material-icons.css';
import firebaseUtil from 'utils/firebase';
import ApiClient from 'api/client';

window.Promise = Bluebird;
Vue.config.productionTip = false;

firebaseUtil.init();
Vue.use(Vuesax);
Vue.use(ApiClient);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#wrapper');
