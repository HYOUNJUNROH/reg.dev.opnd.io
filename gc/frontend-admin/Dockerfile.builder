
ARG BASE_IMAGE=node:16-bullseye-slim
FROM ${BASE_IMAGE} AS builder

WORKDIR /app
COPY .yarn/releases/ .yarn/releases/
COPY .yarn/patches/ .yarn/patches/
COPY .yarn/plugins/ .yarn/plugins/
COPY .yarnrc.yml package.json yarn.lock ./

RUN  --mount=type=secret,id=ssh-script,dst=/root/install_ssh_key.sh \
  --mount=type=secret,id=ssh-key,dst=/root/.ssh/id.key \
  --mount=type=ssh \
  git config --list && \
  sh /root/install_ssh_key.sh /root/.ssh/id.key common/openerd-nuxt3 nas.dev.opnd.io 2222 git.dev.opnd.io  && \
  yarn install
