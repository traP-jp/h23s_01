<script setup>
import { ref } from "vue";
import {
  resultIkaList,
  resultShikaList,
  resultMekaList,
  isPenalty,
  penaltyTimer,
} from ".././store.js";
const props = defineProps({
  message: {
    type: Object,
    required: true,
  },
  type: {
    type: String,
    required: true,
  },
});
const icon = ref(
  `https://q.trap.jp/api/v3/public/icon/${props.message.user}.png`
);
const isCorrect = ref(false);
const isIncorrect = ref(false);

const onClickHandler = () => {
  if (props.type === "game") {
    if (props.message.ika || props.message.shika || props.message.meka) {
      isCorrect.value = true;
      if (props.message.ika) {
        resultIkaList.value.push(props.message);
      }
      if (props.message.shika) {
        resultShikaList.value.push(props.message);
      }
      if (props.message.meka) {
        resultMekaList.value.push(props.message);
      }
    } else {
      isIncorrect.value = true;
      isPenalty.value = true;
      penaltyCount();
    }
  } else if (props.type === "result") {
    // traQのメッセージを新しいタブで開く
    window.open(
      `https://q.trap.jp/messages/${props.message.messageId}`,
      "_blank"
    );
  }
};

// isPenaltyがTrueに切り替わったらカウントダウン
const penaltyCount = () => {
  penaltyTimer.value = 3;
  let timer = setInterval(() => {
    penaltyTimer.value--;
    if (penaltyTimer.value === 0) {
      clearInterval(timer);
      isPenalty.value = false;
    }
  }, 1000);
};
</script>
<template>
  <button
    :disabled="isCorrect || isIncorrect"
    :class="[
      'message_card',
      { correct_card: isCorrect },
      { incorrect_card: isIncorrect },
    ]"
    @click="onClickHandler()"
  >
    <div class="message_header">
      <image class="message_icon" :src="icon" />
      <div class="message_user">{{ message.user }}</div>
    </div>
    <div class="message_separator" />
    <div class="message_channel">#&nbsp;{{ message.channel }}</div>
    <div class="message_text">{{ message.content }}</div>
  </button>
</template>
<style scoped lang="scss">
.message_card {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  width: 100%;
  border-radius: 5px;
  padding: 10px 20px;
  margin: 10px 0;
  background-color: white;
  border: 2px solid white;
  &:hover {
    cursor: pointer;
    border: 2px solid #005bac;
  }
  &:visited {
    border: 2px solid white;
  }
  .message_header {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: left;
    align-items: center;
    .message_icon {
      width: 32px;
      height: 32px;
      border-radius: 50%;
      margin-right: 10px;
      object-fit: cover;
      background-color: #f0f2f5;
    }
    .message_user {
      font-size: 16px;
      font-weight: bold;
    }
  }
  .message_separator {
    width: 100%;
    height: 2px;
    background-color: #f0f2f5;
    margin: 5px 0;
  }
  .message_channel {
    width: 100%;
    text-align: left;
    font-size: 14px;
    color: #6b7d8a;
    margin-bottom: 5px;
  }
  .message_text {
    width: 100%;
    text-align: left;
    font-size: 16px;
  }
}
.correct_card {
  border-color: #00ff00 !important;
}
.incorrect_card {
  border-color: #ff0000 !important;
}
</style>
