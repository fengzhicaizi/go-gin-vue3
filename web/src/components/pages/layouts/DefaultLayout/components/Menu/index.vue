<script lang="ts" setup>
import { ref } from 'vue';
import { Menu, MenuItem, SubMenu } from 'ant-design-vue';
import type { MenuClickEventHandler } from 'ant-design-vue/lib/menu/src/interface';
import Icon from '@/components/Icon/index';
import { useAuthStore } from '@/stores/system/auth';
import { storeToRefs } from 'pinia';
import router from '@/routers';

const emit = defineEmits<{
	(event: 'change', path: string): void;
}>();

const localSelectedKeys = localStorage.getItem('selectedKeys');
const localOpenKeys = localStorage.getItem('openKeys');

const authStore = useAuthStore();
const { menu } = storeToRefs(authStore);

const selectedKeys = ref<string[]>([localSelectedKeys || '']);
const openKeys = ref<string[]>([localOpenKeys || '']);

const changeMenu: MenuClickEventHandler = ({ item }) => {
	emit('change', item.path);
};

// router监控
router.beforeEach((to, from, next) => {
	localStorage.setItem('selectedKeys', to?.meta?.key as string);
	if (to.matched && to.matched[1]) {
		localStorage.setItem('openKeys', to.matched[1]?.meta?.key as string);
	}
	selectedKeys.value = [`${to.meta.key}`];
	next();
});
</script>

<template>
	<Menu v-model:openKeys="openKeys" v-model:selectedKeys="selectedKeys" theme="dark" mode="inline" @click="changeMenu">
		<template v-if="menu && menu.children && menu.children.length">
			<template :key="m.code" v-for="m in menu.children">
				<template v-if="m.children && m.children.length !== 0">
					<SubMenu :key="m.code" :title="m.name">
						<template #icon>
							<Icon :type="m.icon || 'icon-default'" />
						</template>
						<div>
							<template :key="sm.code || ''" v-for="sm in m.children">
								<MenuItem :path="sm.path">{{ sm.name }}</MenuItem>
							</template>
						</div>
					</SubMenu>
				</template>
				<template v-if="!m.children">
					<MenuItem :key="m.code" :path="m.path">
						<template #icon>
							<Icon :type="m.icon || 'icon-default'" />
						</template>
						{{ m.name }}
					</MenuItem>
				</template>
			</template>
		</template>
	</Menu>
</template>
