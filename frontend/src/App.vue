<template>
  <v-app>
    <navigation-drawer/>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import NavigationDrawer from './components/NavigationDrawer.vue';

export default {
  name: "App",
  components: {
    NavigationDrawer,
  },
  data: () => ({
  }),
};
</script>
