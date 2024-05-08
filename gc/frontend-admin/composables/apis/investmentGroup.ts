import axios, { AxiosResponse } from "axios";
import { useUserStore } from "~/store/users";
import { GetInvestmentsResponse, UsersResponse, CancelUserInvestmentRequest } from "../models/models";

export async function getInvestments(page: number, limit: number): Promise<GetInvestmentsResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/investments?page=${page}&limit=${limit}`);

  return response.data as GetInvestmentsResponse;
}

export async function getInvestmentCount(): Promise<number> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/investments/count`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });

  return response.data.data as number;
}

export async function getUsersForInvestmentList(investmentId: number, page: number, limit: number): Promise<UsersResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/investments/${investmentId}/users?page=${page}&limit=${limit}`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data as UsersResponse;
}

export async function getUsersForInvestmentListCount(investmentId: number): Promise<number> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/investments/${investmentId}/users/count`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response.data.data as number;
}

export async function cancelUserInvestment(investmentId: number, userIds: CancelUserInvestmentRequest): Promise<UsersResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post(`/adm/investments/${investmentId}/users/cancel`, userIds).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });

  return response.data as UsersResponse;
}

export async function depositUserInvestment(investmentId: number, userIds: CancelUserInvestmentRequest): Promise<UsersResponse> {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post(`/adm/investments/${investmentId}/users/deposit`, userIds).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });

  return response.data as UsersResponse;
}
