<script setup>
import { gameTimer } from "../store";
</script>

<template>
    <p>{{ gameTimer }}</p>
    <div class="progress-bar" :style="{ width: gameTimer/30 * 100 + '%' }"></div>
</template>

<style>
.progress-bar {
  height: 20px;
  background-color: #323232;
  transition: width 0.3s;
}
</style>