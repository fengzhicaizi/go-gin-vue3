<script lang="ts" setup>
import { Result } from 'ant-design-vue';
</script>

<template>
	<Result class="full" status="404" title="404" sub-title="页面不存在"></Result>
</template>

<style lang="less" scoped>
.full {
	position: absolute;
	left: 12px;
	right: 12px;
	top: 12px;
	bottom: 12px;
	border-radius: 4px;
	background-color: white;
}
</style>
