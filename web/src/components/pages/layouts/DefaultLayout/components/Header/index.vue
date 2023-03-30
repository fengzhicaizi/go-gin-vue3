<script lang="ts" setup>
import { Breadcrumb, BreadcrumbItem, Dropdown, Menu, MenuItem, Modal, message } from 'ant-design-vue';
import { nextTick, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '@/stores/system/auth';
import Icon from '@/components/Icon';
import AuthModal from './Modal.vue';
import type { openFnArgsType } from './Modal.d';
import router from '@/routers';
import config from '@/conf/index';

const route = useRoute();
const authStore = useAuthStore();
const modalRef = ref();

const loginOut = () => {
	Modal.confirm({
		title: '确认退出？',
		onOk() {
			localStorage.clear();
			router.replace({
				path: '/login',
			});
			message.success('退出成功,请重新登录');
		},
	});
};

const openUserInfoModal = () => {
	if (modalRef.value) {
		const props: openFnArgsType = {
			type: 'edit',
			ok() {
				nextTick(() => {
					localStorage.clear();
					router.replace({
						path: '/login',
					});
					message.success('个人信息修改成功,请重新登录');
				});
			},
		};
		modalRef.value.open(props);
	}
};

// 全屏幕
const changeFullScreen = () => {
	const id = config.fullScreenId;
	const dom = document.getElementById(id);

	if (dom) {
		dom.requestFullscreen();
	} else {
		message.error('没有需要全屏的内容');
	}
};
</script>

<template>
	<div class="header">
		<Breadcrumb class="crumb">
			<template v-for="item in route.matched.filter(f => f.name != '')">
				<BreadcrumbItem v-if="item.path !== '/'" :key="item.path">{{ item.name }}</BreadcrumbItem>
			</template>
		</Breadcrumb>
		<div class="ability">
			<Icon title="全屏" type="icon-quanping" @click="changeFullScreen" class="icon"></Icon>
		</div>
		<div class="account">
			<Dropdown trigger="click">
				<div class="dropdown">
					<div class="portrait">
						<Icon type="icon-account" style="font-size: 20px"></Icon>
					</div>
					<div class="name">{{ authStore.authGetter?.realName || '--' }}</div>
				</div>

				<template #overlay>
					<Menu class="menu">
						<MenuItem style="text-align: center" @click="openUserInfoModal"><Icon type="icon-user" class="icon"></Icon>个人设置</MenuItem>
						<MenuItem style="text-align: center" @click="loginOut"><Icon type="icon-tuichu" class="icon"></Icon>退出登录</MenuItem>
					</Menu>
				</template>
			</Dropdown>
		</div>
	</div>
	<AuthModal ref="modalRef"></AuthModal>
</template>

<style lang="less" scoped>
.header {
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: space-between;
	align-items: center;
	margin-left: 24px;
	.crumb {
		flex: 1;
	}
	.ability {
		display: flex;
		flex-direction: row;
		align-items: center;
		.icon {
			font-size: 20px;
			margin-right: 24px;
			transition: all 0.2s;
			&:hover {
				color: #000;
				cursor: pointer;
			}
		}
	}
	.account {
		display: flex;
		flex-direction: row;
		align-items: center;
		cursor: pointer;
	}
}
.dropdown {
	display: flex;
	flex-direction: row;
	align-items: center;
	.portrait {
		height: 30px;
		width: 30px;
		margin-right: 12px;
		border-radius: 50%;
		overflow: hidden;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		border: 1px solid #f1d5d6;
		background-color: #f1d5d6;
	}
}
.menu {
	width: 120px;
	padding: 12px 0;
	.icon {
		font-size: 14px;
		margin-right: 12px;
		color: #333;
	}
}
</style>
