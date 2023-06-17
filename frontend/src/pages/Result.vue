<script setup>
import { onMounted } from "vue";
import { resultIkaList, resultShikaList, resultMekaList } from "../store.js";
import CardColumn from "../components/CardColumn.vue";

onMounted(() => {
  for (let i = 0; i < 10; i++) {
    resultIkaList.value.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "いか大学",
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
    resultShikaList.value.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "しか大学",
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
    resultMekaList.value.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "めか大学",
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
  }
});
</script>
<template>
  <div>
    <h1>Result</h1>
    <div class="message_columns">
      <CardColumn
        :messageList="resultIkaList.reverse()"
        color="#f0f2f5"
        title="いか"
      />
      <CardColumn
        :messageList="resultShikaList.reverse()"
        color="#6b7d8a"
        title="しか"
      />
      <CardColumn
        :messageList="resultMekaList.reverse()"
        color="#f0f2f5"
        title="めか"
      />
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
