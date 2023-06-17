<script setup>
import { status, allList, leftList, middleList, rightList } from "./store.js";
import { onMounted } from "vue";
import Title from "./pages/Title.vue";
import Game from "./pages/Game.vue";
import Result from "./pages/Result.vue";

onMounted(() => {
  for (let i = 1; i < 100; i++) {
    allList.value.push({
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
  for (let i = 0; i < 33; i++) {
    leftList.value.push(allList.value[i]);
  }
  for (let i = 33; i < 66; i++) {
    middleList.value.push(allList.value[i]);
  }
  for (let i = 66; i < 99; i++) {
    rightList.value.push(allList.value[i]);
  }
});
</script>

<template>
  <div>
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
  </div>
</template>

<style scoped lang="scss"></style>
