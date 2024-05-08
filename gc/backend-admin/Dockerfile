# syntax=docker/dockerfile:1.4
ARG BUILDER_IMAGE=golang:1.20-bullseye
ARG BASE_IMAGE=gcr.io/distroless/base-debian11
ARG BUILD_NUMBER=0
FROM ${BUILDER_IMAGE} AS builder

ADD . .

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

ARG VERSION_ARGS=" -X=main.AppName=${APP_NAME} -X=main.GitCommit=${GIT_COMMIT} -X=main.GitBranch=${GIT_BRANCH} -X=main.GitTag=${GIT_TAG} -X=main.GitVersion=${GIT_VERSION} -X=main.GitDateTime=${GIT_DATETIME} "

RUN task generate:swagger && \
  CGO_ENABLED=0 go build -o backend -ldflags="${VERSION_ARGS} -extldflags=-static" -tags timetzdata main.go

FROM ${BASE_IMAGE}

COPY --from=builder /app/backend /backend
CMD ["/backend", "serve"]
