<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message, Row, Col } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { createRoleApi, updateRoleApi } from '@/apis/system/role';
import type { AddRoleBodyType, UpdateRoleBodyType } from '@/apis/system/role.d';
import type { Rule } from 'ant-design-vue/lib/form/index';

export type openFnArgsType = {
	type?: 'edit' | 'add';
	data?: UpdateRoleBodyType;
	ok?: () => void;
	cancel?: () => void;
};

type FormState = {
	name: string;
};

const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('新增');
const modalProps: openFnArgsType = {};
const formRef = ref();
const col = {
	wrapperCol: { span: 20 },
	labelCol: { span: 4 },
};
const form = reactive<FormState>({
	name: '',
});
const rules: Record<string, Rule[]> = {
	name: [{ required: true, message: '角色名必须填写' }],
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
			const fn = modalProps.type === 'add' ? createRoleApi : modalProps.type === 'edit' ? updateRoleApi : null;
			const body: AddRoleBodyType | UpdateRoleBodyType = {
				id: modalProps.data?.id || 0,
				name: value.name || '',
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
			form.name = props.data?.name || '';
			title.value = '编辑';
		}
	});
};

defineExpose({
	open,
});
</script>

<template>
	<Modal width="500px" :visible="visible" :title="title" :confirm-loading="confirmLoading" @cancel="cancel" @ok="ok">
		<Form ref="formRef" v-bind="col" :model="form" :rules="rules" autocomplete="off">
			<Row>
				<Col span="24">
					<FormItem label="角色名" name="name">
						<Input v-model:value="form.name" />
					</FormItem>
				</Col>
			</Row>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
