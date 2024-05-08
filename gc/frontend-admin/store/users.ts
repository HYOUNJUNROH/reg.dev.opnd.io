import axios, { AxiosResponse, AxiosRequestTransformer, AxiosRequestHeaders, AxiosResponseTransformer, AxiosResponseHeaders } from "axios";
import { DateTime } from "luxon";
import { defineStore } from "pinia";
import { User } from "~/composables/models/models";
import { getBaseURL } from "../composables/utils";

export const useUserStore = defineStore(
  "user",
  () => {
    const userInfo = ref<User | null>();

    function createAxiosInstance(withCredential = true, header = {}) {
      // const token = sessionStorage.getItem("token");
      // if (token) {
      //   setAxiosHeaders(token);
      // }

      const instance = axios.create({
        baseURL: getBaseURL(),
        withCredentials: withCredential,
        headers: header,
      });
      return instance;
    }

    function logout() {
      userInfo.value = null;
      setAxiosHeaders("");
      useRouter().push("/login");
    }

    const setAxiosHeaders = (token: string) => {
      if (token) {
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        sessionStorage.setItem("token", token);
      } else {
        sessionStorage.removeItem("token");
        delete axios.defaults.headers.common["Authorization"];
      }
    };

    return { createAxiosInstance, logout, userInfo, setAxiosHeaders };
  },
  {
    persist: true,
  }
);
