<script setup lang="ts">
import { RouterView } from 'vue-router';
import { ConfigProvider } from 'ant-design-vue';
import zhCN from 'ant-design-vue/es/locale/zh_CN';
</script>

<template>
	<ConfigProvider :locale="zhCN">
		<RouterView />
	</ConfigProvider>
</template>

<style lang="less">
@import url('ant-design-vue/dist/antd.less');
@import url('./assets/styles/antd.global.less');
@import url('./assets/styles/global.less');
</style>
