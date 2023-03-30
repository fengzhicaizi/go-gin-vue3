<script lang="ts" setup>
import { ref, onMounted, reactive, nextTick } from 'vue';
import { Table, Button, Input, InputGroup, Select, SelectOption, Popconfirm, message } from 'ant-design-vue';
import Modal from './Modal.vue';
import { getAuthListApi, deleteAuthApi, resetPasswordApi } from '@/apis/system/auth';
import { getRoleListApi } from '@/apis/system/role';
import type { AuthBodyType } from '@/apis/system/auth.d';
import type { QueryParamsType } from '@/apis/system/auth.d';
import type { openFnArgsType } from './Modal.vue';
import type { ColumnsType, TablePaginationConfig } from 'ant-design-vue/lib/table';
import type { UpdateRoleBodyType } from '@/apis/system/role.d';
import Icon from '@/components/Icon';
import dayjs from 'dayjs';

const modal = ref();
const data = ref([]);
const roles = ref<UpdateRoleBodyType[]>([]);
const columns: ColumnsType = [
	{
		title: '序号',
		width: '100px',
		align: 'center',
		customRender: ({ index }) => `${index + 1}`,
	},
	{
		title: '姓名',
		dataIndex: 'realName',
		key: 'realName',
		align: 'center',
	},
	{
		title: '账号',
		dataIndex: 'username',
		key: 'username',
		align: 'center',
	},
	{
		title: '电话号码',
		dataIndex: 'phone',
		key: 'phone',
		align: 'center',
	},
	{
		title: '角色',
		dataIndex: 'roleId',
		key: 'roleId',
		align: 'center',
	},
	{
		title: '创建时间',
		dataIndex: 'createdAt',
		align: 'center',
		customRender: ({ record }) => dayjs(record.createdAt).format('YYYY-MM-DD HH:mm:ss'),
	},
	{
		title: '操作',
		key: 'operation',
		dataIndex: 'operation',
		align: 'center',
		width: 250,
	},
];
const loading = ref(false);
const pagination = reactive({
	total: 0,
	current: 1,
	pageSize: 10,
});
const search = reactive<{
	value: string;
	condition: 'username' | 'realName' | 'phone';
}>({
	value: '',
	condition: 'username',
});

const add = () => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'add',
			ok() {
				nextTick(() => {
					getAuthList();
				});
			},
		};
		modal.value.open(props);
	}
};

const getAuthList = async () => {
	try {
		const param: QueryParamsType = {
			pageSize: pagination.pageSize,
			page: pagination.current,
		};

		if (search.value !== '') {
			param[search.condition] = search.value;
		}

		loading.value = true;

		const res = await getAuthListApi(param);
		const value = res.data.data;

		if (value && value.list && value.page && Array.isArray(value.list)) {
			data.value = value.list;
			pagination.total = value.page.total;
			pagination.current = value.page.page;
			pagination.pageSize = value.page.pageSize;
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	} finally {
		loading.value = false;
	}
};

const getRoleList = async () => {
	try {
		const res = await getRoleListApi();

		if (res) {
			roles.value = res.data.data;
		}
	} catch (err: any) {
		message.error(err);
	}
};

const editAuth = (record: AuthBodyType & { id: number }) => {
	if (modal.value) {
		const props: openFnArgsType = {
			type: 'edit',
			data: record,
			ok() {
				nextTick(() => {
					getAuthList();
				});
			},
		};
		modal.value.open(props);
	}
};

const deleteAuth = async (id: number) => {
	try {
		const res = await deleteAuthApi(id);
		if (res) {
			message.success(res.data.msg);
			nextTick(() => {
				getAuthList();
			});
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	}
};

const resetPassword = async (id: number) => {
	try {
		const res = await resetPasswordApi(id);
		if (res) {
			message.success(res.data.msg);
			nextTick(() => {
				getAuthList();
			});
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	}
};

const change = (p: TablePaginationConfig) => {
	pagination.current = p.current || 1;
	getAuthList();
};

const getRoleName = (id: number) => {
	return roles.value.filter(f => f.id == id)[0]?.name;
};

onMounted(() => {
	getAuthList();
	getRoleList();
});
</script>

<template>
	<div class="main">
		<div class="header">
			<div class="search">
				<InputGroup compact>
					<Select :dropdownStyle="{ textAlign: 'center' }" style="width: 100px; text-align: center" v-model:value="search.condition">
						<SelectOption value="username">账 号</SelectOption>
						<SelectOption value="realName">姓 名</SelectOption>
						<SelectOption value="phone">电话号码</SelectOption>
					</Select>
					<Input v-model:value="search.value" style="width: 200px" allowClear></Input>
				</InputGroup>
				<Button style="margin-left: 24px" @click="getAuthList" type="primary">
					<template #icon><Icon type="icon-sousuo" style="font-size: 14px"></Icon></template>
					查询
				</Button>
			</div>
			<div class="right">
				<Button @click="add" type="primary">新增</Button>
			</div>
		</div>
		<Table size="small" :bordered="true" :data-source="data" :pagination="pagination" :columns="columns" row-key="id" :loading="loading" @change="change">
			<template #bodyCell="{ record, column }">
				<template v-if="column.key === 'operation'">
					<div style="width: 100%; display: flex; flex-direction: row; justify-content: space-around">
						<Button :disabled="record.isRoot" size="small" @click="editAuth(record)">编辑</Button>
						<Popconfirm :disabled="record.isRoot" title="确定删除此行数据，删除后无法恢复?" @confirm="deleteAuth(record.id)">
							<Button :disabled="record.isRoot" size="small" type="ghost" danger>删除</Button>
						</Popconfirm>
						<Popconfirm @confirm="resetPassword(record.id)">
							<template #title>
								<div style="width: 250px; text-align: center">
									<div>确认重置改用户的密码?</div>
									<div>"重置后的密码为123456,请及时更改.</div>
								</div>
							</template>
							<Button size="small" type="dashed" danger>重置密码</Button>
						</Popconfirm>
					</div>
				</template>
				<template v-if="column.key === 'roleId'">
					<div>{{ getRoleName(record.roleId) }}</div>
				</template>
			</template>
		</Table>
		<Modal ref="modal" :roles="roles"></Modal>
	</div>
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
		.search {
			display: flex;
			flex-direction: row;
		}
	}
}
</style>
