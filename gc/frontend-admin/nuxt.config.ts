import { defineNuxtConfig } from "nuxt/config";
import ViteComponents from "unplugin-vue-components/vite";
import IconsResolver from "unplugin-icons/resolver";
import eslintPlugin from "vite-plugin-eslint";

const isDevelopment = process.env.NODE_ENV !== "production";

export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      API_BASE_URL: process.env.API_BASE_URL,

      APP_NAME: process.env.APP_NAME,
      GIT_COMMIT: process.env.GIT_COMMIT,
      GIT_BRANCH: process.env.GIT_BRANCH,
      GIT_TAG: process.env.GIT_TAG,
      GIT_VERSION: process.env.GIT_VERSION,
      GIT_DATETIME: process.env.GIT_DATETIME,
    },
  },
  // privateRuntimeConfig: {},
  ssr: false,
  components: true,
  modules: [
    "@nuxt/image",
    // "nuxt-windicss",
    "@pinia-plugin-persistedstate/nuxt",
    "@pinia/nuxt",
    "unplugin-icons/nuxt",
    "@vueuse/nuxt",
    "@intlify/nuxt3",
    "@nuxtjs/tailwindcss",
  ],
  piniaPersistedstate: {
    cookieOptions: {
      sameSite: "strict",
    },
    storage: "sessionStorage",
    // storage: "localStorage",
  },
  // css
  css: [
    "@mdi/font/css/materialdesignicons.min.css",
    "~/assets/css/tailwind.css",
  ],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  tailwindcss: {
    cssPath: "~/assets/css/tailwind.css",
    configPath: "tailwind.config",
    exposeConfig: false,
    injectPosition: 0,
    viewer: true,
  },
  image: {
    domains: [
      "https://images.unsplash.com",
      "https://source.unsplash.com",
    ],
    provider: isDevelopment ? "public" : "custom",
    providers: {
      public: {
        provider: "~/providers/public",
        options: {
          baseURL: "/",
        },
      },
      custom: {
        provider: "~/providers/custom",
        options: {
          baseURL: "https://s3.dev.opnd.io/image",
        },
      },
    },
  },
  intlify: {
    localeDir: "locales",
    vueI18n: {
      locale: "ko-KR",
    },
  },
  vite: {
    plugins: [
      ViteComponents({
        resolvers: [
          IconsResolver({
            componentPrefix: "",
          }),
        ],
        dts: true,
      }),
      eslintPlugin({
        cache: false,
      }),
    ],
    logLevel: "warn",
    clearScreen: false,
    // esbuild: {
    //   drop:
    //     process.env.GIT_BRANCH === "release"
    //       ? [
    //           "console",
    //           "debugger",
    //         ]
    //       : [],
    // },
  },
});
