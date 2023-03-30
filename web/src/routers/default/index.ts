import type { RouteRecordRaw } from 'vue-router';

import Login from '@/components/pages/auth/Login/index.vue';

const routes: RouteRecordRaw[] = [
	{
		path: '/login',
		name: 'login',
		component: Login,
	},
];

export default routes;
