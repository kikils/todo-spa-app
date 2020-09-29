import Vue from 'vue'
import firebase from 'firebase'
import Router from 'vue-router'

Vue.use(Router)

function loadComponent (view) {
  return () => import(/* webpackChunkName: "view-[request]" */ `@/components/${view}.vue`)
}

let router =  new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'todoApp',
      component: loadComponent('TodoApp'),
      meta: {requiresAuth: true},
    },
    {
      path: '/signin',
      name: 'signin',
      component: loadComponent('Signin')
    },
    {
      path: '/signup',
      name: 'signup',
      component: loadComponent('Signup')
    },
  ]
})

router.beforeEach((to, from, next) => {
    const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
    if (requiresAuth) {
        firebase.auth().onAuthStateChanged(function (user) {
          if (user) {
            next()
          } else {
            next({ name: 'signin' })
          }
        })
      } else {
        next()
      }
  })
  
  export default router