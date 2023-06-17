<script setup>
const props = defineProps({
  message: {
    type: Object,
    required: true,
  },
});
</script>
<template>
  <div class="message_card">
    <div>{{ message.user }}</div>
    <div>#&nbsp;{{ message.channel }}</div>
    <div>{{ message.content }}</div>
  </div>
</template>
<style scoped lang="scss">
.message_card {
  width: calc(100% - 40px);
  border-radius: 5px;
  padding: 10px;
  margin: 10px;
  background-color: white;
}
</style>
