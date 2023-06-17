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
onMounted(() => {});
</script>
<template>
  <div>
    <h1>リザルト</h1>
    <div class="scoreboard">
      <div class="scoreboard_title">
        トータルスコア:&nbsp;&nbsp;{{ totalScore }}
      </div>
      <div class="scoreboard_detail">
        <img class="scoreboard_detail_icon" src="/ika.svg" />
        <div class="scoreboard_detail_score">{{ ikaScore }}</div>
        <img class="scoreboard_detail_icon" src="/shika.svg" />
        <div class="scoreboard_detail_score">{{ shikaScore }}</div>
        <img class="scoreboard_detail_icon" src="/meka.svg" />
        <div class="scoreboard_detail_score">{{ mekaScore }}</div>
      </div>
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
        icon="/ika.svg"
        type="result"
      />
      <CardColumn
        :messageList="resultShikaList.reverse()"
        color="#6b7d8a"
        title="しか"
        icon="/shika.svg"
        type="result"
      />
      <CardColumn
        :messageList="resultMekaList.reverse()"
        color="#f0f2f5"
        title="めか"
        icon="/meka.svg"
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
  .scoreboard_title {
    font-size: 40px;
    font-weight: bold;
    margin-bottom: 20px;
  }
  .scoreboard_detail {
    display: flex;
    justify-content: center;
    align-items: center;
    .scoreboard_detail_icon {
      width: 64px;
      height: 64px;
      margin: 0 20px;
    }
    .scoreboard_detail_score {
      font-size: 32px;
      font-weight: bold;
      margin: 0 20px;
    }
  }
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
