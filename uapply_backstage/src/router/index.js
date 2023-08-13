import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import Login from "../views/login/login.vue"
import Master from "../views/master/master.vue"
import SuperMaster from "../views/super_master/super_master.vue"

const routes = [
  {
    path: '/',
    redirect: '/login',
    // redirect: '/SuperMaster',
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/Master',
    name: 'Master',
    component: Master,
    redirect:"/Master/config",
    children:[
      {
        path:'config',
        name:"MasterConfig",
        component: ()=> import("@/views/master/master_pages/master_config")
      },
      {
        path:'new',
        name:"MasterNew",
        component: ()=> import("@/views/master/master_pages/master_new")
      },
      {
        path:'interviewed',
        name:"MasterInterviewed",
        component: ()=> import("@/views/master/master_pages/master_interviewed")
      },
      {
        path:'eliminate',
        name:"MasterEliminate",
        component: ()=> import("@/views/master/master_pages/master_eliminate")
      },
      {
        path:'admission',
        name:"MasterAdmission",
        component: ()=> import("@/views/master/master_pages/master_admission")
      },
      {
        path:'data',
        name:"MasterData",
        component: ()=> import("@/views/master/master_pages/master_data")
      },
    ]
  },
  {
    path: '/SuperMaster',
    name: 'SuperMaster',
    component: SuperMaster,
    redirect:"/SuperMaster/config",
    children:[
      {
        path:"config",
        name:"SuperConfig",
        component: ()=> import("@/views/super_master/super_master_pages/super_config.vue")
      },
      {
        path:"data",
        name:"SuperData",
        component: ()=> import("@/views/super_master/super_master_pages/super_data.vue")
      }
    ]
  },
  {
    path: '/error_pages',
    name: 'ErrorPages',
    redirect:"/error_pages/error_404",
    children:[
      {
        path:"error_404",
        name:"error_404",
        component: () => import("@/views/error_pages/error_404.vue")
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

export default router
