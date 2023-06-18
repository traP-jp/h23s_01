<script setup>
import { onMounted } from "vue";
import { status } from "../store.js";

onMounted(() => {
  // 3秒後にリザルト画面に遷移
  setTimeout(() => {
    status.value = "result";
  }, 3000);
});
</script>
<template>
  <div class="gameover">
    <div class="gameover_text">ゲーム終了！</div>
  </div>
</template>
<style scoped lang="scss">
.gameover {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  background-color: rgb(240, 242, 245, 0.8);
  z-index: 100;
  .gameover_text {
    font-size: 48px;
    font-weight: bold;
    color: #005bac;
  }
}
</style>
