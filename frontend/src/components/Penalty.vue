<script setup>
import { penaltyTimer } from ".././store.js";
</script>
<template>
  <div class="penalty">
    <div class="penalty_text">Loading...</div>
    <div class="penalty_timer">{{ penaltyTimer }}</div>
  </div>
</template>
<style scoped lang="scss">
.penalty {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  font-size: 32px;
  font-weight: bold;
  background-color: #f0f2f5;
  z-index: 100;
  opacity: 0.8;
  .penalty_text {
    opacity: 1;
  }
  .penalty_timer {
    font-size: 64px;
    font-weight: bold;
    color: #005bac;
    opacity: 1;
  }
}
</style>
