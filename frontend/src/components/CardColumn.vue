<script setup>
import MessageCard from "./MessageCard.vue";
const props = defineProps({
  messageList: {
    type: Array,
    required: true,
  },
  color: {
    type: String,
    required: true,
  },
  title: {
    type: String,
  },
  type: {
    type: String,
    required: true,
  },
});
</script>
<template>
  <div class="card_column">
    <div class="column_title">{{ title }}</div>
    <div v-if="type === 'game'" :class="['column_contents', 'game_column']">
      <TransitionGroup name="message_cards">
        <template v-for="message in messageList" :key="message.messageId">
          <MessageCard :message="message" :type="type" />
        </template>
      </TransitionGroup>
    </div>
    <div v-if="type === 'result'" :class="['column_contents', 'result_column']">
      <TransitionGroup name="message_cards">
        <template v-for="message in messageList" :key="message.messageId">
          <MessageCard :message="message" :type="type" />
        </template>
      </TransitionGroup>
    </div>
  </div>
</template>
<style scoped lang="scss">
.title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
}
.card_column {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  width: calc(100% / 3);
  .column_title {
    width: 100%;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 10px;
  }
  .column_contents {
    width: 100%;
    background-color: v-bind(color);
    padding: 5px 20px;
    .messsage_cards-move {
      transition: all 1s ease;
    }
    .message_cards-enter-active {
      transition: all 0.5s ease;
    }
    .message_cards-leave-active {
      transition: all 0.5s ease;
    }
    .message_cards-enter-from {
      opacity: 0;
      transform: translateX(10px);
    }
  }
  .game_column {
    height: 550px;
    overflow-y: hidden;
  }
  .result_column {
    height: 550px;
    overflow-y: scroll;
  }
}
</style>
