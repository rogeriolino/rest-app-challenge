import Vue from 'vue'
import Router from 'vue-router'
import store from '@/store'
import Login from '@/pages/Login'
import Home from '@/pages/Home'

Vue.use(Router)

function requireAuth (to, from, next) {
  if (!store.getters.loggedIn) {
    next({
      name: 'login',
      query: { redirect: to.fullPath }
    })
  } else {
    next()
  }
}

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/logout',
      name: 'logout',
      beforeEnter (to, from, next) {
        store.dispatch('logout').then(() => {
          next({ name: 'login' })
        })
      }
    },
    {
      path: '/',
      name: 'home',
      component: Home,
      beforeEnter: requireAuth
    },
    {
      path: '*',
      beforeEnter (to, from, next) {
        next({ name: 'home' })
      }
    }
  ]
})
