import axios, { AxiosResponse } from "axios";
import { GetInvestmentsResponse, UsersResponse } from "~/composables/models/models";
import { useUserStore } from "~/store/users";
import { getBaseURL } from "../utils";

export async function getUsers(page: number, limit: number): Promise<UsersResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/users?page=${page}&limit=${limit}`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data as UsersResponse;
}

export async function getUserCount(): Promise<number> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/users/count`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data.data as number;
}

export async function getInvestmentsForUserList(userId: number, page: number, limit: number): Promise<GetInvestmentsResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/users/${userId}/investments?page=${page}&limit=${limit}`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data as GetInvestmentsResponse;
}

export async function getInvestmentsForUserListCount(userId: number): Promise<number> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/users/${userId}/investments/count`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data.data as number;
}
