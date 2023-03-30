<script lang="ts" setup>
import { reactive, ref } from 'vue';
import type { UnwrapRef } from 'vue';
import { Form, FormItem, Button, Input, message } from 'ant-design-vue';
import Icon from '@/components/Icon';
import crypto from 'crypto-js';
import { loginApi } from '@/apis/system/auth';
import type { LoginBodyType } from '@/apis/system/auth.d';
import type { Rule } from 'ant-design-vue/lib/form/index';
import type { FormInstance } from 'ant-design-vue';
import router from '@/routers';

interface FormState {
	username: string;
	password: string;
}

const labelCol = { style: { width: '0px' } };
const wrapperCol = { span: 24 };

const form: UnwrapRef<FormState> = reactive({
	username: '',
	password: '',
});
const formRef = ref<FormInstance>();

const rules: Record<string, Rule[]> = {
	username: [{ required: true, message: '账号不能为空', trigger: 'change' }],
	password: [{ required: true, message: '密码不能为空' }],
};

const login = async () => {
	const values = await formRef.value?.validate();

	if (values) {
		const body: LoginBodyType = {
			username: values.username,
			password: crypto.MD5(encodeURI(values.password)).toString(),
		};

		try {
			const res = await loginApi(body);
			const { data } = res;

			if (data.code === 200) {
				const token = data.data;
				localStorage.setItem('token', `Bearer ${token}`);
				message.success('登录成功');
				router.push({
					path: '/home',
				});
			}
		} catch (err: any) {
			message.error(err.response.data.detail);
		}
	}
};
</script>

<template>
	<div class="login">
		<div class="login-box">
			<div class="title">登录</div>
			<Form class="form" ref="formRef" :model="form" :rules="rules" :label-col="labelCol" :wrapper-col="wrapperCol" @finish="login">
				<FormItem has-feedback name="username">
					<Input v-model:value="form.username" size="large" placeholder="请输入账号" autocomplete="off">
						<template #prefix>
							<Icon type="icon-user" />
						</template>
					</Input>
				</FormItem>
				<FormItem has-feedback name="password">
					<Input v-model:value="form.password" type="password" size="large" placeholder="请输入密码" autocomplete="off">
						<template #prefix>
							<Icon type="icon-password" />
						</template>
					</Input>
				</FormItem>

				<!-- <div style="margin-top: 24px" class="verification">
					<Input style="flex: 1" type="password" size="large" placeholder="请输入验证码">
						<template #prefix>
							<Icon type="icon-yanzhengma" />
						</template>
					</Input>
					<div class="code"></div>
				</div> -->
				<Button style="margin-top: 12px" class="btn" size="large" html-type="submit" type="primary">登录</Button>
			</Form>
		</div>
	</div>
</template>

<style lang="less" scoped>
.login {
	width: 100vw;
	height: 100vh;
	background-image: url('/th.jpg');
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: flex-end;
	.login-box {
		width: 350px;
		background: white;
		margin-right: 48px;
		border-radius: 15px;
		display: flex;
		flex-direction: column;
		align-items: center;
		box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
		padding: 24px 24px 48px;
		.form {
			width: 100%;
		}
		.title {
			font-size: 20px;
			margin: 30px auto 30px;
		}
		.btn {
			width: 100%;
		}
		.verification {
			display: flex;
			flex-direction: row;

			.code {
				width: 100px;
				border: 1px solid red;
				margin-left: 12px;
			}
		}
	}
	:deep(.ant-input-lg::placeholder) {
		font-size: 14px;
	}
	:deep(.ant-input-prefix) {
		margin-right: 12px;
	}
}
</style>
