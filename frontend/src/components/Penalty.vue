<script setup>
import { penaltyTimer } from ".././store.js";
</script>
<template>
  <div class="penalty">
    <div class="penalty_text">復帰まで...</div>
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
  background-color: rgb(240, 242, 245, 0.8);
  z-index: 100;
  .penalty_text {
    font-size: 32px;
    font-weight: bold;
  }
  .penalty_timer {
    font-size: 64px;
    font-weight: bold;
    color: #005bac;
  }
}
</style>
