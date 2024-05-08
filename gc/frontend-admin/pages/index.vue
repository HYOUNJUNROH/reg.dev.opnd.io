<template>
  <div class="mx-auto w-full max-w-7xl pb-8 px-4 sm:px-6 lg:pb-8">
    <div class="px-4">
      <div class="sm:hidden">
        <label for="tabs" class="sr-only">Select a tab</label>
        <!-- Use an "onChange" listener to redirect the user to the selected tab URL. -->
        <select
          id="tabs"
          v-model="currentTab"
          name="tabs"
          class="block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
        >
          <option :value="'member'">회원관리</option>
          <option :value="'project'">프로젝트관리</option>
          <option :value="'banner'">배너관리</option>
        </select>
      </div>
      <div class="hidden sm:block">
        <div class="border-b border-gray-200">
          <nav class="-mb-px flex space-x-8" aria-label="Tabs">
            <button
              :class="[
                currentTab === 'member'
                  ? 'border-indigo-500 text-indigo-600'
                  : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700',
                'whitespace-nowrap border-b-2 py-4 px-1 text-sm font-medium',
              ]"
              @click="currentTab = 'member'"
            >
              회원관리
            </button>
            <button
              :class="[
                currentTab === 'project'
                  ? 'border-indigo-500 text-indigo-600'
                  : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700',
                'whitespace-nowrap border-b-2 py-4 px-1 text-sm font-medium',
              ]"
              @click="currentTab = 'project'"
            >
              프로젝트관리
            </button>
            <button
              :class="[
                currentTab === 'banner'
                  ? 'border-indigo-500 text-indigo-600'
                  : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700',
                'whitespace-nowrap border-b-2 py-4 px-1 text-sm font-medium',
              ]"
              @click="currentTab = 'banner'"
            >
              배너관리
            </button>
          </nav>
        </div>
      </div>
    </div>
    <div class="py-4">
      <ManagementMemberPanel id="ManagementMemberPanel" />
      <ManagementProjectPanel id="ManagementProjectPanel" />
      <ManagementBannerPanel id="ManagementBannerPanel" />
    </div>
  </div>
</template>

<script lang="ts" setup>
const router = useRouter();
const route = useRoute();

const currentTab = ref("member");

watch(
  () => route.query,
  () => {
    currentTab.value = route.query.type as string;
  }
);

watch(
  () => currentTab.value,
  () => {
    pushPage();
  }
);

onMounted(() => {
  if (route.query.type) {
    currentTab.value = route.query.type as string;
  }
  router.push({
    path: route.path,
    query: {
      page: 1,
      limit: 10,
      type: currentTab.value,
    },
  });
});

function pushPage() {
  router.push({
    path: route.path,
    query: {
      page: 1,
      limit: 10,
      type: currentTab.value,
    },
  });
}

definePageMeta({
  layout: "index",
  middleware: ["auth"],
});
</script>
