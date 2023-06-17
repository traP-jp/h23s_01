<script setup>
import { ref,onMounted } from "vue";
import { countDownTimer } from "../store.js";
import { gameTimer } from "../store.js";
import TimeBar from "../components/TimeBar.vue";

const countDown = () => {
  let timer = setInterval(() => {
    countDownTimer.value--;
    if (countDownTimer.value === -1) {
      clearInterval(timer);
      Timer(); 
    }
  }, 1000);
};

const Timer = () => {
  let timer = setInterval(() => {
    gameTimer.value--;
    if (gameTimer.value === 0) {
      clearInterval(timer);
    }
  }, 1000);
};

onMounted(() => {
  countDown();
});



</script>

<template>
  <div class="category">
    <h1>Game</h1>
  </div>
  
  <div>
    <TimeBar />
  </div>
  <div class="start">
    <p v-if="countDownTimer > 0" style="font-size: 64px;">{{ countDownTimer }}</p>
    <p v-if="countDownTimer == 0" style="font-size: 64px;">START!</p>
  </div>

  <div class="game">
    
  </div>
</template>

<style>
.start {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 50vh;
  text-align: center;
  font-size: 24px;
  font-weight: bold;
}

</style>
