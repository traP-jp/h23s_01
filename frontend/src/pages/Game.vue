<script setup>
import {
  leftList,
  middleList,
  rightList,
  countDownTimer,
  gameTimer,
  isPenalty,
  resultIkaList,
  resultShikaList,
  resultMekaList,
  ikaScore,
  shikaScore,
  mekaScore,
  allList,
  API_URL,
  status,
} from ".././store.js";
import { ref, onMounted, watch } from "vue";
import TimeBar from "../components/TimeBar.vue";
import CardColumn from "../components/CardColumn.vue";
import CountDown from "../components/CountDown.vue";
import Penalty from "../components/Penalty.vue";
import GameOver from "../components/GameOver.vue";
import axios from "axios";
import gamebgm from "../sound/bgm1.mp3";
import countsound from "../sound/countdown.mp3";
import finishwhistle from "../sound/finishwhistle.mp3";

const isStarted = ref(false);
const isEnded = ref(false);

const gameBgm = new Audio(gamebgm);
const countSound = new Audio(countsound);
const finishWhistle = new Audio(finishwhistle);

onMounted(() => {
  countDown();
  resultIkaList.value = [];
  resultShikaList.value = [];
  resultMekaList.value = [];
  ikaScore.value = 0;
  shikaScore.value = 0;
  mekaScore.value = 0;

  // messageをAPIから取得
  axios
    .get(`${API_URL}/api/message`, {
      withCredentials: true,
    })
    .then((res) => {
      console.log(res.data.messages);
      allList.value = shuffleArray(res.data.messages);
      const N = allList.value.length;
      leftList.value = allList.value.slice(0, N / 3);
      middleList.value = allList.value.slice(N / 3, (2 * N) / 3);
      rightList.value = allList.value.slice((2 * N) / 3, N);
    })
    .catch((err) => {
      console.log(err);
    });
});

const shuffleArray = (array) => {
  const cloneArray = [...array];

  for (let i = cloneArray.length - 1; i >= 0; i--) {
    let rand = Math.floor(Math.random() * (i + 1));
    // 配列の要素の順番を入れ替える
    let tmpStorage = cloneArray[i];
    cloneArray[i] = cloneArray[rand];
    cloneArray[rand] = tmpStorage;
  }

  return cloneArray;
};
// 時間管理
const countDown = () => {
  countDownTimer.value = 3;
  gameTimer.value = 30;
  let timer = setInterval(() => {
    countDownTimer.value--;
    if (countDownTimer.value === 2) {
      countSound.play(); // 音源が2秒分しかないのでここに入れました
    }
    if (countDownTimer.value === -1) {
      clearInterval(timer);
      isStarted.value = true;
    }
  }, 1000);
};

// ゲームが開始＆allListの取得が完了したらwatchでゲーム時間のカウントを開始する
watch(
  () => isStarted.value && allList.value.length > 0,
  () => {
    gameBgm.play();
    addLeftIndex();
    addMiddleIndex();
    addRightIndex();
    let timer = setInterval(() => {
      gameTimer.value--;
      if (gameTimer.value === 0) {
        clearInterval(timer);
        isEnded.value = true;
      }
    }, 1000);
  }
);

// ゲームが終了したらリザルト画面への遷移を行う
watch(isEnded, () => {
  if (isEnded.value) {
    //gameBgm.pause();
    //gameBgm.currentTime = 0; // bgm停止,終了の合図用の音源再生
    finishWhistle.play();
    // 3秒後にリザルト画面に遷移
    setTimeout(() => {
      status.value = "result";
    }, 3000);
  }
});
// メッセージ追加の管理
const leftIndex = ref(0);
const middleIndex = ref(0);
const rightIndex = ref(0);

const randomInterval = () => {
  return 1000 + Math.floor(Math.random() * 2000);
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
    <CountDown />
  </div>
  <div v-else class="game">
    <GameOver v-if="gameTimer === 0" />
    <Penalty v-else-if="isPenalty" />
    <TimeBar />

    <div class="message_columns">
      <CardColumn
        :messageList="leftList.slice(0, leftIndex).reverse()"
        color="#f0f2f5"
        title=""
        type="game"
      />
      <CardColumn
        :messageList="middleList.slice(0, middleIndex).reverse()"
        color="#6b7d8a"
        title=""
        type="game"
      />
      <CardColumn
        :messageList="rightList.slice(0, rightIndex).reverse()"
        color="#f0f2f5"
        title=""
        type="game"
      />
    </div>
  </div>
</template>

<style>
.message_columns {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  margin-top: 40px;
}
.game {
  position: relative;
}
</style>
