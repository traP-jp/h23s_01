<script setup>
import { resultIkaList, resultShikaList, resultMekaList } from "../store.js";
import CardColumn from "../components/CardColumn.vue";
</script>
<template>
  <div>
    <h1>Result</h1>
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
  </div>
</template>
<style scoped lang="scss">
.message_columns {
  display: flex;
  justify-content: space-around;
  width: 100%;
}
</style>
