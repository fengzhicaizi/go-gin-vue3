<script lang="ts" setup>
import { Layout, LayoutSider, LayoutHeader, LayoutContent, Spin } from 'ant-design-vue';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons-vue';
import { RouterView } from 'vue-router';
import { onBeforeMount, ref } from 'vue';
import Menu from './components/Menu/index.vue';
import Header from './components/Header/index.vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/system/auth';

const collapsed = ref(false);
const router = useRouter();
const authStore = useAuthStore();

const changeMenu = (path: string) => {
	router.push(path);
};

onBeforeMount(() => {
	authStore.getAuth();
});
</script>

<template>
	<div class="loading" v-if="!authStore.auth?.id">
		<Spin size="large"></Spin>
	</div>
	<Layout v-else class="defaut-layout">
		<LayoutSider class="sider" v-model:collapsed="collapsed" :trigger="null" collapsible>
			<div class="logo"></div>
			<Menu @change="changeMenu" />
		</LayoutSider>
		<Layout>
			<LayoutHeader class="header">
				<div class="outlined">
					<MenuUnfoldOutlined v-if="collapsed" class="trigger" @click="() => (collapsed = !collapsed)" />
					<MenuFoldOutlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
				</div>
				<div class="other">
					<Header></Header>
				</div>
			</LayoutHeader>
			<LayoutContent class="content">
				<div class="layout">
					<RouterView></RouterView>
				</div>
			</LayoutContent>
		</Layout>
	</Layout>
</template>

<style lang="less" scoped>
.loading {
	width: 100vw;
	height: 100vh;
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: center;
}
.defaut-layout {
	// position: relative;
	width: 100vw;
	height: 100vh;
	.sider {
		.logo {
			height: 32px;
			margin: 16px;
			background: rgba(255, 255, 255, 0.3);
		}
	}
	.header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		.outlined {
			display: flex;
			flex-direction: row;
			align-items: center;
		}
		.other {
			flex: 1;
		}
		.trigger {
			font-size: 20px;
		}
	}
	.content {
		position: relative;
		.layout {
			position: absolute;
			left: 0;
			right: 0;
			top: 0;
			bottom: 0;
			overflow: auto;
		}
	}
}
:deep(.ant-layout-header) {
	background-color: white;
	padding: 0 24px;
}
</style>
