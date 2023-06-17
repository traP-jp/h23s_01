<script setup>
import { countDownTimer } from ".././store.js";
</script>
<template>
  <div class="countdown_text" v-if="countDownTimer > 0">
    {{ countDownTimer }}
  </div>
  <div class="countdown_text" v-if="countDownTimer == 0">START!</div>
</template>
<style scoped lang="scss">
.countdown_text {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 50vh;
  text-align: center;
  font-size: 100px;
  font-weight: bold;
}
</style>
