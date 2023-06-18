<script setup>
import { status, user_id, user_name, API_URL } from "./store.js";
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
