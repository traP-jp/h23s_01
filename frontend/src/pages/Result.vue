<script setup>
import { messages } from "../store.js";
import CardColumn from "../components/CardColumn.vue";
</script>
<template>
  <div>
    <h1>Result</h1>
    <div class="message_columns">
      <CardColumn :messageList="messages.allList" color="#F0F2F5" />
      <CardColumn :messageList="messages.allList" color="white" />
      <CardColumn :messageList="messages.allList" color="#F0F2F5" />
    </div>
  </div>
</template>
<style scoped lang="scss">
.message_columns {
  display: flex;
  justify-content: space-around;
  width: 100%;
}
</style>
