import Vue from 'vue'
import VueRouter from 'vue-router'
import Brief from '../views/Brief.vue'
import Tag from '../views/Tag.vue'
import Device from '../views/Device.vue'
import Log from '../views/Log.vue'
import Personnel from '../views/Personnel.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Brief',
    component: Brief
  },
  {
    path: '/tag',
    name: 'Tag',
    component: Tag
  },
  {
    path: '/device',
    name: 'Device',
    component: Device
  },
  {
    path: '/personnel',
    name: 'Personnel',
    component: Personnel
  },
  {
    path: '/log',
    name: 'Log',
    component: Log
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router