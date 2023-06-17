<script setup>
import { gameTimer } from "../store";
</script>

<template>
  <div style="display: flex">
    <h2 style="margin: auto">{{ gameTimer }}</h2>
    <div class="remaining-bar">
      <div class="remaining" :style="`width: ${(gameTimer / 30) * 100}%`"></div>
    </div>
  </div>
</template>

<style>
.remaining-bar {
  width: 80%;
  height: 20px;
  margin: auto;
  background-color: lightgray;
  border-radius: 10px;
  overflow: hidden;
}

.remaining {
  height: 100%;
  background-color: #646fcc;
}
</style>
