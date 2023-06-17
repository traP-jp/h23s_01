import { reactive, ref, computed } from "vue";

// 認証の状態を管理する
export const auth = reactive({
  //TODO: 後で足す
});

// ゲームの状態を管理する
export const status = ref("title");

// タイマーを管理する
export const countDownTimer = ref(3);
export const gameTimer = ref(30)
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

// メッセージリストを管理する
export const messages = reactive({
  allList: [],
  correctList: computed(() => {
    return messages.allList.filter((message) => {
      return message.ika || message.shika || message.meka;
    });
  }),
  incorrectList: computed(() => {
    return messages.allList.filter((message) => {
      return !message.ika && !message.shika && !message.meka;
    });
  }),

  // ゲーム画面用
  leftList: [],
  middleList: [],
  rightList: [],

  // リザルト画面用
  resultIkaList: [],
  resultShikaList: [],
  resultMekaList: [],
});

// ランキングを管理するc
export const ranking = reactive({
  list: [],
});
