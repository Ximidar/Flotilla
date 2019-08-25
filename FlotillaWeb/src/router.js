import Vue from 'vue'
import Router from 'vue-router'
import About from '@/views/About.vue'
import FilesVue from '@/views/Files/Files.vue'
import Status from '@/views/Status.vue'
import Util from '@/views/Util.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'About',
      component: About
    },
    // {
    //   path: '/about',
    //   name: 'About',
    //   component: About
    // },
    {
      path: '/files',
      name: 'Files',
      component: FilesVue
    },
    {
      path: '/util',
      name: 'Util',
      component: Util
    },
    {
      path: '/status',
      name: 'Status',
      component: Status
    }
  ]
})
