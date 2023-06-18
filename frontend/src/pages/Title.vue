<script setup>
import { status } from "../store";
import { user_name } from "../store";
import clicksound from "../sound/clicksound5.mp3";

const play = () => {
  const clickSound = new Audio(clicksound);
  clickSound.volume = 0.2;
  clickSound.play();
  status.value = "game";
};
</script>
<template>
  <div>
    <h1>いかしかめかアクティビティズ</h1>
    <div class="user_info">
      <img
        class="user_icon"
        :src="`https://q.trap.jp/api/v3/public/icon/${user_name}`"
      />
      <div class="user_name">{{ user_name }}</div>
    </div>
    <button @click="play" class="play">あそぶ</button>
  </div>
</template>

<style scoped lang="scss">
.play {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 32px;
  padding: 10px 60px;
  border-radius: 10px;
  &:hover {
    background-color: #0066cc;
  }
}
.user_info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 40px;
  .user_icon {
    width: 128px;
    height: 128px;
    border-radius: 50%;
    margin: 20px 0;
    object-fit: cover;
    background-color: #f0f2f5;
  }
  .user_name {
    font-size: 32px;
    font-weight: bold;
  }
}
</style>
