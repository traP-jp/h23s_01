<script setup>
import MessageCard from "./MessageCard.vue";
const props = defineProps({
  messageList: {
    type: Array,
    required: true,
  },
  color: {
    type: String,
  },
});
</script>
<template>
  <div class="card_column">
    <div v-for="message in messageList" :key="message.messageId">
      <MessageCard :message="message" />
    </div>
  </div>
</template>
<style scoped lang="scss">
.card_column {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  width: calc(100% / 3);
  background-color: v-bind(color);
}
</style>
