<script>
import { leftList, middleList, rightList } from ".././store.js";
import { ref, watch, onMounted } from "vue";

export default {
  setup() {
    const currentLeftIndex = ref(0);
    const currentMiddleIndex = ref(0);
    const currentRightIndex = ref(0);

    const randomInterval = () => {
      return Math.floor(Math.random() * 5) * 1000;
    };

    const currentLeftMessageList = ref([]);
    const currentMiddleMessageList = ref([]);
    const currentRightMessageList = ref([]);

    onMounted(() => {
      addLeftMessage();
      addMiddleMessage();
      addRightMessage();
    });

    const addLeftMessage = () => {
      if (currentLeftIndex.value >= leftList.value.length) {
        return; // 終了条件
      }
      currentLeftMessageList.value.push(leftList.value[currentLeftIndex.value]);
      currentLeftIndex.value++;
      setTimeout(addLeftMessage, randomInterval());
    };

    const addMiddleMessage = () => {
      if (currentMiddleIndex.value >= middleList.value.length) {
        return; // 終了条件
      }
      currentMiddleMessageList.value.push(
        middleList.value[currentMiddleIndex.value]
      );
      currentMiddleIndex.value++;
      setTimeout(addMiddleMessage, randomInterval());
    };

    const addRightMessage = () => {
      if (currentRightIndex.value >= rightList.value.length) {
        return; // 終了条件
      }
      currentRightMessageList.value.push(
        rightList.value[currentRightIndex.value]
      );
      currentRightIndex.value++;
      setTimeout(addRightMessage, randomInterval());
    };

    return {
      currentLeftMessageList,
      currentMiddleMessageList,
      currentRightMessageList,
    };
  },
};
</script>
<template>
  <div>
    <h1>Game</h1>

    <div class="lists-container">
      <div class="list">
        <div>leftMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentLeftMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>

      <div class="list">
        <div>middleMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentMiddleMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>

      <div class="list">
        <div>rightMessages</div>
        <ul style="text-align: left">
          <li v-for="message in currentRightMessageList" :key="message">
            <div>{{ message }}</div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style>
.lists-container {
  display: flex;
}

.list {
  flex: 1;
  margin-right: 50px;
}
</style>
