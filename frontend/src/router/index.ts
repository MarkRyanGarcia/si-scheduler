import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Schedules from '../views/Schedules.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', name: 'Home', component: Home },
  { path: '/about', name: 'About', component: About },
  { path: '/schedules', name: 'Schedules', component: Schedules },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router