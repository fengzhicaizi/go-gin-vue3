<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message, Select, SelectOption } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { createAuthApi, updateAuthApi } from '@/apis/system/auth';
import type { UpdateRoleBodyType } from '@/apis/system/role.d';
import type { AuthBodyType, UpdateAuthBodyType } from '@/apis/system/auth.d';
import type { Rule } from 'ant-design-vue/lib/form/index';

export type openFnArgsType = {
	type?: 'edit' | 'add';
	data?: AuthBodyType & { id: number };
	ok?: () => void;
	cancel?: () => void;
};

type FormState = {
	username: string;
	roleId: string;
	realName: string;
	phone: string;
};

const props = defineProps<{
	roles: UpdateRoleBodyType[];
}>();
const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('新增');
const modalProps: openFnArgsType = {};
const formRef = ref();
const col = {
	wrapperCol: { span: 15 },
	labelCol: { span: 6 },
};
const form = reactive<FormState>({
	username: '',
	roleId: '',
	realName: '',
	phone: '',
});
const rules: Record<string, Rule[]> = {
	username: [{ required: true, message: '账号必须填写' }],
	roleId: [{ required: true, message: '角色不能为空' }],
};

const cancel = () => {
	modalProps.cancel && modalProps.cancel();
	visible.value = false;
};

const ok = async () => {
	confirmLoading.value = true;
	try {
		const value = await formRef.value.validate();
		if (value) {
			const fn = modalProps.type === 'add' ? createAuthApi : modalProps.type === 'edit' ? updateAuthApi : null;
			const body: AuthBodyType | UpdateAuthBodyType = {
				id: modalProps.data?.id || null,
				roleId: Number(value.roleId) || 0,
				username: value.username || '',
				realName: value.realName || '',
				phone: value.phone || '',
			};
			const res = fn && (await fn(body));
			if (res && res.status === 200) {
				message.success(res.data.msg);
				nextTick(() => {
					visible.value = false;
					modalProps.ok && modalProps.ok();
				});
			}
		}
	} catch (err: any) {
		message.error(err.response.data.msg);
	} finally {
		confirmLoading.value = false;
	}
};

const open = (props: openFnArgsType) => {
	modalProps.type = props.type;
	modalProps.ok = props.ok;
	modalProps.cancel = props.cancel;
	modalProps.data = props.data;
	visible.value = true;
	nextTick(() => {
		if (modalProps.type === 'add') {
			formRef.value.resetFields && formRef.value.resetFields();
			title.value = '新增';
		}

		if (modalProps.type === 'edit') {
			formRef.value.resetFields && formRef.value.resetFields();
			form.username = props.data?.username || '';
			form.roleId = props.data?.roleId ? String(props.data?.roleId) : '';
			form.realName = props.data?.realName || '';
			form.phone = props.data?.phone || '';
			title.value = '编辑';
		}
	});
};

defineExpose({
	open,
});
</script>

<template>
	<Modal :visible="visible" :title="title" :confirm-loading="confirmLoading" @cancel="cancel" @ok="ok">
		<Form ref="formRef" v-bind="col" :model="form" :rules="rules" autocomplete="off">
			<FormItem label="账号" name="username">
				<Input v-model:value="form.username" />
			</FormItem>
			<FormItem label="角色" name="roleId">
				<Select v-model:value="form.roleId">
					<SelectOption v-for="list in props.roles" :value="`${list.id}`" :key="list.id">{{ list.name }}</SelectOption>
				</Select>
			</FormItem>
			<FormItem label="姓名" name="realName">
				<Input v-model:value="form.realName" />
			</FormItem>
			<FormItem label="电话号码" name="phone">
				<Input v-model:value="form.phone" />
			</FormItem>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
