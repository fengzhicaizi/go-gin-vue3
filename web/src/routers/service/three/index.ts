import type { RouteRecordRaw } from 'vue-router';
import RouterView from '@/components/RouterView/index.vue';
import ThreeEarthView from '@/views/three/earth/index.vue';
import ThreeHunanView from '@/views/three/hunan/index.vue';

const route: RouteRecordRaw = {
	path: '/three',
	name: 'Three',
	component: RouterView,
	meta: {
		key: 'three',
	},
	children: [
		{
			path: '/three/earth',
			name: '3D地球',
			component: ThreeEarthView,
			meta: {
				key: 'three-earth',
			},
		},
		{
			path: '/three/hunan3DMap',
			name: '湖南3D地图',
			component: ThreeHunanView,
			meta: {
				key: 'three-hunan3DMap',
			},
		},
	],
};

export default route;
