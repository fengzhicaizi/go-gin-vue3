<script lang="ts" setup>
import { Modal, Tree, message } from 'ant-design-vue';
import { ref, nextTick } from 'vue';
import { updateRoleApi } from '@/apis/system/role';
import { getMenuApi } from '@/apis/system/menu';
import type { TreeProps } from 'ant-design-vue';

export type openType = {
	id: number;
	menuIds: string;
	ok?: () => void;
};

const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('配置菜单');
const modalProps: openType = {
	id: 0,
	menuIds: '',
};
const treeData = ref<TreeProps['treeData']>([]);
const checkedKeys = ref<number[]>([]);
const fieldNames: TreeProps['fieldNames'] = {
	title: 'name',
	key: 'id',
};

const cancel = () => {
	visible.value = false;
};

const ok = async () => {
	confirmLoading.value = true;
	try {
		const res = await updateRoleApi({
			id: modalProps.id,
			menuIds: checkedKeys.value.join(','),
		});

		if (res) {
			message.success(res.data.msg);
			visible.value = false;
			modalProps.ok && modalProps.ok();
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	} finally {
		confirmLoading.value = false;
	}
};

const open = async (props: openType) => {
	treeData.value = [];
	modalProps.id = props.id;
	modalProps.ok = props?.ok;
	visible.value = true;
	await queryAllMenusByTree();
	nextTick(() => {
		checkedKeys.value = props.menuIds.split(',').map(f => parseInt(f));
	});
};

const queryAllMenusByTree = async () => {
	confirmLoading.value = true;
	try {
		const res = await getMenuApi();

		if (res && res.data?.data && res.data.data.children) {
			treeData.value = [res.data.data];
		}
	} catch (err: any) {
		message.error(err);
	} finally {
		confirmLoading.value = false;
	}
};

defineExpose({
	open,
});
</script>

<template>
	<Modal width="500px" :visible="visible" :title="title" :confirm-loading="confirmLoading" @cancel="cancel" @ok="ok">
		<div class="modal-body">
			<Tree v-if="treeData?.length" :height="400" :field-names="fieldNames" default-expand-all v-model:checkedKeys="checkedKeys" checkable :tree-data="treeData"> </Tree>
		</div>
	</Modal>
</template>

<style lang="less">
.modal-body {
	height: 400px;
	width: 400px;
	margin: 0 auto;
}
</style>
