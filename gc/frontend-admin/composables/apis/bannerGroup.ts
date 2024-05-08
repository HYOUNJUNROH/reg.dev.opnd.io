import axios, { AxiosResponse } from "axios";
import { Banner, BannerPriorities, Banners, GetInvestmentsResponse, UsersResponse } from "~/composables/models/models";
import { useUserStore } from "~/store/users";
import { getBaseURL } from "../utils";

// bannerGroup.GET("", adm.GetBanners)
export async function getBanners(): Promise<Banner[]> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/banners`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data.data.banners as Banner[];
}

// bannerGroup.POST("", adm.PostBanner)
export async function postBanner(formData: FormData): Promise<AxiosResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post(`/adm/banners`, formData).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response;
}

// bannerGroup.POST("/priority", adm.PostPriorityBanner)
export async function postPriorityBanner(infos: BannerPriorities): Promise<AxiosResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post(`/adm/banners/priority`, infos).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response;
}

// bannerGroup.DELETE("/:id", adm.DeleteBanner)
export async function deleteBanner(id: number): Promise<AxiosResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const formData = new FormData();
  formData.append("id", id.toString());
  const response = await instance
    .delete(`/adm/banners/${id}`, {
      data: formData,
    })
    .catch((error) => {
      if (error.response.status === 401) {
        user.logout();
      }
    });
  return response;
}
