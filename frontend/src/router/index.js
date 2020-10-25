import Vue from 'vue'
import VueRouter from 'vue-router'
import Brief from '../views/Brief.vue'
import Tag from '../views/Tag.vue'
import Locker from '../views/Locker.vue'
import DoorEvent from '../views/DoorEvent.vue'
import TagRecord from '../views/TagRecord.vue'
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
    path: '/tag-record',
    name: 'TagRecord',
    component: TagRecord
  },
  {
    path: '/door-event',
    name: 'DoorEvent',
    component: DoorEvent
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router