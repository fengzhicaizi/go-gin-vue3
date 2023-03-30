<script setup lang="ts">
import { onMounted } from 'vue';
import { startCesium } from './cesium';

onMounted(() => {
	// startCesium();
});
</script>

<template>
	<div id="full-screen" class="cesium"></div>
</template>

<style lang="less" scoped>
.cesium {
	width: 100%;
	height: 100%;
	position: relative;
}
</style>
