import urlJoin from "url-join";
import { useUserStore } from "~/store/users";

export const getBaseURL = (): string | undefined => {
  if (useRuntimeConfig !== null && typeof useRuntimeConfig !== "undefined") {
    return urlJoin(useRuntimeConfig().public.API_BASE_URL);
    // return urlJoin(useRuntimeConfig().API_BASE_URL, "/api");
  }
  return undefined;
};

/// /adm/images/{이미지파일명}
export function getImage(image: string) {
  return urlJoin(getBaseURL()!, "/adm/images", image);
}

export async function getImageFile(image: string) {
  const user = useUserStore();
  const instance = user.createAxiosInstance();
  const response = await instance.get(`/adm/images/${image}`).catch((error) => {
    if (error.response.status === 401) {
      user.logout();
    }
  });
  return response;
}
