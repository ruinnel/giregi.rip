import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import firebaseUtil from 'utils/firebase';

import VueToastr from 'vue-toastr';

import Validator from 'components/Validator';
import { ValidationProvider, ValidationObserver } from 'utils/validator';
import FormatUtil from 'utils/format';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import '@fortawesome/fontawesome-free/css/all.min.css';

import './assets/scss/tabler.scss';

Vue.config.productionTip = false;
Vue.use(VueToastr, {
  defaultTimeout: 2000,
  defaultProgressBar: false,
  defaultPosition: 'toast-bottom-right',
});
Vue.component('ValidationProvider', ValidationProvider);
Vue.component('ValidationObserver', ValidationObserver);
Vue.component('Validator', Validator);

Vue.prototype.formatNumber = (num) => FormatUtil.formatNumber(num);
Vue.prototype.formatDate = (date, format) => FormatUtil.formatDate(date, format);
Vue.prototype.formatDateTime = (date, format) => FormatUtil.formatDateTime(date, format);

firebaseUtil.init();

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
