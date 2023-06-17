<script setup>
import { status } from "./store.js";
import Title from "./pages/Title.vue";
import Game from "./pages/Game.vue";
import Result from "./pages/Result.vue";
import { onMounted } from "vue";
import { messages } from "./store.js";
onMounted(() => {
  for (let i = 0; i < 100; i++) {
    messages.allList.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "いかしかめか大学",
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
  }
});
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
