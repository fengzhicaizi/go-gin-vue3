import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue()],
	resolve: {
		alias: {
			// eslint-disable-next-line @typescript-eslint/ban-ts-comment
			// @ts-ignore
			'@': process.cwd() + '/src',
		},
	},
	css: {
		preprocessorOptions: {
			less: {
				javascriptEnabled: true,
			},
		},
	},
	server: {
		port: 9000,
		hmr: false, // 关闭热更新
	},
});
