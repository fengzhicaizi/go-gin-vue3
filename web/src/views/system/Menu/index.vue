<script lang="ts" setup>
import { Table, Button, message, Popconfirm } from 'ant-design-vue';
import type { ColumnsType } from 'ant-design-vue/lib/table';
import { ref, onMounted, nextTick } from 'vue';
import { getMenuApi, deleteMenuApi } from '@/apis/system/menu';
import type { UpdateMenuBodyType } from '@/apis/system/menu.d';
import Icon from '@/components/Icon';
import Modal from './Modal.vue';
import type { openFnArgsType } from './Modal.vue';

const data = ref([]);
const loading = ref(false);
const modal = ref();

const columns: ColumnsType = [
	{
		title: '菜单名',
		dataIndex: 'name',
		align: 'center',
		key: 'name',
		width: 200,
	},
	{
		title: '菜单地址',
		dataIndex: 'path',
		key: 'path',
		align: 'left',
		width: 200,
	},
	{
		title: '菜单编号',
		dataIndex: 'code',
		align: 'left',
		width: 200,
	},
	{
		title: '图标',
		align: 'center',
		key: 'icon',
		width: 100,
	},
	{
		title: '排序',
		dataIndex: 'sort',
		align: 'center',
		width: 100,
	},
	{
		title: '备注',
		dataIndex: 'mark',
		align: 'center',
	},
	{
		title: '操作',
		key: 'operation',
		dataIndex: 'operation',
		align: 'center',
		width: 200,
	},
];

const add = (pid: number) => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'add',
			pid: pid || 1,
			ok() {
				nextTick(() => {
					getMenu();
				});
			},
		};
		modal.value.open(props);
	}
};

const editMenu = (record: UpdateMenuBodyType) => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'edit',
			pid: record.pid,
			data: record,
			ok() {
				nextTick(() => {
					getMenu();
				});
			},
		};
		modal.value.open(props);
	}
};

const deleteMenu = async (id: string) => {
	try {
		const res = await deleteMenuApi(id);
		if (res) {
			message.success(res.data.msg);
			nextTick(() => {
				getMenu();
			});
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	}
};

const getMenu = async () => {
	try {
		loading.value = true;

		const res = await getMenuApi();

		if (res && res.data?.data && res.data.data.children && Array.isArray(res.data.data.children)) {
			data.value = res.data.data.children;
		}
	} catch (err) {
	} finally {
		loading.value = false;
	}
};

onMounted(() => {
	getMenu();
});
</script>

<template>
	<div class="main">
		<div class="header">
			<div class="search"></div>
			<div class="right">
				<Button type="primary" @click="add(1)">新增</Button>
			</div>
		</div>
		<Table size="small" :bordered="true" :data-source="data" :columns="columns" :row-key="record => record.id" :pagination="false" :loading="loading">
			<template #bodyCell="{ record, column }">
				<template v-if="column.key === 'name'">
					<div style="text-align: left; text-indent: 12px">
						{{ record.name }}
					</div>
				</template>
				<!-- <template v-if="column.key === 'path'">
					<div style="text-align: left; text-indent: 24px">{{ record.path }}</div>
				</template> -->
				<template v-if="column.key === 'icon'">
					<template v-if="record.icon">
						<Icon :type="record.icon"></Icon>
					</template>
				</template>
				<template v-if="column.key === 'operation'">
					<div style="width: 100%; display: flex; flex-direction: row; justify-content: space-around">
						<Button :disabled="record.username === 'admin'" size="small" @click="add(record.id)">新增</Button>
						<Button :disabled="record.username === 'admin'" size="small" @click="editMenu(record)">编辑</Button>
						<Popconfirm title="确定删除此行数据，删除后无法恢复?" @confirm="deleteMenu(record.id)">
							<Button size="small" type="ghost" danger>删除</Button>
						</Popconfirm>
					</div>
				</template>
			</template>
		</Table>
	</div>
	<Modal ref="modal"></Modal>
</template>

<style lang="less" scoped>
.main {
	margin: 12px;
	padding: 24px;
	border-radius: 6px;
	background: white;
	.header {
		height: 50px;
		display: flex;
		flex-direction: row;
		align-items: flex-start;
		justify-content: space-between;
	}
}
</style>
