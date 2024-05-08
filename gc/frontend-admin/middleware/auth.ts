import { useUserStore } from "~/store/users";

export default defineNuxtRouteMiddleware(() => {
  const user = useUserStore();
  watchEffect(() => {
    if (!user.userInfo?.id) {
      window.location.href = "/login";
    }
  });
});
