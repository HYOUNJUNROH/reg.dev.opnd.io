ARG BASE_IMAGE=node:16-bullseye-slim
ARG BUILDER_IMAGE=node:16-bullseye-slim
FROM ${BUILDER_IMAGE} AS builder

ENV API_BASE_URL="https://admin.geopop.co.kr/"
ENV DATA_CDN_URL="https://static.cdn.ntruss.com"
ENV SUPABASE_URL=reserved
ENV SUPABASE_KEY=reserved

ADD . .

ENV NODE_ENV=production

ARG APP_NAME
ENV APP_NAME=${APP_NAME}
ARG GIT_COMMIT
ENV GIT_COMMIT=${GIT_COMMIT}
ARG GIT_BRANCH
ENV GIT_BRANCH=${GIT_BRANCH}
ARG GIT_TAG
ENV GIT_TAG=${GIT_TAG}
ARG GIT_VERSION
ENV GIT_VERSION=${GIT_VERSION}
ARG GIT_DATETIME
ENV GIT_DATETIME=${GIT_DATETIME}

RUN yarn build

FROM ${BASE_IMAGE}

COPY --from=builder /app/.output .output

ENV NITRO_HOST=0.0.0.0
ENV NITRO_PORT=3000

ENV NITRO_PUBLIC_API_BASE_URL="https://admin.geopop.co.kr/"
ENV NITRO_PUBLIC_DATA_CDN_URL=https://static.cdn.ntruss.com
ENV NITRO_PUBLIC_SUPABASE_URL=reserved
ENV NITRO_PUBLIC_SUPABASE_KEY=reserved
ENV NUXT_TELEMETRY_DISABLED=1

EXPOSE 3000

WORKDIR /
CMD [".output/server/index.mjs"]
