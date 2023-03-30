<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message, Row, Col, Textarea } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { createMenuApi, updateMenuApi } from '@/apis/system/menu';
import type { AddMenuBodyType, UpdateMenuBodyType } from '@/apis/system/menu.d';
import type { Rule } from 'ant-design-vue/lib/form/index';

export type openFnArgsType = {
	type?: 'edit' | 'add';
	pid: number;
	data?: AddMenuBodyType & {
		id?: number;
		pid: number;
	};
	ok?: () => void;
	cancel?: () => void;
};

type FormState = {
	name: string;
	code: string;
	path: string;
	sort: number;
	icon: string;
	mark: string;
};

const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('新增');
const modalProps: openFnArgsType = {
	pid: 0,
};
const formRef = ref();
const col = {
	wrapperCol: { span: 20 },
	labelCol: { span: 4 },
};
const form = reactive<FormState>({
	name: '',
	code: '',
	path: '',
	sort: 1,
	icon: '',
	mark: '',
});
const rules: Record<string, Rule[]> = {
	name: [{ required: true, message: '菜单名必须填写' }],
	code: [{ required: true, message: 'code必须填写' }],
	path: [{ required: true, message: '路径必须填写' }],
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
			const fn = modalProps.type === 'add' ? createMenuApi : modalProps.type === 'edit' ? updateMenuApi : null;
			const body: AddMenuBodyType | UpdateMenuBodyType = {
				id: modalProps.data?.id || 0,
				pid: modalProps.pid || 1,
				name: value.name || '',
				code: value.code || '',
				path: value.path || '',
				sort: Math.floor(value.sort) || 1,
				icon: value.icon || '',
				mark: value.mark || '',
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
	modalProps.pid = props.pid;
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
			form.path = props.data?.path || '';
			form.code = props.data?.code || '';
			form.sort = props.data?.sort || 1;
			form.icon = props.data?.icon || '';
			form.mark = props.data?.mark || '';
			title.value = '编辑';
		}
	});
};

defineExpose({
	open,
});
</script>

<template>
	<Modal width="800px" :visible="visible" :title="title" :confirm-loading="confirmLoading" @cancel="cancel" @ok="ok">
		<Form ref="formRef" v-bind="col" :model="form" :rules="rules" autocomplete="off">
			<Row>
				<Col span="12">
					<FormItem label="菜单名" name="name">
						<Input v-model:value="form.name" placeholder="菜单名" />
					</FormItem>
				</Col>
				<Col span="12">
					<FormItem label="路径" name="path">
						<Input v-model:value="form.path" placeholder="例如: /system/menu" />
					</FormItem>
				</Col>
				<Col span="12">
					<FormItem label="code" name="code">
						<Input v-model:value="form.code" placeholder="用 '-' 连接各级路由,例如：system-menu" />
					</FormItem>
				</Col>
				<Col span="12">
					<FormItem label="排序" name="sort">
						<Input v-model:value="form.sort" />
					</FormItem>
				</Col>
				<Col span="12">
					<FormItem label="图标" name="icon">
						<Input v-model:value="form.icon" placeholder="icon-font图标,例如：icon-menu" />
					</FormItem>
				</Col>
				<Col span="24">
					<FormItem :label-col="{ span: 2 }" :wrapper-col="{ span: 22 }" label="备注" name="mark">
						<Textarea v-model:value="form.mark" />
					</FormItem>
				</Col>
			</Row>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
