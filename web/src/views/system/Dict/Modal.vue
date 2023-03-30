<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message, Row, Col } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { createDictApi, updateDictApi } from '@/apis/system/dict';
import type { DictType } from '@/apis/system/dict.d';
import type { Rule } from 'ant-design-vue/lib/form/index';

export type openFnArgsType = {
	type?: 'edit' | 'add';
	id?: number;
	data?: DictType;
	ok?: () => void;
	cancel?: () => void;
};

type FormState = {
	name: string;
	type: string;
	remark: string;
};

const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('新增');
const modalProps: openFnArgsType = {};
const formRef = ref();
const col = {
	wrapperCol: { span: 18 },
	labelCol: { span: 5 },
};
const form = reactive<FormState>({
	name: '',
	type: '',
	remark: '',
});
const rules: Record<string, Rule[]> = {
	name: [{ required: true, message: '字典名称必须填写' }],
	type: [{ required: true, message: '字典类型必须填写' }],
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
			const body: DictType = {
				id: modalProps.id,
				name: value.name || '',
				type: value.type || '',
				remark: value.remark || '',
			};
			const res = await (modalProps.type === 'add' ? createDictApi(body) : modalProps.type === 'edit' ? updateDictApi(body) : null);
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
	modalProps.id = props.id;
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
			form.type = props.data?.type || '';
			form.remark = props.data?.remark || '';
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
					<FormItem label="字典名称" name="name">
						<Input v-model:value="form.name" />
					</FormItem>
					<FormItem label="字典类型" name="type">
						<Input v-model:value="form.type" :disabled="modalProps.type === 'edit'" />
					</FormItem>
					<FormItem label="备注" name="remark">
						<Input v-model:value="form.remark" />
					</FormItem>
				</Col>
			</Row>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
