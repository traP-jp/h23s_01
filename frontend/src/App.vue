<script setup>
import {
  status,
  allList,
  leftList,
  middleList,
  rightList,
  user_id,
  user_name,
} from "./store.js";
import { onMounted } from "vue";
import Title from "./pages/Title.vue";
import Game from "./pages/Game.vue";
import Result from "./pages/Result.vue";
import axios from "axios";

onMounted(() => {
  const searchParams = new URLSearchParams(window.location.search);
  if (searchParams.toString() !== "") {
    axios
      .get("http://localhost:8080/api/me", {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res.data);
        user_id.value = res.data.id;
        user_name.value = res.data.name;
      })
      .catch(() => {
        axios
          .get("http://localhost:8080/api/oauth2/callback?" + searchParams, {
            withCredentials: true,
          })
          .then(
            axios
              .get("http://localhost:8080/api/me", {
                withCredentials: true,
              })
              .then((res) => {
                console.log(res.data);
                user_id.value = res.data.id;
                user_name.value = res.data.name;
              })
          );
      });
  } else {
    axios
      .get("http://localhost:8080/api/me", {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res.data);
        user_id.value = res.data.id;
        user_name.value = res.data.name;
      })
      .catch(
        axios.get("http://localhost:8080/api/oauth2/authorize", {
          withCredentials: true,
        })
      );
  }
  for (let i = 1; i < 100; i++) {
    allList.value.push({
      user: "ikura-hamu",
      messageId: i,
      channel: "gps/times/ikura-hamu",
      content: "いかしかめか大学",
      createdAt: "2023-06-14T05:46:02.585Z",
      ika: true,
      shika: true,
      meka: true,
    });
  }
  for (let i = 0; i < 33; i++) {
    leftList.value.push(allList.value[i]);
  }
  for (let i = 33; i < 66; i++) {
    middleList.value.push(allList.value[i]);
  }
  for (let i = 66; i < 99; i++) {
    rightList.value.push(allList.value[i]);
  }
});
</script>

<template>
  <div>
    <div v-if="!user_id">
      <a href="http://localhost:8080/api/oauth2/authorize">認証用リンク</a>
    </div>
    <div v-else>
      <select v-model="status">
        <option value="title">Title</option>
        <option value="game">Game</option>
        <option value="result">Result</option>
      </select>
      <div v-if="status === 'title'">
        <Title />
      </div>
      <div v-if="status === 'game'">
        <Game />
      </div>
      <div v-if="status === 'result'">
        <Result />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
