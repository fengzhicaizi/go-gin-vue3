<script lang="ts" setup>
import { nextTick, reactive, ref } from 'vue';
import { Drawer, Button } from 'ant-design-vue';
import Table from './Table.vue';

export type openType = {
	type: string;
	id: number;
	ok?: () => void;
};

const drawer = reactive({
	visible: false,
	title: '',
});
const table = ref();
const modalProps: {
	type?: string;
	id?: number;
	ok?: () => void;
} = {};

const open = async (props: openType) => {
	modalProps.type = props.type;
	modalProps.id = props.id;
	drawer.title = props.type;
	drawer.visible = true;
	nextTick(() => {
		//
		table.value.query(props.type);
	});
};

const close = () => {
	drawer.visible = false;
};

const add = () => {
	table.value.add(modalProps.type);
};

defineExpose({
	open,
});
</script>

<template>
	<Drawer class="drawer" width="700px" :visible="drawer.visible" @close="close">
		<template #title>
			<div class="header">
				<div>{{ drawer.title }}</div>
				<Button type="primary" @click="add">新增</Button>
			</div>
		</template>
		<Table ref="table" class="table"></Table>
	</Drawer>
</template>

<style lang="less" scoped>
.drawer {
	.header {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
	}
	.table {
		margin: -24px;
	}
}
</style>
