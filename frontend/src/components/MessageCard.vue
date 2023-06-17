<script setup>
const props = defineProps({
  message: {
    type: Object,
    required: true,
  },
});
const icon = "https://q.trap.jp/api/v3/public/icon/ikura-hamu.png";
</script>
<template>
  <div class="message_card">
    <div class="message_header">
      <image class="message_icon" :src="icon" />
      <div class="message_user">{{ message.user }}</div>
    </div>
    <div class="message_separator" />
    <div class="message_channel">#&nbsp;{{ message.channel }}</div>
    <div class="message_text">{{ message.content }}</div>
  </div>
</template>
<style scoped lang="scss">
.message_card {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  width: calc(100% - 40px);
  border-radius: 5px;
  padding: 10px 20px;
  margin: 10px 0;
  background-color: white;
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
</style>
