<script>
import {
  leftList,
  middleList,
  rightList,
  countDownTimer,
  gameTimer,
} from ".././store.js";
import { ref, onMounted } from "vue";
import TimeBar from "../components/TimeBar.vue";

export default {
  setup() {
    const currentLeftIndex = ref(0);
    const currentMiddleIndex = ref(0);
    const currentRightIndex = ref(0);

    const randomInterval = () => {
      return Math.floor(Math.random() * 5) * 1000;
    };

    const currentLeftMessageList = ref([]);
    const currentMiddleMessageList = ref([]);
    const currentRightMessageList = ref([]);

    onMounted(() => {
      addLeftMessage();
      addMiddleMessage();
      addRightMessage();
    });

    const addLeftMessage = () => {
      if (currentLeftIndex.value >= leftList.value.length) {
        return; // 終了条件
      }
      currentLeftMessageList.value.push(leftList.value[currentLeftIndex.value]);
      currentLeftIndex.value++;
      setTimeout(addLeftMessage, randomInterval());
    };

    const addMiddleMessage = () => {
      if (currentMiddleIndex.value >= middleList.value.length) {
        return; // 終了条件
      }
      currentMiddleMessageList.value.push(
        middleList.value[currentMiddleIndex.value]
      );
      currentMiddleIndex.value++;
      setTimeout(addMiddleMessage, randomInterval());
    };

    const addRightMessage = () => {
      if (currentRightIndex.value >= rightList.value.length) {
        return; // 終了条件
      }
      currentRightMessageList.value.push(
        rightList.value[currentRightIndex.value]
      );
      currentRightIndex.value++;
      setTimeout(addRightMessage, randomInterval());
    };

    return {
      currentLeftMessageList,
      currentMiddleMessageList,
      currentRightMessageList,
    };
  },
};

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
    <div class="lists-container">
      <div class="list">
        <div>leftMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentLeftMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>

      <div class="list">
        <div>middleMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentMiddleMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>

      <div class="list">
        <div>rightMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentRightMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>
    </div>
    <p>{{ messages.allList }}</p>
  </div>

  <div>
    <TimeBar />
  </div>
  <div class="start">
    <p v-if="countDownTimer > 0" style="font-size: 64px">
      {{ countDownTimer }}
    </p>
    <p v-if="countDownTimer == 0" style="font-size: 64px">START!</p>
  </div>

  <div class="game"></div>
</template>

<style>
.lists-container {
  display: flex;
}
.list {
  flex: 1;
  margin-right: 50px;
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
</style>
