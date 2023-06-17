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
      content: "いかしかめか大学" + i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
    allList.value.push({
      user: "aya_se",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content:
        "東京工業大学は、東京都目黒区大岡山に本部を置く日本の国立大学である。" +
        i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: false,
      shika: false,
      meka: false,
    });
  }
  for (let i = 0; i < 200; i += 3) {
    leftList.value.push(allList.value[i]);
  }
  for (let i = 1; i < 200; i += 3) {
    middleList.value.push(allList.value[i]);
  }
  for (let i = 2; i < 200; i += 3) {
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
