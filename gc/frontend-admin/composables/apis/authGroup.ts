import axios, { AxiosResponse } from "axios";
import { User } from "~/composables/models/models";
import { useUserStore } from "~/store/users";
import { getBaseURL } from "../utils";

/// POST /adm/auth/login
export async function login(email: string, password: string): Promise<User> {
  const formData = new FormData();
  formData.append("email", email);
  formData.append("password", password);

  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post("/adm/auth/login", formData).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });

  user.userInfo = response.data.data as User;
  user.setAxiosHeaders(user.userInfo.token);
  return response.data as User;
}

/// POST /adm/auth/update-password
export async function updatePassword(oldPassword: string, newPassword: string) {
  const formData = new FormData();
  formData.append("old_password", oldPassword);
  formData.append("new_password", newPassword);

  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.post("/adm/auth/update-password", formData).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });

  return response.data;
}
