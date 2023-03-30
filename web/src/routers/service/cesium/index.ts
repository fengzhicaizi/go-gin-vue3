import type { RouteRecordRaw } from 'vue-router';
import RouterView from '@/components/RouterView/index.vue';
import CesiumDemo1View from '@/views/cesium/demo1/index.vue';

const route: RouteRecordRaw = {
	path: '/cesium',
	name: 'cesium',
	component: RouterView,
	meta: {
		key: 'cesium',
	},
	children: [
		{
			path: '/cesium/demo1',
			name: 'cesium-demo1',
			component: CesiumDemo1View,
			meta: {
				key: 'cesium-demo1',
			},
		},
	],
};

export default route;
