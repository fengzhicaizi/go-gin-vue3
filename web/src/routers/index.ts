import { createRouter, createWebHashHistory } from 'vue-router';
import defaultRoutes from './default/index';
import { serviceRoutes } from './service/index';

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [...defaultRoutes, ...serviceRoutes],
});

export default router;
