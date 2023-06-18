<script setup>
import { status } from "../store";
import { user_name } from "../store";
//import clicksound from "../sound/clicksound5.mp3";

const play = () => {
  //const clickSound = new Audio(clicksound);
  //clickSound.volume = 0.2;
  //clickSound.play();
  status.value = "game";
};
</script>
<template>
  <div>
    <h1>いかしかめかアクティビティズ</h1>
    <div class="user_info">
      <img
        class="user_icon"
        :src="`https://q.trap.jp/api/v3/public/icon/${user_name}`"
      />
      <div class="user_name">{{ user_name }}</div>
    </div>
    <button @click="play" class="play">あそぶ</button>
    <div class="howtoplay">
      <h2>遊び方</h2>
      <div class="howtoplaytext">
        <p>
          1.一定の間隔でtraQのgps/times以下のイカした投稿が流れてきます。<br />
          2.流れてくる投稿のうち，<hilight>「いか」「しか」「めか」</hilight>のどれかが含まれた投稿を見つけて選択すればポイントが得られます。<br />(ひらがな，カタカタ，漢字が対象です。例えば「
          医科歯科」，「いい感じ」，「イカしてる」などです。)<br />
          3.「いか」「しか」「めか」のどれも含まれない投稿を選択すると，<hilight>3秒間投稿を選択することができなく</hilight>なります。<br />
          4.制限時間が終了するとリザルト画面に移動し，スコアと「いか」「しか」「めか」が含まれた投稿の一覧が表示されます。<br />
          <hilight2
            >リザルト画面からtraQの投稿に飛ぶことができるので，素敵な投稿を見つけたら是非timesに遊びに行きましょう!</hilight2
          >
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.play {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 32px;
  padding: 10px 60px;
  border-radius: 10px;
  &:hover {
    background-color: #0066cc;
  }
}
.user_info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 40px;
  .user_icon {
    width: 128px;
    height: 128px;
    border-radius: 50%;
    margin: 20px 0;
    object-fit: cover;
    background-color: #f0f2f5;
  }
  .user_name {
    font-size: 32px;
    font-weight: bold;
  }
}
.howtoplay {
  padding: 0 20px 10px;
  border: solid 3px #005bac;
  border-radius: 10px;
  margin: 40px;
  width: 900px;
}
.howtoplaytext {
  text-align: left;
  line-height: 30px;
}
hilight {
  color: #f01d2e;
  font-weight: bold;
}
hilight2 {
  color: #f5b847;
  font-weight: bold;
}
</style>
