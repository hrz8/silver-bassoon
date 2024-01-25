import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';

import ErrorPage from './pages/ErrorPage.vue';
import HomePage from './pages/HomePage.vue';
import OrdersPage from './pages/OrdersPage.vue';

const routes: Readonly<RouteRecordRaw[]> = [
  {path: '/', component: HomePage},
  {path: '/orders', component: OrdersPage},
  {path: '/error/:errorType', component: ErrorPage},
  {path: '/:pathMatch(.*)*', redirect: '/error/not-found'},
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export {router};
