<script lang="ts" setup>
import { Modal, Form, FormItem, Input, message, Row, Col } from 'ant-design-vue';
import { ref, reactive, nextTick } from 'vue';
import { addDictDataApi, updateDictDataApi } from '@/apis/system/dict';
import type { DictData } from '@/apis/system/dict.d';
import type { Rule } from 'ant-design-vue/lib/form/index';

export type openFnArgsType = {
	type?: 'edit' | 'add';
	id: number;
	dictType: string;
	data?: DictData;
	ok?: () => void;
	cancel?: () => void;
};

type FormState = {
	label: string;
	value: string;
	sort: number;
};

const visible = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const title = ref<string>('新增');
const modalProps: openFnArgsType = {
	id: 0,
	dictType: '',
};
const formRef = ref();
const col = {
	wrapperCol: { span: 18 },
	labelCol: { span: 5 },
};
const form = reactive<FormState>({
	label: '',
	value: '',
	sort: 1,
});
const rules: Record<string, Rule[]> = {
	label: [{ required: true, message: '字典名称必须填写' }],
	value: [{ required: true, message: '字典类型必须填写' }],
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
			const body: DictData = {
				id: modalProps.id,
				label: value.label || '',
				value: value.value || '',
				sort: parseInt(value.sort) || 1,
			};
			const res = await (modalProps.type === 'add' ? addDictDataApi(modalProps.dictType, body) : modalProps.type === 'edit' ? updateDictDataApi(body) : null);
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
	modalProps.dictType = props.dictType;
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
			form.label = props.data?.label || '';
			form.value = props.data?.value || '';
			form.sort = props.data?.sort || 1;
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
					<FormItem label="字典名" name="label">
						<Input v-model:value="form.label" />
					</FormItem>
					<FormItem label="字典值" name="value">
						<Input v-model:value="form.value" />
					</FormItem>
					<FormItem label="排序" name="sort">
						<Input v-model:value="form.sort" />
					</FormItem>
				</Col>
			</Row>
		</Form>
	</Modal>
</template>

<style lang="less"></style>
