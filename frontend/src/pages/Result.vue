<script setup>
import { onMounted } from "vue";
import {
  resultIkaList,
  resultShikaList,
  resultMekaList,
  ikaScore,
  shikaScore,
  mekaScore,
  totalScore,
  status,
} from "../store.js";
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
    <div class="scoreboard">
      <h1>トータルスコア: {{ totalScore }}</h1>
      <h2>
        いかスコア：{{ ikaScore }}　しかスコア：{{ shikaScore }}　めかスコア：{{
          mekaScore
        }}
      </h2>
    </div>
    <div class="messages_title">今回出会ったメッセージ</div>
    <div class="messages_description">
      traQのチャンネルにアクセスしてみよう！
    </div>
    <div class="message_columns">
      <CardColumn
        :messageList="resultIkaList.reverse()"
        color="#f0f2f5"
        title="いか"
        type="result"
      />
      <CardColumn
        :messageList="resultShikaList.reverse()"
        color="#6b7d8a"
        title="しか"
        type="result"
      />
      <CardColumn
        :messageList="resultMekaList.reverse()"
        color="#f0f2f5"
        title="めか"
        type="result"
      />
    </div>
    <div class="back_button_container">
      <button class="back_button" @click="status = 'title'">
        タイトルに戻る
      </button>
      <button class="back_button" @click="status = 'game'">もういちど</button>
    </div>
  </div>
</template>
<style scoped lang="scss">
.message_columns {
  display: flex;
  justify-content: space-around;
  width: 100%;
  margin-bottom: 60px;
}
.scoreboard {
  background-color: #e0e2ff;
  margin: 60px;
  padding: 32px;
  border-radius: 10px;
}
.messages_title {
  font-size: 32px;
  font-weight: bold;
  margin: 20px;
}
.messages_description {
  font-size: 20px;
  margin-bottom: 20px;
}
.back_button_container {
  display: flex;
  justify-content: center;
  .back_button {
    background-color: #005bac;
    color: white;
    font-weight: bold;
    font-size: 20px;
    padding: 10px 30px;
    margin: 0 20px;
    border-radius: 10px;
    &:hover {
      background-color: #0066cc;
    }
  }
}
</style>
