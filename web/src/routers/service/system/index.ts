import type { RouteRecordRaw } from 'vue-router';
import SystemMenuView from '@/views/system/Menu/index.vue';
import SystemAuthView from '@/views/system/Auth/index.vue';
import SystemRoleView from '@/views/system/Role/index.vue';
import SystemDictView from '@/views/system/Dict/index.vue';
import RouterView from '@/components/RouterView/index.vue';

const route: RouteRecordRaw = {
	path: '/system',
	name: '系统设置',
	component: RouterView,
	meta: {
		key: 'system',
	},
	children: [
		{
			path: '/system/menu',
			name: '菜单设置',
			component: SystemMenuView,
			meta: {
				key: 'system-menu',
			},
		},
		{
			path: '/system/role',
			name: '角色管理',
			component: SystemRoleView,
			meta: {
				key: 'system-role',
			},
		},
		{
			path: '/system/dict',
			name: '字典管理',
			component: SystemDictView,
			meta: {
				key: 'system-dict',
			},
		},
		{
			path: '/system/auth',
			name: '用户设置',
			component: SystemAuthView,
			meta: {
				key: 'system-auth',
			},
		},
	],
};

export default route;
