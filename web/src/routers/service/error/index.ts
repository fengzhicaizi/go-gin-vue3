import type { RouteRecordRaw } from 'vue-router';
import Error404View from '@/components/pages/error/404.vue';

const routes: RouteRecordRaw[] = [
	{
		path: '/:pathMatch(.*)*',
		name: '404',
		component: Error404View,
		meta: {
			key: '404',
		},
	},
];

export default routes;
