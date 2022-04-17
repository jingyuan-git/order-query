// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
window.axios = require('axios');
import Vue from "vue";
import BootstrapVue from "bootstrap-vue";
import App from "./App";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import router from './router'

Vue.config.devtools = true;

Vue.use(BootstrapVue);

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
