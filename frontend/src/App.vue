<script setup>
import { status } from "./store.js";
import Title from "./pages/Title.vue";
import Game from "./pages/Game.vue";
import Result from "./pages/Result.vue";
</script>

<template>
  <select v-model="status">
    <option value="title">Title</option>
    <option value="game">Game</option>
    <option value="result">Result</option>
  </select>
  <div v-if="status === 'title'">
    <Title />
  </div>
  <div v-if="status === 'game'">
    <Game />
  </div>
  <div v-if="status === 'result'">
    <Result />
  </div>
</template>

<style scoped lang="scss"></style>
