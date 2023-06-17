<script setup>
import {
  status,
  allList,
  leftList,
  middleList,
  rightList,
  user_id,
  user_name,
  API_URL,
} from "./store.js";
import { onMounted } from "vue";
import Title from "./pages/Title.vue";
import Game from "./pages/Game.vue";
import Result from "./pages/Result.vue";
import Background from "./components/Background.vue";
import axios from "axios";

onMounted(() => {
  // traQ認証の実装
  const searchParams = new URLSearchParams(window.location.search);
  if (searchParams.toString() !== "") {
    axios
      .get(`${API_URL}/api/me`, {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res.data);
        user_id.value = res.data.id;
        user_name.value = res.data.name;
      })
      .catch((err) => {
        console.log(err);
        axios
          .get(`${API_URL}/api/oauth2/callback?${searchParams}`, {
            withCredentials: true,
          })
          .then((res) => {
            console.log(res);
            axios
              .get(`${API_URL}/api/me`, {
                withCredentials: true,
              })
              .then((res) => {
                console.log(res.data);
                user_id.value = res.data.id;
                user_name.value = res.data.name;
              });
          });
      });
  } else {
    axios
      .get(`${API_URL}/api/me`, {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res.data);
        user_id.value = res.data.id;
        user_name.value = res.data.name;
      })
      .catch((err) => {
        console.log(err);
        user_id.value = "unauthorized";
      });
  }
  for (let i = 0; i < 50; i++) {
    allList.value.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "いかのお寿司大学" + i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: false,
      meka: false,
    });
    allList.value.push({
      user: "masky5859",
      messageId: i,
      channel: "gps/times/masky5859",
      content: "かくかくしかじかエコカー減税ダイハツから" + i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: false,
      shika: true,
      meka: false,
    });
    allList.value.push({
      user: "Akira_256",
      messageId: i,
      channel: "gps/times/Akira_256",
      content: "ときめかないラブストーリー" + i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: false,
      shika: false,
      meka: true,
    });
    allList.value.push({
      user: "aya_se",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content:
        "東京工業大学は、東京都目黒区大岡山に本部を置く日本の国立大学である。" +
        i,
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: false,
      shika: false,
      meka: false,
    });
  }
  for (let i = 0; i < 200; i += 3) {
    leftList.value.push(allList.value[i]);
  }
  for (let i = 1; i < 200; i += 3) {
    middleList.value.push(allList.value[i]);
  }
  for (let i = 2; i < 200; i += 3) {
    rightList.value.push(allList.value[i]);
  }
});
</script>

<template>
  <div>
    <Background />
    <div v-if="user_id === 'unauthorized'">
      <a class="authorize_link" :href="`${API_URL}/api/oauth2/authorize`"
        >認証用リンク</a
      >
    </div>
    <div v-else-if="user_id">
      <div v-if="status === 'title'">
        <Title />
      </div>
      <div v-if="status === 'game'">
        <Game />
      </div>
      <div v-if="status === 'result'">
        <Result />
      </div>
      <select v-model="status">
        <option value="title">Title</option>
        <option value="game">Game</option>
        <option value="result">Result</option>
      </select>
    </div>
  </div>
</template>

<style scoped lang="scss">
.authorize_link {
  background-color: #005bac;
  color: white;
  font-weight: bold;
  font-size: 20px;
  width: 128px;
  padding: 10px 20px;
  border-radius: 10px;
  &:hover {
    background-color: #0066cc;
  }
}
</style>
