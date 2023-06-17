<script setup>
import { status } from "../store";

const play = () => {
  status.value = "game";
};
</script>
<template>
  <div>
    <div>
      <h1>Title</h1>
      <button @click="play" class="play">Play</button>
    </div>
    <div class="howtoplay">
      <h2>遊び方</h2>
      <div class="howtoplaytext">
        <p>
          1.一定の間隔でtraQのgps/times以下のイカした投稿が流れてきます。<br />
          2.流れてくる投稿のうち，<hilight>「いか」「しか」「めか」</hilight>のどれかが含まれた投稿を見つけて選択すればポイントが得られます。<br />(ひらがな，カタカタ，漢字が対象です。例えば「
          医科歯科」，「いい感じ」，「イカしてる」などです。)<br />
          3.「いか」「しか」「めか」のどれも含まれない投稿を選択すると，<hilight>3秒間投稿を選択することができなく</hilight>なります。<br />
          4.制限時間が終了するとリザルト画面に移動し，スコアと「いか」「しか」「めか」が含まれた投稿の一覧が表示されます。<br />
          <hilight2>リザルト画面からtraQの投稿に飛ぶことができるので，素敵な投稿を見つけたら是非timesに遊びに行きましょう!</hilight2>
        </p>
      </div>
    </div>
  </div>
</template>

<style>
.play {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 24px;
  width: 128px;
  padding: 4px;
}
.howtoplay {
  padding: 10px;
  padding-top: 0px;
  border: solid 3px #646fcc;
  border-radius: 10px;
  margin: 30px;
  width: 895px;
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
