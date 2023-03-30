import axios from 'axios';
import config from '@/conf';
import { message } from 'ant-design-vue';
import router from '@/routers';

const request = axios.create({
	baseURL: config.baseUrl,
	timeout: 5000,
});

request.interceptors.request.use(
	config => {
		if (config.headers) {
			config.headers.Authorization = localStorage.getItem('token') || '';
		}
		return config;
	},
	err => {
		Promise.reject(err);
	}
);

request.interceptors.response.use(
	response => {
		const { data } = response;
		if (data && data.code === 401) {
			message.error('账号验证失败，请重新登录');
			router.push('/login');
			return;
		}
		return response;
	},
	err => {
		console.log(err);
	}
);

export default request;
