<script lang="ts" setup>
import { Table, Popconfirm, Button, message } from 'ant-design-vue';
import type { ColumnsType } from 'ant-design-vue/lib/table';
import { nextTick, ref } from 'vue';
import { deleteDictDataApi, getDictDataApi } from '@/apis/system/dict';
import type { DictData } from '@/apis/system/dict.d';
import Modal from './Modal.vue';

const data = ref([]);
const loading = ref(false);
const modal = ref();
let currentType = '';
const columns: ColumnsType = [
	{
		title: '序号',
		width: 100,
		align: 'center',
		customRender: ({ index }) => `${index + 1}`,
	},
	{
		title: '字典名',
		dataIndex: 'label',
		align: 'center',
	},
	{
		title: '字典值',
		dataIndex: 'value',
		align: 'center',
	},
	{
		title: '排序',
		dataIndex: 'sort',
		align: 'center',
	},
	{
		title: '操作',
		key: 'operation',
		dataIndex: 'operation',
		align: 'center',
		width: 150,
	},
];

const query = async (type: string) => {
	currentType = type;
	loading.value = true;
	try {
		const res = await getDictDataApi(type);

		if (res && res.data.data && Array.isArray(res.data.data)) {
			data.value = res.data.data;
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	} finally {
		loading.value = false;
	}
};

const add = (type: string) => {
	modal.value.open({
		type: 'add',
		dictType: type,
		ok: () => {
			query(currentType);
		},
	});
};

const edit = (
	record: DictData & {
		id: number;
	}
) => {
	//
	modal.value.open({
		type: 'edit',
		id: record.id,
		data: record,
		ok: () => {
			query(currentType);
		},
	});
};

const del = async (id: number) => {
	try {
		const res = await deleteDictDataApi(id);

		if (res) {
			nextTick(() => {
				query(currentType);
			});
		}
	} catch (err: any) {
		message.error(err.response.data.detail);
	}
};

defineExpose({
	add,
	query,
});
</script>

<template>
	<div>
		<Table size="small" :data-source="data" :columns="columns" :loading="loading" :pagination="false">
			<template #bodyCell="{ record, column }">
				<template v-if="column.key === 'operation'">
					<div style="width: 100%; display: flex; flex-direction: row; justify-content: space-around">
						<Button :disabled="record.username === 'admin'" size="small" @click="edit(record)">编辑</Button>
						<Popconfirm title="确定删除此行数据，删除后无法恢复?" @confirm="del(record.id)">
							<Button size="small" type="ghost" danger>删除</Button>
						</Popconfirm>
					</div>
				</template>
			</template>
		</Table>
		<Modal ref="modal"></Modal>
	</div>
</template>

<style lang="less"></style>
