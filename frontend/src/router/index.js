import Vue from 'vue'
import VueRouter from 'vue-router'
import Brief from '../views/Brief.vue'
import Tag from '../views/Tag.vue'
import Locker from '../views/Locker.vue'
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
    path: '/locker',
    name: 'Locker',
    component: Locker
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