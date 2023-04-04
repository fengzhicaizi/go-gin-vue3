import type { RouteRecordRaw } from 'vue-router';
import HomeView from '@/views/Home/index.vue';

const route: RouteRecordRaw = {
	path: '/home',
	name: '首页',
	component: HomeView,
	meta: {
		key: 'system',
	},
};

export default route;
