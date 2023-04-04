import type { RouteRecordRaw } from 'vue-router';
import DefaultLayout from '@/components/pages/layouts/DefaultLayout/index.vue';
import systemRoute from './system';
import errorRoutes from './error';
import homeRoute from './home';
import cesiumRoute from './cesium';
import threeRoute from './three';

const addRoutes: RouteRecordRaw[] = [homeRoute, threeRoute, cesiumRoute, systemRoute].concat(errorRoutes);

const serviceRoutes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'main',
		component: DefaultLayout,
		redirect: '/home',
		children: addRoutes,
	},
];

export { addRoutes, serviceRoutes };
