<template>
  <div class="relative hidden sm:flex sm:flex-1 sm:items-center sm:justify-center">
    <div class="inline-flex items-center">
      <nav
        class="isolate inline-flex -space-x-px rounded-md shadow-sm"
        aria-label="Pagination"
      >
        <a
          v-if="lpage > 1"
          role="button"
          class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
          @click="
            if (Number(pageInfo.page) - 1 > 0) {
              pageInfo.page = Number(pageInfo.page) - 1;
              pushPage();
            }
          "
        >
          <span class="sr-only">Previous</span>
          <ChevronLeftIcon
            class="h-5 w-5"
            aria-hidden="true"
          />
        </a>
        <a
          v-for="(i, idx) in rpage - lpage + 1"
          :key="idx"
          role="button"
          aria-current="page"
          class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
          :class="{
            'bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-700 hover:bg-indigo-600': i + lpage - 1 === Number(pageInfo.page),
          }"
          @click="
            {
              pageInfo.page = i + lpage - 1;
              pushPage();
            }
          "
          >{{ i + lpage - 1 }}</a
        >
        <a
          v-if="rpage < maxPage"
          role="button"
          class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
          @click="
            if (Number(pageInfo.page) + 1 <= maxPage) {
              pageInfo.page = Number(pageInfo.page) + 1;
              pushPage();
            }
          "
          ><span class="sr-only">Next</span>
          <ChevronRightIcon
            class="h-5 w-5"
            aria-hidden="true"
        /></a>
      </nav>
      <div class="absolute right-0 inline-flex ml-2">
        <select
          v-model="pageInfo.limit"
          class="block rounded-md px-6 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 h-9"
          @change="
            updatePage();
            pushPage();
          "
        >
          <option :value="10">10개씩 보기</option>
          <option :value="30">30개씩 보기</option>
          <option :value="50">50개씩 보기</option>
        </select>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { PropType } from "vue";
import { ref } from "vue";
import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/vue/20/solid";
import { Pagenation } from "~/composables/models/models";
// import { Pagenation } from "@/composables/util";

const route = useRoute();
const router = useRouter();

const props = defineProps({
  pagenation: {
    type: Object as PropType<Pagenation>,
    required: true,
  },
  totalCount: {
    type: Number,
    required: true,
  },
});

const pageInfo = ref(props.pagenation);
const emits = defineEmits(["update-pagenation"]);

const maxPage = computed(() => {
  return Math.max(Math.ceil(props.totalCount / pageInfo.value.limit!) || 0, 1);
});
const rpage = computed(() => {
  return Math.min(Number(pageInfo.value.page ?? 1) + 4, maxPage.value);
});

const lpage = computed(() => {
  return Math.max(Number(pageInfo.value.page ?? 1) - 4, 1);
});

function updatePage() {
  pageInfo.value.page = Math.min(Math.ceil(props.totalCount / pageInfo.value.limit!) || 1, pageInfo.value.page);
}

function pushPage() {
  console.log("pushPage", pageInfo.value);
  if (!pageInfo.value.isCallBack || pageInfo.value.isCallBack === false)
    router.push({
      path: route.path,
      query: pageInfo.value,
    });
  else emits("update-pagenation", pageInfo.value);
}
</script>

<style lang="scss"></style>
