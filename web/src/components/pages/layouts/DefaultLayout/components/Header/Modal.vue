<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { updateAuthApi } from '@/apis/system/auth';
import type { UpdateAuthBodyType } from '@/apis/system/auth.d';
import type { Rule } from 'ant-design-vue/lib/form/index';
import { useAuthStore } from '@/stores/system/auth';
import crypto from 'crypto-js';
import type { openFnArgsType } from './Modal.d';

const authStore = useAuthStore();

type FormState = {
	username: string;
	realName: string;
	phone: string;
	password: string;
};

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
	realName: '',
	phone: '',
	password: '',
});
const rules: Record<string, Rule[]> = {
	username: [{ required: true, message: '账号必须填写' }],
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
			const fn = updateAuthApi;
			const body: UpdateAuthBodyType = {
				id: modalProps.data?.id || null,
				username: value.username || '',
				realName: value.realName || '',
				phone: value.phone || '',
				roleId: 0,
			};
			if (value.password !== '') {
				body.password = crypto.MD5(encodeURI(value.password)).toString();
			}
			const res = fn && (await fn(body));
			if (res && res.status === 200) {
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
	modalProps.data = authStore.authGetter;
	visible.value = true;
	nextTick(() => {
		if (modalProps.type === 'edit') {
			formRef.value.resetFields && formRef.value.resetFields();
			form.username = modalProps.data?.username || '';
			form.realName = modalProps.data?.realName || '';
			form.phone = modalProps.data?.phone || '';
			title.value = '个人信息';
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
			<FormItem label="姓名" name="realName">
				<Input v-model:value="form.realName" />
			</FormItem>
			<FormItem label="电话号码" name="phone">
				<Input v-model:value="form.phone" />
			</FormItem>
			<FormItem label="新密码" name="password">
				<Input type="password" v-model:value="form.password" />
			</FormItem>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
