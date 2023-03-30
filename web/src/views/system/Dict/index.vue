<script lang="ts" setup>
import { Table, Button, message, Popconfirm } from 'ant-design-vue';
import type { ColumnsType } from 'ant-design-vue/lib/table';
import { ref, onMounted, nextTick } from 'vue';
import { getDictlistApi, deleteDictApi } from '@/apis/system/dict';
import type { DictType } from '@/apis/system/dict.d';
import Modal from './Modal.vue';
import type { openFnArgsType } from './Modal.vue';
import AssignModal from './drawer/index.vue';
import type { openType } from './drawer/index.vue';
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
		title: '字典名称',
		dataIndex: 'name',
		align: 'center',
		width: 200,
	},
	{
		title: '字典类型',
		dataIndex: 'type',
		align: 'center',
		width: 150,
	},
	{
		title: '备注',
		dataIndex: 'remark',
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
					getDictList();
				});
			},
		};
		modal.value.open(props);
	}
};

const edit = (record: DictType & { id: number }) => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'edit',
			id: record.id,
			data: record,
			ok() {
				nextTick(() => {
					getDictList();
				});
			},
		};
		modal.value.open(props);
	}
};

const assgin = (
	row: DictType & {
		id: number;
	}
) => {
	if (assignModal.value) {
		const props: openType = {
			type: row.type,
			id: row.id,
			ok() {
				getDictList();
			},
		};
		assignModal.value.open(props);
	}
};

const del = async (id: number) => {
	try {
		const res = await deleteDictApi(id);
		if (res) {
			message.success(res.data.msg);
			nextTick(() => {
				getDictList();
			});
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	}
};

const getDictList = async () => {
	try {
		loading.value = true;

		const res = await getDictlistApi();

		if (res && res.data?.data && Array.isArray(res.data.data)) {
			data.value = res.data.data;
		}
	} catch (err) {
	} finally {
		loading.value = false;
	}
};

onMounted(() => {
	getDictList();
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
						<Button :disabled="record.username === 'admin'" size="small" @click="assgin(record)">配置</Button>
						<Button :disabled="record.username === 'admin'" size="small" @click="edit(record)">编辑</Button>
						<Popconfirm title="确定删除此行数据，删除后无法恢复?" @confirm="del(record.id)">
							<Button size="small" type="ghost" danger>删除</Button>
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
