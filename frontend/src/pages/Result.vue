<script setup>
import { onMounted } from "vue";
import clicksound from "../sound/clicksound5.mp3";
import resultbgm from "../sound/bgm2.mp3";
import {
  resultIkaList,
  resultShikaList,
  resultMekaList,
  ikaScore,
  shikaScore,
  mekaScore,
  totalScore,
  status,
  user_name,
  API_URL,
  user_id,
} from "../store.js";
import CardColumn from "../components/CardColumn.vue";
import MessageCard from "../components/MessageCard.vue";
import axios from "axios";
onMounted(() => {});
const shareMessage = {
  messageId: "",
  type: "result",
  channel: "random/gamers/ikashikameka",
  content: `@${user_name.value} は「いかしかめかアクティビティズ」で${totalScore.value}点獲得しました！`,
  user: user_name.value,
};
const shareScore = () => {
  // axiosでメッセージ投稿
  /*
  await axios
    .post(`${API_URL}/api/score`, {
      withCredentials: true,
      body: {
        score: totalScore.value,
      },
    })
    .then(() => {
      alert("スコアをシェアしました！");
      axios
        .get(`${API_URL}/api/score/highest`, {
          withCredentials: true,
        })
        .then((res) => {
          console.log(res.data);
        });
    });
  */
  console.log(user_id.value);
  axios
    .post(`${API_URL}/api/post`, {
      withCredentials: true,
      id: user_id.value,
      score: totalScore.value,
    })
    .then(() => {
      alert("スコアをシェアしました！");
    });
};

const resultBgm = new Audio(resultbgm);

const replay = () => {
  resultBgm.pause();
  resultBgm.currentTime = 0;
  const clickSound = new Audio(clicksound);
  clickSound.volume = 0.2;
  clickSound.play();
  status.value = "game";
};

const toTitle = () => {
  resultBgm.pause();
  resultBgm.currentTime = 0;
  const clickSound = new Audio(clicksound);
  clickSound.volume = 0.2;
  clickSound.play();
  status.value = "title";
};

onMounted(() => {
  resultBgm.play();
});
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
    <div class="share_score">
      <div class="share_header">
        <img class="share_icon" src="/traq.png" />
        <div class="share_title">でスコアをシェアしよう！</div>
      </div>
      <MessageCard :message="shareMessage" :type="game" />
      <button class="submit_button" @click="shareScore()">シェアする</button>
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
      <button class="back_button" @click="toTitle">タイトルに戻る</button>
      <button class="back_button" @click="replay">もういちど</button>
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
  background-color: #d5e7ff;
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

.share_score {
  width: 550px;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin: 0 auto 60px;
  padding: 32px;
  background-color: #f0f2f5;
  border-radius: 10px;
  .share_header {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
    .share_icon {
      width: 64px;
      height: 64px;
      margin-right: 20px;
    }
    .share_title {
      font-size: 24px;
      font-weight: bold;
      margin-right: 20px;
    }
  }
  .submit_button {
    background-color: #005bac;
    color: white;
    font-weight: bold;
    font-size: 20px;
    padding: 10px 30px;
    margin: 20px 20px 0;
    border-radius: 10px;
    &:hover {
      background-color: #0066cc;
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
