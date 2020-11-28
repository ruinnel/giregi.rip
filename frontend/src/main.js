import Vue from 'vue';
import App from './App.vue';
import Bluebird from 'bluebird';
import router from './router';
import store from './store';
import firebase from 'firebase/app';
import googleConfig from 'config/google-config';
import VueAnalytics from 'vue-analytics';

import VueToastr from 'vue-toastr';
import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';

import Validator from 'components/Validator';
import { ValidationProvider, ValidationObserver } from 'utils/validator';
import FormatUtil from 'utils/format';
import Dialog from 'components/dialog';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import '@fortawesome/fontawesome-free/css/all.min.css';

import './assets/scss/tabler.scss';

window.Promise = Bluebird;

Vue.config.productionTip = false;
Vue.use(VueToastr, {
  defaultTimeout: 2000,
  defaultProgressBar: false,
  defaultPosition: 'toast-bottom-right',
});
Vue.use(Dialog);
Vue.use(Loading);
Vue.component('ValidationProvider', ValidationProvider);
Vue.component('ValidationObserver', ValidationObserver);
Vue.component('Validator', Validator);

Vue.prototype.formatNumber = (num) => FormatUtil.formatNumber(num);
Vue.prototype.formatDate = (date, format) => FormatUtil.formatDate(date, format);
Vue.prototype.formatDateTime = (date, format) => FormatUtil.formatDateTime(date, format);

const { ga, firebase: firebaseConfig } = googleConfig;
firebase.initializeApp(firebaseConfig);
Vue.use(VueAnalytics, { id: ga, router });

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
