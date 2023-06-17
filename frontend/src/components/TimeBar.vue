<script setup>
import { gameTimer } from "../store";
import { computed } from "vue";
const barColor = computed(() => {
  if (gameTimer.value > 10) {
    return "#005bac";
  } else if (gameTimer.value > 5) {
    return "#ffcc00";
  } else {
    return "#ff0000";
  }
});
</script>

<template>
  <div style="display: flex">
    <h2 class="timer_text" style="margin: auto">{{ gameTimer }}</h2>
    <div class="remaining_bar">
      <div class="remaining" :style="`width: ${(gameTimer / 30) * 100}%`"></div>
    </div>
  </div>
</template>

<style>
.remaining_bar {
  width: 80%;
  height: 20px;
  margin: auto;
  background-color: lightgray;
  border-radius: 10px;
  overflow: hidden;
}

.timer_text {
  font-size: 40px;
  font-weight: bold;
  color: v-bind(barColor);
  margin: auto;
}

.remaining {
  height: 100%;
  background-color: v-bind(barColor);
}
</style>
