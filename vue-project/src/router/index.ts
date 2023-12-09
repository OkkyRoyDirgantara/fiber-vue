import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateProduct from '../components/CreateProduct.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../views/DashboardView.vue'),
      children: [
        {
          path: '',
          name: 'dashboard-home',
          component: () => import('../views/DashboardHomeView.vue')
        },
        {
          path: '/products',
          name: 'products',
          component: () => import('../views/ProductsView.vue'),
          children: [
            {
              path: 'create',
              name: 'products-create',
              component: () => CreateProduct
            },
            {
              path: 'update-stock',
              name: 'products-update',
              component: () => import('../components/UpdateProduct.vue')
            }
          ]
        }
      ]
    }
  ]
})

export default router
