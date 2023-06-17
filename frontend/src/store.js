import { reactive, ref, computed } from "vue";

// 認証の状態を管理する
export const auth = reactive({
  //TODO: 後で足す
});

// ゲームの状態を管理する
export const status = ref("title");

// タイマーを管理する
export const countDownTimer = ref(3);
export const gameTimer = ref(30);

// 得点を管理する
export const score = reactive({
  ika: 0,
  shika: 0,
  meka: 0,
  total: computed(() => {
    return score.ika + score.shika + score.meka;
  }),
  highest: 0,
});

export const allList = ref([]);
export const correctList = computed(() => {
  return messages.allList.filter((message) => {
    return message.ika || message.shika || message.meka;
  });
});
export const incorrectList = computed(() => {
  return messages.allList.filter((message) => {
    return !message.ika && !message.shika && !message.meka;
  });
});
export const leftList = ref([]);
export const middleList = ref([]);
export const rightList = ref([]);
export const resultIkaList = ref([]);
export const resultShikaList = ref([]);
export const resultMekaList = ref([]);

// ランキングを管理する
export const ranking = reactive({
  list: [],
});
