<script setup lang="ts">
import { onMounted } from 'vue';
import { start } from './index';

onMounted(() => {
	start();
});
</script>

<template>
	<div id="full-screen" class="map"></div>
</template>

<style lang="less" scoped>
.map {
	width: 100%;
	height: 100%;
	position: relative;
}
</style>
