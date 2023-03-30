<script lang="ts" setup>
import { Table, Button, message, Popconfirm } from 'ant-design-vue';
import type { ColumnsType } from 'ant-design-vue/lib/table';
import { ref, onMounted, nextTick } from 'vue';
import { getRoleListApi, deleteRoleApi } from '@/apis/system/role';
import type { UpdateRoleBodyType } from '@/apis/system/role.d';
import Modal from './Modal.vue';
import type { openFnArgsType } from './Modal.vue';
import AssignModal from './AssignModal.vue';
import type { openType } from './AssignModal.vue';
import dayjs from 'dayjs';

const data = ref([]);

const columns: ColumnsType = [
	{
		title: '序号',
		width: 100,
		align: 'center',
		customRender: ({ index }) => `${index + 1}`,
	},
	{
		title: '角色名',
		dataIndex: 'name',
		align: 'center',
	},
	{
		title: '创建时间',
		dataIndex: 'createdAt',
		align: 'center',
		width: 200,
		customRender: ({ record }) => dayjs(record.createdAt).format('YYYY-MM-DD HH:mm:ss'),
	},
	{
		title: '操作',
		key: 'operation',
		dataIndex: 'operation',
		align: 'center',
		width: 200,
	},
];

const loading = ref(false);
const modal = ref();
const assignModal = ref();

const add = () => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'add',
			ok() {
				nextTick(() => {
					getRoleList();
				});
			},
		};
		modal.value.open(props);
	}
};

const edit = (record: UpdateRoleBodyType) => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'edit',
			data: record,
			ok() {
				nextTick(() => {
					getRoleList();
				});
			},
		};
		modal.value.open(props);
	}
};

const assgin = (row: UpdateRoleBodyType) => {
	if (assignModal.value) {
		const props: openType = {
			id: row.id,
			menuIds: row.menuIds || '',
			ok() {
				getRoleList();
			},
		};
		assignModal.value.open(props);
	}
};

const deleteMenu = async (id: string) => {
	try {
		const res = await deleteRoleApi(id);
		if (res) {
			message.success(res.data.msg);
			nextTick(() => {
				getRoleList();
			});
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	}
};

const getRoleList = async () => {
	try {
		loading.value = true;

		const res = await getRoleListApi();

		if (res && res.data?.data && Array.isArray(res.data.data)) {
			data.value = res.data.data;
		}
	} catch (err: any) {
		message.error(err);
	} finally {
		loading.value = false;
	}
};

onMounted(() => {
	getRoleList();
});
</script>

<template>
	<div class="main">
		<div class="header">
			<div class="search"></div>
			<div class="right">
				<Button type="primary" @click="add()">新增</Button>
			</div>
		</div>
		<Table size="small" :bordered="true" :data-source="data" :columns="columns" :row-key="record => record.id" :pagination="false" :loading="loading">
			<template #bodyCell="{ record, column }">
				<template v-if="column.key === 'operation'">
					<div style="width: 100%; display: flex; flex-direction: row; justify-content: space-around">
						<Button size="small" @click="assgin(record)" :disabled="record.isAllMenus">配置</Button>
						<Button size="small" @click="edit(record)" :disabled="record.isAllMenus">编辑</Button>
						<Popconfirm title="确定删除此行数据，删除后无法恢复?" @confirm="deleteMenu(record.id)">
							<Button size="small" type="ghost" danger :disabled="record.isAllMenus">删除</Button>
						</Popconfirm>
					</div>
				</template>
			</template>
		</Table>
	</div>
	<Modal ref="modal"></Modal>
	<AssignModal ref="assignModal"></AssignModal>
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
