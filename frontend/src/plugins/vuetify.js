import Vue from 'vue';
import Vuetify from 'vuetify';
import vuetify_ko from 'vuetify/es5/locale/ko'
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#ee44aa',
        secondary: '#424242',
        accent: '#82B1FF',
        error: '#FF5252',
        info: '#2196F3',
        success: '#4CAF50',
        warning: '#FFC107'
      },
    },
  },
  lang: {
    locales: {vuetify_ko},
    current: 'vuetify_ko',
  }
});
