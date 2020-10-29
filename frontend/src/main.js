import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'
import router from './router'
import { consoleBuildInfo } from 'vue-cli-plugin-build-info/plugin'

Vue.config.productionTip = false
consoleBuildInfo()

window.LeaveItUsApp = new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
