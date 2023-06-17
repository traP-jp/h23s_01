<script setup>
import { status } from "../store";

const play = () => {
  status.value = "game";
};
</script>
<template>
  <div>
    <h1>Title</h1>
    <button @click="play" class="play">Play</button>
  </div>
</template>

<style>
.play {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 24px;
  width: 128px;
  padding: 4px;
}
</style>
