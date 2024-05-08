<template>
  <div
    v-if="pageInfo.type === 'project'"
    class="px-4 sm:px-2 lg:px-4"
  >
    <!-- <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-base font-semibold leading-6 text-gray-900">검색 조건</h1>
      </div>
    </div>
    <div class="mt-4 sm:mt-0 sm:flex-none">
      <div>
        <div class="mt-4 flex rounded-md shadow-sm">
          <div class="relative flex flex-grow items-stretch focus-within:z-10">
            <Listbox
              v-model="selected"
              as="div"
            >
              <ListboxLabel class="sr-only"> Sort to </ListboxLabel>
              <div class="relative">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white py-2.5 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ selected.name }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronUpDownIcon
                      class="h-5 w-5 text-gray-400"
                      aria-hidden="true"
                    />
                  </span>
                </ListboxButton>

                <transition
                  leave-active-class="transition ease-in duration-100"
                  leave-from-class="opacity-100"
                  leave-to-class="opacity-0"
                >
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-max overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption
                      v-for="sort in sorts"
                      :key="sort.id"
                      v-slot="{ active, selected }"
                      as="template"
                      :value="sort"
                    >
                      <li :class="[active ? 'bg-indigo-600 text-white' : 'text-gray-900', 'relative cursor-default select-none py-2 pl-8 pr-4']">
                        <span :class="[selected ? 'font-semibold' : 'font-normal', 'block truncate']">{{ sort.name }}</span>

                        <span
                          v-if="selected"
                          :class="[active ? 'text-white' : 'text-indigo-600', 'absolute inset-y-0 left-0 flex items-center pl-1.5']"
                        >
                          <CheckIcon
                            class="h-5 w-5"
                            aria-hidden="true"
                          />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <input
              id="search"
              type="search"
              name="search"
              class="block w-full rounded-md mx-2 border-0 py-2.5 px-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              placeholder="검색어를 입력하세요."
            />
            <button
              type="button"
              class="block rounded-md bg-indigo-600 px-3 py-2.5 text-center text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              <MagnifyingGlassIcon
                class="h-5 w-5 text-white"
                aria-hidden="true"
              />
            </button>
          </div>
        </div>
      </div>
    </div> -->
    <div class="flow-root">
      <!-- <div class="mb-2 sm:flex sm:items-center">
        <div class="sm:flex-auto">
          <h1 class="text-base font-semibold leading-6 text-gray-900">검색 결과</h1>
        </div>
      </div> -->
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <table class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                <th
                  scope="col"
                  class="py-3.5 pl-4 pr-3 text-sm font-semibold text-gray-900 text-center sm:pl-3"
                >
                  고유번호
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  생성일
                </th>
                <!-- <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  수정일
                </th> -->
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  공개일
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  상태
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  참여인원수
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  구매수량/총수량
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                >
                  프로젝트 이름
                </th>
                <th
                  scope="col"
                  class="relative py-3.5 pl-3 pr-4 text-center sm:pr-3"
                >
                  관리
                </th>
              </tr>
            </thead>
            <tbody class="bg-white">
              <tr
                v-for="(list, index) in investmentList"
                :key="index"
                :class="index % 2 === 0 ? undefined : 'bg-gray-50'"
              >
                <td class="whitespace-nowrap text-center py-4 pl-4 pr-3 text-sm font-medium text-gray-700 sm:pl-3">
                  {{ list.id }}
                </td>
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ list.created_at ? DateTime.fromISO(list.created_at).toFormat("yyyy.MM.dd") : "-" }}
                </td>
                <!-- <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ "list.editDt" }}
                </td> -->
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ list.start_date ? DateTime.fromISO(list.start_date).toFormat("yyyy.MM.dd") : "-" }}
                </td>
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ list.status }}
                </td>
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ list.user_count }}
                </td>
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">{{ list.current_invest }}/{{ list.max_invest }}</td>
                <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                  {{ list.title }}
                </td>
                <td class="relative whitespace-nowrap text-center py-4 text-sm font-medium">
                  <button
                    class="rounded-md mx-1 border-0 px-2 py-1.5 ring-1 ring-inset ring-indigo-600 text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
                    @click="
                      selectProjectName = list.title;
                      selectProjectId = list.id;
                      loadUsers(list.id);
                      memberDialog = true;
                    "
                  >
                    참여인원
                  </button>
                  <button
                    disabled
                    class="rounded-md mx-1 border-0 px-2 py-1.5 ring-1 ring-inset ring-indigo-600 text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
                  >
                    정보수정
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!-- paging -->
      <div class="flex items-center justify-between mt-4 border-t border-gray-200 bg-white py-3">
        <BoardPageArea
          :pagenation="pageInfo"
          :total-count="investmentTotalCount"
        />
      </div>
    </div>

    <!-- 참여회원목록 dialog -->
    <TransitionRoot
      as="template"
      :show="memberDialog"
    >
      <Dialog
        as="div"
        class="relative z-10"
        @close="memberDialog = false"
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
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-screen-lg sm:p-6">
                <div class="mb-4 sm:flex-auto">
                  <strong class="text-lg font-semibold leading-6 text-gray-900">참여 회원 목록({{ selectProjectName }})</strong>
                </div>
                <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                  <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                    <div class="relative">
                      <div
                        v-if="selectedTr.length > 0"
                        class="absolute left-14 top-0 flex h-12 items-center space-x-3 bg-white sm:left-12"
                      >
                        <button
                          type="button"
                          class="inline-flex items-center rounded bg-white px-2 py-1 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-30 disabled:hover:bg-white"
                          @click="
                            cancelSelectUserInvestment(selectProjectId);
                            resultDialog = true;
                          "
                        >
                          선택회원 참여취소
                        </button>
                      </div>
                      <table class="min-w-full table-fixed divide-y divide-gray-300">
                        <thead>
                          <tr>
                            <th
                              scope="col"
                              class="relative px-7 sm:w-12 sm:px-6"
                            >
                              <input
                                type="checkbox"
                                class="absolute left-4 top-1/2 -mt-2 h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
                                :checked="indeterminate || (selectedTr?.length ?? 0) === (investmentUserList?.length ?? 0)"
                                :indeterminate="indeterminate"
                                @change="selectedTr = $event.target.checked ? investmentUserList.map((p) => p.id) : []"
                              />
                            </th>
                            <th
                              scope="col"
                              class="py-3.5 pr-3 text-center text-sm font-semibold text-gray-900"
                            >
                              고유번호
                            </th>
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              아이디
                            </th>
                            <!-- <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              회원구분
                            </th> -->
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              성별
                            </th>
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              이름
                            </th>
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              전화번호
                            </th>
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              구매수량
                            </th>
                            <th
                              scope="col"
                              class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                            >
                              입금여부
                            </th>
                            <th
                              scope="col"
                              class="relative text-center py-3.5 pl-3 pr-4 sm:pr-3"
                            >
                              관리
                            </th>
                          </tr>
                        </thead>
                        <tbody class="divide-y divide-gray-200 bg-white">
                          <tr
                            v-for="(list, index) in investmentUserList"
                            :key="index"
                            :class="[selectedTr.includes(list.id) && 'bg-gray-50']"
                          >
                            <td class="relative px-7 sm:w-12 sm:px-6">
                              <div
                                v-if="selectedTr.includes(list.id)"
                                class="absolute inset-y-0 left-0 w-0.5 bg-indigo-600"
                              />
                              <input
                                v-model="selectedTr"
                                type="checkbox"
                                class="absolute left-4 top-1/2 -mt-2 h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
                                :value="list.id"
                              />
                            </td>
                            <td :class="['whitespace-nowrap text-center py-4 pr-3 text-sm font-medium', selectedTr.includes(list.id) ? 'text-indigo-600' : 'text-gray-900']">
                              {{ list.id }}
                            </td>
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.user_id }}
                            </td>
                            <!-- <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ "list.type" }}
                            </td> -->
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.gender }}
                            </td>
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.name }}
                            </td>
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.phone2 }}
                            </td>
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.amount }}
                            </td>
                            <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                              {{ list.paid_status }}
                            </td>
                            <td class="relative whitespace-nowrap text-center py-4 text-sm font-medium">
                              <!-- <button
                                class="rounded-md mx-1 border-0 px-2 py-1.5 ring-1 ring-inset ring-indigo-600 text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
                                @click="cancelSingleUserInvestment(selectProjectId, list.id)"
                              >
                                참여취소
                              </button> -->
                              <button
                                :disabled="list.paid_status === 'paid'"
                                class="rounded-md mx-1 border-0 px-2 py-1.5 ring-1 ring-inset ring-indigo-600 text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
                                @click="depositSingleUserInvestment(selectProjectId, list.id)"
                              >
                                입금완료
                              </button>
                            </td>
                          </tr>
                        </tbody>
                      </table>
                    </div>
                  </div>
                </div>
                <!-- paging -->
                <div class="flex items-center justify-between mt-4 border-t border-gray-200 bg-white py-3">
                  <BoardPageArea
                    :pagenation="investmentUserPageInfo"
                    :total-count="investmentUserPageTotalCount"
                    @update-pagenation="loadPageUsers"
                  />
                </div>
                <div class="absolute top-5 right-4">
                  <button
                    type="button"
                    class="block rounded-md text-indigo-600 text-center font-semibold leading-6 shadow-sm"
                    @click="memberDialog = false"
                  >
                    <span class="sr-only">close</span>
                    <XMarkIcon
                      class="h-7 w-7 text-gray-400"
                      aria-hidden="true"
                    />
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>

    <!-- 참여취소 결과 목록 dialog -->
    <TransitionRoot
      as="template"
      :show="resultDialog"
    >
      <Dialog
        as="div"
        class="relative z-10"
        @close="resultDialog = false"
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
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-screen-lg sm:p-6">
                <div class="mb-4 sm:flex-auto">
                  <strong class="text-lg font-semibold leading-6 text-gray-900">참여취소 결과 목록</strong>
                </div>
                <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                  <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                    <table class="min-w-full divide-y divide-gray-300">
                      <thead>
                        <tr>
                          <th
                            scope="col"
                            class="py-3.5 pl-4 pr-3 text-sm font-semibold text-gray-900 text-center sm:pl-3"
                          >
                            고유번호
                          </th>
                          <th
                            scope="col"
                            class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                          >
                            아이디
                          </th>
                          <th
                            scope="col"
                            class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                          >
                            이름
                          </th>
                          <th
                            scope="col"
                            class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                          >
                            전화번호
                          </th>
                          <th
                            scope="col"
                            class="px-3 py-3.5 text-sm font-semibold text-gray-900 text-center"
                          >
                            구매수량
                          </th>
                          <th
                            scope="col"
                            class="relative py-3.5 pl-3 pr-4 text-center sm:pr-3"
                          >
                            취소성공여부
                          </th>
                        </tr>
                      </thead>
                      <tbody class="bg-white">
                        <tr
                          v-for="(list, index) in investmentCancelUserList"
                          :key="index"
                          :class="index % 2 === 0 ? undefined : 'bg-gray-50'"
                        >
                          <td class="whitespace-nowrap text-center py-4 pl-4 pr-3 text-sm font-medium text-gray-700 sm:pl-3">
                            {{ list.id }}
                          </td>
                          <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                            {{ list.user_id }}
                          </td>
                          <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                            {{ list.name }}
                          </td>
                          <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                            {{ list.phone2 }}
                          </td>
                          <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                            {{ list.amount }}
                          </td>
                          <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-700">
                            {{ list.cancel_status }}
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
                <!-- paging -->
                <!-- <div
                  class="flex items-center justify-between mt-4 border-t border-gray-200 bg-white py-3"
                >
                  <BoardPageArea :pagenation="pageInfo" />
                </div>
                <div class="absolute top-5 right-4">
                  <button
                    type="button"
                    class="block rounded-md text-indigo-600 text-center font-semibold leading-6 shadow-sm"
                    @click="resultDialog = false"
                  >
                    <span class="sr-only">close</span>
                    <XMarkIcon
                      class="h-7 w-7 text-gray-400"
                      aria-hidden="true"
                    />
                  </button>
                </div> -->
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from "vue";
import { DateTime } from "luxon";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { CheckIcon, ChevronUpDownIcon, MagnifyingGlassIcon } from "@heroicons/vue/20/solid";
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot, Listbox, ListboxButton, ListboxLabel, ListboxOption, ListboxOptions } from "@headlessui/vue";
import { getInvestments, getInvestmentCount, getUsersForInvestmentList, getUsersForInvestmentListCount, cancelUserInvestment, depositUserInvestment } from "~/composables/apis/investmentGroup";
import { CancelUserInvestmentRequest, Investment, User, Pagenation } from "~/composables/models/models";

const investmentTotalCount = ref<number>(0);
const investmentUserPageTotalCount = ref<number>(0);

const route = useRoute();

const memberDialog = ref(false);
const resultDialog = ref(false);

const selectedTr = ref<number[]>([]);
const checked = ref(false);
const indeterminate = computed(() => selectedTr.value.length > 0 && selectedTr.value.length < investmentUserList.value.length);

const sorts = [
  { id: 1, name: "고유번호" },
  { id: 2, name: "프로젝트 이름" },
];

const selected = ref(sorts[1]);

const pageInfo = ref<Pagenation>({
  page: 1,
  limit: 10,
  type: "project",
  search: "title",
  isCallBack: false,
});
const investmentUserPageInfo = ref<Pagenation>({
  page: 1,
  limit: 10,
  type: "projectUser",
  search: "title",
  isCallBack: true,
});

const investmentList = ref<Investment[]>([]);
const investmentUserList = ref<User[]>([]);
const investmentCancelUserList = ref<User[]>([]);
const selectProjectName = ref<string>("");
const selectProjectId = ref<number>(-1);

watch(
  () => route.query,
  async () => {
    pageInfo.value = {
      page: Number(route.query.page) || 1,
      limit: Number(route.query.limit) || 10,
      type: (route.query.type as string) || "project",
      search: (route.query.search as string) || "title",
      keyword: (route.query.keyword as string) || "",
    };
    console.log(pageInfo.value);
    await loadPage();
  }
);

onMounted(async () => {
  pageInfo.value = {
    page: Number(route.query.page) || 1,
    limit: Number(route.query.limit) || 10,
    type: (route.query.type as string) || "project",
    search: (route.query.search as string) || "title",
    keyword: (route.query.keyword as string) || "",
  };
  await loadPage();
});

async function loadPage() {
  if (!pageInfo.value || pageInfo.value.type !== "project") return;
  investmentTotalCount.value = await getInvestmentCount();
  const result = await getInvestments(pageInfo.value.page, pageInfo.value.limit);
  investmentList.value = result.data.investments;
}

async function loadUsers(investmentId: number) {
  selectedTr.value = [];
  investmentUserPageInfo.value.id = investmentId;
  investmentUserPageInfo.value.page = 1;
  investmentUserPageInfo.value.limit = 10;
  const result = await getUsersForInvestmentList(investmentId, investmentUserPageInfo.value.page, investmentUserPageInfo.value.limit);
  investmentUserPageTotalCount.value = await getUsersForInvestmentListCount(investmentId);
  investmentUserList.value = result.data.users;
}

async function loadPageUsers(page: Pagenation) {
  console.log(page);
  investmentUserPageInfo.value = page;
  if (!investmentUserPageInfo.value.id) return;
  selectedTr.value = [];
  const result = await getUsersForInvestmentList(investmentUserPageInfo.value.id, investmentUserPageInfo.value.page, investmentUserPageInfo.value.limit);
  investmentUserList.value = result.data.users;
}

async function cancelSingleUserInvestment(investmentId: number, userId: number) {
  investmentCancelUserList.value = [];
  const result = await cancelUserInvestment(investmentId, {
    user_ids: [
      {
        id: userId,
      },
    ],
  } as CancelUserInvestmentRequest);
  investmentCancelUserList.value = result.data.users;

  const successList = investmentCancelUserList.value.filter((user) => user.cancel_status === "success");
  if (successList.length > 0) {
    investmentUserList.value = investmentUserList.value.filter((user) => !successList.map((s) => s.id).includes(user.id));
  } else {
    alert("취소가 불가능한 회원입니다.");
  }
}

async function depositSingleUserInvestment(investmentId: number, userId: number) {
  investmentCancelUserList.value = [];
  const result = await depositUserInvestment(investmentId, {
    user_ids: [
      {
        id: userId,
      },
    ],
  } as CancelUserInvestmentRequest);
  if (result.success) {
    /// update paid status
    const user = investmentUserList.value.find((u) => u.id === userId);
    if (user) {
      user.paid_status = "paid";
    }
  }
}

async function cancelSelectUserInvestment(investmentId: number) {
  investmentCancelUserList.value = [];
  const result = await cancelUserInvestment(investmentId, {
    user_ids: selectedTr.value.map((id) => {
      return {
        id,
      };
    }),
  } as CancelUserInvestmentRequest);
  investmentCancelUserList.value = result.data.users;

  const successList = investmentCancelUserList.value.filter((user) => user.cancel_status === "success");
  if (successList.length > 0) {
    investmentUserList.value = investmentUserList.value.filter((user) => !successList.map((s) => s.id).includes(user.id));
  }
}

async function depositSelectUserInvestment(investmentId: number) {
  investmentCancelUserList.value = [];
  const result = await depositUserInvestment(investmentId, {
    user_ids: selectedTr.value.map((id) => {
      return {
        id,
      };
    }),
  } as CancelUserInvestmentRequest);
}
</script>
