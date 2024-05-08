import Toast from "vue-toastification";
import { defineNuxtPlugin } from "#app";
import "vue-toastification/dist/index.css"; // if needed

const options = {
  position: "top-right", // 포지션
  timeout: 5000, // 시간
  pauseOnFocusLoss: false, // 다른 곳 클릭 시 시간 멈춤
  pauseOnHover: false, // 모달 마우스 호버 시 시간 멈춤
  closeOnClick: true, // 클릭 시 모달 닫기
  draggable: true, // 드레그 해서 모달 닫기
  draggablePercent: 0.55, // 드래그 해야 하는 퍼센트 (0.55% 지나면 사라짐)
  closeButton: "button", // 닫기 버튼 (false시 닫기버튼 사라짐)
  showCloseButtonOnHover: false, // 마우스 호버 시 닫기 버튼 생성
  rtl: false, // 왼쪽정렬 (true시 오른쪽 정렬)
  icon: true, // 텍스트 옆 아이콘 유무
};

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(Toast, options);
});
