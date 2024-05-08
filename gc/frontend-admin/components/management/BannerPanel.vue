<template>
  <div
    v-if="pageInfo.type === 'banner'"
    class="px-4 sm:px-2 lg:px-4"
  >
    <div class="flex flex-wrap mt-4 sm:px-4 sm:items-center">
      <div class="inline-flex sm:flex-auto sm:items-center">
        <button
          type="button"
          class="inline-flex rounded-md ring-1 ring-inset ring-indigo-600 px-3 py-2 text-center text-sm font-semibold shadow-sm text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
          @click="changePriority('up')"
        >
          <ChevronUpIcon
            class="h-5 w-5 mr-1"
            aria-hidden="true"
          />
          올리기
        </button>
        <button
          type="button"
          class="inline-flex ml-2 rounded-md ring-1 ring-inset ring-indigo-600 px-3 py-2 text-center text-sm font-semibold shadow-sm text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
          @click="changePriority('down')"
        >
          <ChevronDownIcon
            class="h-5 w-5 mr-1"
            aria-hidden="true"
          />
          내리기
        </button>
      </div>
      <div class="ml-auto sm:mt-0 sm:flex-none">
        <button
          type="button"
          class="inline-flex rounded-md ring-1 ring-inset ring-indigo-600 px-3 py-2 text-center text-sm font-semibold shadow-sm text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
          @click="deleteDialog = true"
        >
          배너 삭제
        </button>
      </div>
    </div>
    <div class="mt-2 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <div class="relative">
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
                      :checked="indeterminate || (selectedTr?.length ?? 0) === (bannerList?.length ?? 0)"
                      :indeterminate="indeterminate"
                      @change="selectedTr = $event.target.checked ? bannerList.map((p) => p.id) : []"
                    />
                  </th>
                  <th
                    scope="col"
                    class="min-w-[4rem] py-3.5 pr-3 text-center text-sm font-semibold text-gray-900"
                  >
                    고유번호
                  </th>
                  <th
                    scope="col"
                    class="min-w-[4rem] px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                  >
                    생성일
                  </th>
                  <th
                    scope="col"
                    class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                  >
                    수정일
                  </th>
                  <th
                    scope="col"
                    class="min-w-[12rem] px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                  >
                    배너이미지
                  </th>
                  <th
                    scope="col"
                    class="min-w-[6rem] px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                  >
                    이미지 설명
                  </th>
                  <th
                    scope="col"
                    class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                  >
                    관리
                  </th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 bg-white">
                <tr
                  v-for="(banner, index) in bannerList"
                  :key="index"
                  :class="[selectedTr.includes(banner.id) && 'bg-gray-50']"
                >
                  <td class="relative px-7 sm:w-12 sm:px-6">
                    <div
                      v-if="selectedTr.includes(banner.id)"
                      class="absolute inset-y-0 left-0 w-0.5 bg-indigo-600"
                    />
                    <input
                      v-model="selectedTr"
                      type="checkbox"
                      class="absolute left-4 top-1/2 -mt-2 h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
                      :value="banner.id"
                    />
                  </td>
                  <td :class="['whitespace-nowrap text-center py-4 pr-3 text-sm font-medium', selectedTr.includes(banner.id) ? 'text-indigo-600' : 'text-gray-900']">
                    {{ banner.id }}
                  </td>
                  <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                    {{ banner.created_at ? DateTime.fromISO(banner.created_at).toFormat("yyyy.MM.dd") : "-" }}
                  </td>
                  <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                    {{ banner.updated_at ? DateTime.fromISO(banner.updated_at).toFormat("yyyy.MM.dd") : "-" }}
                  </td>
                  <td class="whitespace-nowrap text-center px-1 py-4 text-sm text-gray-500">
                    <div class="flex h-10 w-50">
                      <img
                        class="m-auto flex-shrink-0 h-10 w-50 object-cover"
                        :src="getImage(banner.image)"
                        alt=""
                      />
                    </div>
                  </td>
                  <td class="max-w-400 px-3 py-4 text-sm text-gray-500">
                    {{ banner.description }}
                  </td>
                  <td class="whitespace-nowrap text-center px-3 py-4 text-sm text-gray-500">
                    <button
                      class="rounded-md mx-1 border-0 px-2 py-1.5 ring-1 ring-inset ring-indigo-600 text-indigo-600 hover:text-indigo-900 hover:ring-1 hover:ring-inset disabled:text-gray-300 disabled:ring-gray-300"
                      @click="
                        edit = true;
                        selectBanner = banner;
                        loadBannerToEdit();
                        dialogOpen = true;
                      "
                    >
                      수정
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    <div class="flex justify-center pt-4 border-t border-gray-200 px-4 sm:px-0">
      <button
        type="button"
        class="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        @click="
          edit = false;
          dialogOpen = true;
        "
      >
        + 배너 항목 추가 (최대 10개)
      </button>
    </div>

    <!-- 추가 & 수정 dialog -->
    <TransitionRoot
      as="template"
      :show="dialogOpen"
    >
      <Dialog
        as="div"
        class="relative z-10"
        @close="closeDialog()"
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
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-4/5 sm:max-w-screen-lg sm:p-6">
                <div class="absolute top-5 right-4">
                  <button
                    type="button"
                    class="block rounded-md text-indigo-600 text-center font-semibold leading-6 shadow-sm"
                    @click="closeDialog()"
                  >
                    <span class="sr-only">close</span>
                    <XMarkIcon
                      class="h-7 w-7 text-gray-400"
                      aria-hidden="true"
                    />
                  </button>
                </div>
                <div
                  v-if="edit"
                  class="mb-4 sm:flex-auto items-center"
                >
                  <strong class="text-lg font-semibold leading-6 text-gray-900">배너 항목 수정</strong>
                  <span>({{ selectBanner?.id }})</span>
                </div>
                <div v-else>
                  <strong class="text-lg font-semibold leading-6 text-gray-900">배너 항목 추가</strong>
                </div>

                <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                  <div class="col-span-full">
                    <label
                      for="cover-photo"
                      class="block text-sm font-medium leading-6 text-gray-900"
                      >배너이미지</label
                    >
                    <div
                      class="relative mt-2 w-full rounded-lg border border-dashed border-gray-900/25"
                      style="padding-top: 20%"
                    >
                      <div class="flex justify-center items-center absolute top-0 right-0 left-0 bottom-0 text-center">
                        <!-- img 태그는 경로가 있을 경우에만 표출 -->
                        <slot v-if="addImageSrc">
                          <img
                            class="w-full h-full object-cover"
                            :src="addImageSrc"
                            alt=""
                          />
                        </slot>
                        <div class="upload-text flex flex-col justify-center items-center">
                          <PhotoIcon
                            class="mx-auto h-12 w-12 text-gray-300"
                            aria-hidden="true"
                          />
                          <div class="mt-2 flex text-sm leading-6 text-gray-600">
                            <label
                              for="file-upload"
                              class="cursor-pointer rounded-md bg-white font-semibold text-indigo-600 focus-within:outline-none focus-within:ring-offset-0 hover:text-indigo-500 after:content-[''] after:absolute after:top-0 after:bottom-0 after:left-0 after:right-0"
                            >
                              <span>Upload a file</span>
                              <input
                                id="file-upload"
                                ref="fileUploadRef"
                                name="file-upload"
                                type="file"
                                class="sr-only"
                                @change="fileUpload"
                              />
                            </label>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="col-span-full">
                    <label
                      for="img-alterate"
                      class="block text-sm font-medium leading-6 text-gray-900"
                      >이미지설명</label
                    >
                    <div class="mt-2">
                      <input
                        id="img-alterate"
                        v-model="addBannerImageDetail"
                        type="text"
                        name="img-alterate"
                        autocomplete="img-alterate"
                        placeholder="이미지에 대한 설명을 입력하세요."
                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      />
                    </div>
                  </div>
                </div>
                <div class="flex mt-5 sm:mt-8 items-center justify-center text-center">
                  <slot v-if="edit">
                    <button
                      type="button"
                      class="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 sm:ml-3 sm:w-auto"
                      @click="editBanner()"
                    >
                      수정
                    </button>
                  </slot>
                  <slot v-else>
                    <button
                      type="button"
                      class="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 sm:ml-3 sm:w-auto"
                      @click="addBanner()"
                    >
                      추가
                    </button>
                  </slot>
                  <button
                    type="button"
                    class="inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:w-auto sm:ml-2"
                    @click="closeDialog()"
                  >
                    취소
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>

    <!-- 삭제 dialog -->
    <TransitionRoot
      as="template"
      :show="deleteDialog"
    >
      <Dialog
        as="div"
        class="relative z-10"
        @close="deleteDialog = false"
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
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-3/4 sm:max-w-lg">
                <div class="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                  <div class="sm:flex sm:items-center">
                    <div class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                      <ExclamationTriangleIcon
                        class="h-6 w-6 text-red-600"
                        aria-hidden="true"
                      />
                    </div>
                    <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                      <DialogTitle
                        as="h3"
                        class="text-base font-semibold leading-6 text-gray-900"
                      >
                        배너 항목 삭제({{ selectedTr.length ?? 0 }})
                      </DialogTitle>
                    </div>
                  </div>
                </div>
                <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <button
                    type="button"
                    class="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto"
                    @click="deleteSelectBanners()"
                  >
                    삭제
                  </button>
                  <button
                    ref="cancelButtonRef"
                    type="button"
                    class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
                    @click="deleteDialog = false"
                  >
                    취소
                  </button>
                </div>
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
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from "@headlessui/vue";
import { getUserCount } from "~/composables/apis/userGroup";
import { Banner, BannerPriorities, Pagenation } from "~/composables/models/models";

import { XMarkIcon, ExclamationTriangleIcon, ArrowDownTrayIcon, ChevronUpIcon, ChevronDownIcon, PhotoIcon } from "@heroicons/vue/24/outline";
import { deleteBanner, getBanners, postBanner, postPriorityBanner } from "~/composables/apis/bannerGroup";
import { getImage } from "../../composables/utils";
import urlJoin from "url-join";

const route = useRoute();
const router = useRouter();

const dialogOpen = ref(false);
const deleteDialog = ref(false);

const selectedTr = ref<number[]>([]);

const edit = ref(false);

const selectBanner = ref<Banner | null>(null);

const fileUploadRef = ref<HTMLInputElement | null>(null);
const addImage = ref<File | null>(null);
const addImageSrc = ref<string | null>(null);
const addBannerImageDetail = ref<string>("");

const indeterminate = computed(() => selectedTr.value.length > 0 && selectedTr.value.length < bannerList.value.length);

const pageInfo = ref<Pagenation>({
  page: 1,
  limit: 10,
  type: "banner",
  search: "title",
  isCallBack: false,
});

const bannerList = ref<Banner[]>([] as Banner[]);
watch(
  () => route.query,
  async () => {
    pageInfo.value = {
      page: Number(route.query.page) || 1,
      limit: Number(route.query.limit) || 10,
      type: (route.query.type as string) || "",
      search: (route.query.search as string) || "title",
      keyword: (route.query.keyword as string) || "",
    };
    await loadPage();
  }
);

onMounted(async () => {
  pageInfo.value = {
    page: Number(route.query.page) || 1,
    limit: Number(route.query.limit) || 10,
    type: (route.query.type as string) || "banner",
    search: (route.query.search as string) || "title",
    keyword: (route.query.keyword as string) || "",
  };
  await loadPage();
});

async function loadPage() {
  if (!pageInfo.value || pageInfo.value.type !== "banner") return;
  const result = await getBanners();
  result.sort((a, b) => a.priority - b.priority);
  bannerList.value = result;
  console.log(result);
}

async function addBanner() {
  if (!pageInfo.value || pageInfo.value.type !== "banner") return;

  if (!addImage.value) {
    alert("이미지를 선택해주세요.");
    return;
  }

  const formData = new FormData();
  formData.append("image", addImage.value);
  formData.append("created_at", DateTime.now().toISO());
  formData.append("updated_at", DateTime.now().toISO());
  formData.append("description", addBannerImageDetail.value);
  formData.append("priority", ((bannerList.value?.length ?? 0) + 1).toString());

  const res = await postBanner(formData);
  if (res) {
    closeDialog();
    await loadPage();
  }
}

async function editBanner() {
  if (!pageInfo.value || pageInfo.value.type !== "banner") return;

  // if (!addImage.value) return;

  const formData = new FormData();
  formData.append("id", selectBanner.value?.id.toString() ?? "");
  if (addImage.value) formData.append("image", addImage.value);
  else formData.append("image", selectBanner.value?.image ?? "");
  formData.append("created_at", selectBanner.value?.created_at ?? "");
  formData.append("updated_at", DateTime.now().toISO());
  formData.append("description", addBannerImageDetail.value);
  formData.append("priority", selectBanner.value?.priority.toString() ?? "");

  const res = await postBanner(formData);
  if (res) {
    closeDialog();
    await loadPage();
  }
}

function fileUpload() {
  if (!fileUploadRef.value) return;
  const file = fileUploadRef.value.files?.[0];
  if (!file) return;
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = () => {
    console.log(reader.result);
    addImageSrc.value = reader.result as string;
    addImage.value = file;
  };
}

function getFileSrc(imageName: string) {
  const base = getBaseURL();
  return urlJoin(String(base), "adm/banners/usr/src/app/uploads", imageName);
}

async function changePriority(arrow: "up" | "down") {
  if (!pageInfo.value || pageInfo.value.type !== "banner") return;

  if (selectedTr.value.length !== 1) {
    alert("하나의 배너를 선택해주세요.");
    return;
  }

  const selectedBanner = bannerList.value.find((banner) => banner.id === selectedTr.value[0]);
  if (!selectedBanner) return;

  const selectedPriority = selectedBanner.priority;

  const targetBanner = bannerList.value.find((banner) => banner.priority === (arrow === "up" ? selectedPriority - 1 : selectedPriority + 1));
  if (!targetBanner) return;

  const infos: BannerPriorities = {
    banner_priorities: [
      {
        id: selectedBanner.id,
        priority: targetBanner.priority,
      },
      {
        id: targetBanner.id,
        priority: selectedBanner.priority,
      },
    ],
  };

  const res = await postPriorityBanner(infos);
  if (res) {
    await loadPage();
  }
}

async function deleteSelectBanners() {
  if (!pageInfo.value || pageInfo.value.type !== "banner") return;

  if (selectedTr.value.length === 0) {
    alert("삭제할 배너를 선택해주세요.");
    return;
  }

  const ids = selectedTr.value;
  // const res = await deleteBanner(ids);
  await Promise.all(ids.map((id) => deleteBanner(id)))
    .then(async (res) => {
      console.log(res);
      if (res) {
        await loadPage();
      }
    })
    .finally(() => {
      deleteDialog.value = false;
    });
}

function closeDialog() {
  /// init
  addImage.value = null;
  addImageSrc.value = null;
  addBannerImageDetail.value = "";

  dialogOpen.value = false;
}

async function loadBannerToEdit() {
  if (!selectBanner.value) return;
  addImageSrc.value = getImage(selectBanner.value.image);
  addBannerImageDetail.value = selectBanner.value.description;
}
</script>
<style lang="scss">
img + .upload-text {
  svg {
    display: none;
  }
  label {
    font-size: 0;
  }
}
</style>
