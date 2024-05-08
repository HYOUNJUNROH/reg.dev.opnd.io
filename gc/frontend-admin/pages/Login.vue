<template>
  <div class="flex h-screen min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <img
        class="mx-auto h-10 w-auto"
        src="/img/logo.svg"
        alt="GEOPOP"
        width="229"
      />
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <div>
        <label
          for="email"
          class="block text-sm font-medium leading-6 text-gray-900"
          >ID</label
        >
        <div class="mt-2">
          <input
            id="email"
            v-model="email"
            name="email"
            type="email"
            autocomplete="email"
            class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            @keyup.enter="goLogin"
          />
        </div>
      </div>

      <div class="mt-4">
        <div class="flex items-center justify-between">
          <label
            for="password"
            class="block text-sm font-medium leading-6 text-gray-900"
            >Password</label
          >
        </div>
        <div class="mt-2">
          <input
            id="password"
            v-model="password"
            name="password"
            type="password"
            autocomplete="current-password"
            class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            @keyup.enter="goLogin"
          />
        </div>
      </div>
      <div class="mt-6">
        <button
          type="submit"
          class="flex w-full justify-center rounded-md bg-[#023e7d] px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
          @click="goLogin"
        >
          Sign in
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { login } from "@/composables/apis/authGroup";

const router = useRouter();

const email = ref<string>("");
const password = ref<string>("");

console.log("API_BASE_URL ", useRuntimeConfig().API_BASE_URL ?? "");

async function goLogin() {
  const result = await login(email.value, password.value);

  if (result) {
    // alert("로그인 성공");
    router.push("/");
  } else {
    alert("로그인 실패");
  }
}

definePageMeta({
  layout: "login",
});
</script>
<style lang="scss"></style>
