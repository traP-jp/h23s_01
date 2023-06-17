<template>
  <div class="background-container"></div>
</template>
<style scoped lang="scss">
.background-container {
  position: fixed;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  overflow: hidden;
  z-index: -1;
  transform: rotate(-5deg);
}

@keyframes moveBackground {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 1000% 1000%;
  }
}
.background-container {
  background-image: url("ikashikameka.jpg");
  background-repeat: repeat;
  animation: fadeIn 1s ease-in-out, moveBackground 720s linear infinite;
}
/* フェードインアニメーションを作成 */
@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>
