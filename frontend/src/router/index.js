import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{
			path: '/',
			component: Dashboard
		},
		{
			path: '/calculator',
			component: () => import('../views/Calculator.vue')
		},
		{
			path: '/map',
			component: () => import('../views/Map.vue')
		},
	],
})

export default router