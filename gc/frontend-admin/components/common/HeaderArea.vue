<template>
  <header>
    <div class="mx-auto flex w-full max-w-7xl items-center justify-between py-8 px-4 sm:px-6 lg:px-8">
      <div class="flex flex-1 items-center gap-x-6">
        <img
          src="/img/logo.svg"
          alt="GEOPOP"
          width="200"
        />
      </div>
      <div class="flex flex-1 items-center justify-end gap-x-8">
        <div class="flex -m-2.5 p-2.5 text-gray-400 hover:text-gray-500">
          <span class="sr-only">Your profile</span>
          <UserIcon
            class="h-6 w-6"
            aria-hidden="true"
          />
          <span class="pl-2">{{ user.userInfo.name }}</span>
        </div>
        <button
          type="button"
          class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500"
          @click="
            oldPassword = '';
            newPassword = '';
            newPasswordCheck = '';
            pwDialog = true;
          "
        >
          <span class="sr-only">비밀번호 변경</span>
          <KeyIcon
            class="h-6 w-6"
            aria-hidden="true"
          />
        </button>
        <button
          type="button"
          class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500"
          @click="logout()"
        >
          <span class="sr-only">로그아웃</span>
          <ArrowRightOnRectangleIcon
            class="h-6 w-6"
            aria-hidden="true"
          />
        </button>
      </div>
    </div>

    <TransitionRoot
      as="template"
      :show="pwDialog"
    >
      <Dialog
        as="div"
        class="relative z-10"
        @close="pwDialog = false"
      >
        <TransitionChild
          as="template"
          enter="ease-out duration-300"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="ease-in duration-200"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </TransitionChild>

        <div class="fixed inset-0 z-10 overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild
              as="template"
              enter="ease-out duration-300"
              enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enter-to="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leave-from="opacity-100 translate-y-0 sm:scale-100"
              leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-md sm:p-6">
                <div>
                  <div class="mt-2 text-center sm:mt-2">
                    <DialogTitle
                      as="h3"
                      class="text-xl font-semibold leading-6 text-gray-900"
                    >
                      비밀번호 변경
                    </DialogTitle>
                    <div class="mt-6 space-y-8 border-b border-gray-900/10 pb-12 sm:space-y-0 sm:divide-gray-900/10 sm:border-t sm:pb-0">
                      <div class="sm:grid sm:grid-cols-3 border-0 sm:items-start sm:gap-4 sm:py-3">
                        <label
                          for="pw1"
                          class="block pl-6 text-left text-sm font-medium leading-6 text-gray-900 sm:pt-1.5"
                          >기존 비밀번호</label
                        >
                        <div class="mt-2 sm:col-span-2 sm:mt-0">
                          <div class="flex overflow-hidden rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-1 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                            <input
                              id="pw1"
                              v-model="oldPassword"
                              type="text"
                              name="password"
                              autocomplete="password"
                              class="block flex-1 border-0 bg-transparent py-1.5 px-1.5 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                              placeholder="기존 비밀번호를 입력하세요."
                            />
                          </div>
                        </div>
                      </div>
                      <div class="sm:grid sm:grid-cols-3 border-0 sm:items-start sm:gap-4 sm:py-3">
                        <label
                          for="pw1"
                          class="block pl-6 text-left text-sm font-medium leading-6 text-gray-900 sm:pt-1.5"
                          >새 비밀번호</label
                        >
                        <div class="mt-2 sm:col-span-2 sm:mt-0">
                          <div class="flex overflow-hidden rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-1 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                            <input
                              id="pw1"
                              v-model="newPassword"
                              type="text"
                              name="password"
                              autocomplete="password"
                              class="block flex-1 border-0 bg-transparent py-1.5 px-1.5 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                              placeholder="새 비밀번호를 입력하세요."
                            />
                          </div>
                        </div>
                      </div>
                      <div class="sm:grid sm:grid-cols-3 border-0 sm:items-start sm:gap-4 sm:py-3">
                        <label
                          for="pw1"
                          class="block pl-6 text-left text-sm font-medium leading-6 text-gray-900 sm:pt-1.5"
                          >새 비밀번호 확인</label
                        >
                        <div class="mt-2 sm:col-span-2 sm:mt-0">
                          <div class="flex overflow-hidden rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-1 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                            <input
                              id="pw1"
                              v-model="newPasswordCheck"
                              type="text"
                              name="password"
                              autocomplete="password"
                              class="block flex-1 border-0 bg-transparent py-1.5 px-1.5 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                              placeholder="새 비밀번호를 다시 입력하세요."
                            />
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="mt-5 sm:mt-6 sm:grid sm:grid-flow-row-dense sm:grid-cols-2 sm:gap-3">
                    <button
                      type="button"
                      class="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 sm:col-start-2"
                      @click="changePassword()"
                    >
                      변경
                    </button>
                    <button
                      ref="cancelButtonRef"
                      type="button"
                      class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:col-start-1 sm:mt-0"
                      @click="pwDialog = false"
                    >
                      취소
                    </button>
                  </div>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </header>
</template>

<script setup>
import { ref } from "vue";
import { UserIcon, KeyIcon, ArrowRightOnRectangleIcon } from "@heroicons/vue/24/outline";
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from "@headlessui/vue";
import { updatePassword } from "~/composables/apis/authGroup";
import { useUserStore } from "../../store/users";

const user = useUserStore();

const pwDialog = ref(false);

const oldPassword = ref("");
const newPassword = ref("");
const newPasswordCheck = ref("");

async function changePassword() {
  /// check newPwd
  if (newPassword.value !== newPasswordCheck.value) {
    alert("새 비밀번호가 일치하지 않습니다.");
    return;
  }

  const result = await updatePassword(oldPassword.value, newPassword.value);
  if (result) {
    alert("비밀번호가 변경되었습니다.");
    pwDialog.value = false;
  } else {
    alert("비밀번호 변경에 실패하였습니다.");
  }
}

async function logout() {
  user.logout();
  window.location.href = "/Login";
}

// watchEffect(() => {
//   if (!user.value) {
//     window.location.href = "/Login";
//   }
// });
</script>
