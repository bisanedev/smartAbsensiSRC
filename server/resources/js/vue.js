import Vue from 'vue';
import VModal from 'vue-js-modal';
import Notifications from 'vue-notification';
import Croppa from 'vue-croppa';
import 'vue-croppa/dist/vue-croppa.css';
import vSelect from 'vue-select'
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
import moment from 'moment';

Vue.use(VModal,{ dialog: true, dynamic: true });
Vue.use(Notifications);
Vue.use(Croppa);

Vue.component('DatePicker', DatePicker);
Vue.component('v-select', vSelect);

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('DD/MM/YYYY')
  }
});

Vue.filter('leftPad',  function (value) {
    return ('00'+value).slice(-2);
});

import router from './router';

const app = new Vue({		
  router
}).$mount('#app');