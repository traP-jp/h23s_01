<script setup>
import {
  leftList,
  middleList,
  rightList,
  countDownTimer,
  gameTimer,
  isPenalty,
} from ".././store.js";
import { ref, onMounted } from "vue";
import TimeBar from "../components/TimeBar.vue";
import CardColumn from "../components/CardColumn.vue";
import Penalty from "../components/Penalty.vue";

onMounted(() => {
  addLeftIndex();
  addMiddleIndex();
  addRightIndex();
  countDown();
});

// 時間管理
const countDown = () => {
  countDownTimer.value = 3;
  gameTimer.value = 30;
  let timer = setInterval(() => {
    countDownTimer.value--;
    if (countDownTimer.value === -1) {
      clearInterval(timer);
      Timer();
      isStarted.value = true;
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

const isStarted = ref(false);

// メッセージ追加の管理
const leftIndex = ref(0);
const middleIndex = ref(0);
const rightIndex = ref(0);

const randomInterval = () => {
  return 1000 + Math.floor(Math.random() * 1000);
};

const addLeftIndex = () => {
  if (leftIndex.value >= leftList.value.length) {
    return; // 終了条件
  }
  leftIndex.value++;
  setTimeout(addLeftIndex, randomInterval());
};

const addMiddleIndex = () => {
  if (middleIndex.value >= middleList.value.length) {
    return; // 終了条件
  }
  middleIndex.value++;
  setTimeout(addMiddleIndex, randomInterval());
};

const addRightIndex = () => {
  if (rightIndex.value >= rightList.value.length) {
    return; // 終了条件
  }
  rightIndex.value++;
  setTimeout(addRightIndex, randomInterval());
};
</script>

<template>
  <div v-if="!isStarted" class="count_down">
    <p v-if="countDownTimer > 0" style="font-size: 64px">
      {{ countDownTimer }}
    </p>
    <p v-if="countDownTimer == 0" style="font-size: 64px">START!</p>
  </div>
  <div v-else class="game">
    <Penalty v-if="isPenalty" />
    <TimeBar />
    <div class="category">
      <div class="message_columns">
        <CardColumn
          :messageList="leftList.slice(0, leftIndex).reverse()"
          color="#f0f2f5"
          title="Left"
          type="game"
        />
        <CardColumn
          :messageList="middleList.slice(0, middleIndex).reverse()"
          color="#6b7d8a"
          title="Middle"
          type="game"
        />
        <CardColumn
          :messageList="rightList.slice(0, rightIndex).reverse()"
          color="#f0f2f5"
          title="Right"
          type="game"
        />
      </div>
    </div>
  </div>
</template>

<style>
.message_columns {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
}
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
.game {
  position: relative;
}
</style>
