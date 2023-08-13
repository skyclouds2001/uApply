import { createRouter, createWebHashHistory } from 'vue-router'
import Index from "../views/pages/index.vue"
import Resume from "../views/pages/resume.vue"
import Apply from "../views/pages/apply.vue"
import Status from "../views/pages/status.vue"
import ErrorPage from "../views/pages/error_page.vue"

const routes = [
  {
    path: '/',
    name: 'index',
    // component: Index,
    component: () => import("@/views/pages/index.vue")
  },
  {
    path: '/resume',
    name: 'resume',
    component: Resume,
  },
  {
    path: '/apply',
    name: 'apply',
    component: Apply
  },
  {
    path: '/status',
    name: 'status',
    component: Status
  },
  {
    path: '/error_page',
    name: 'error',
    component: ErrorPage
  }
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

export default router
