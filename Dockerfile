FROM node:18 as builder

WORKDIR /build

# 复制前端依赖文件
COPY web/package.json .
COPY web/yarn.lock .

# 安装前端依赖
RUN yarn --frozen-lockfile

# 复制前端源码和版本文件
COPY ./web .
COPY ./VERSION .

# 清理可能存在的构建缓存并构建前端
RUN rm -rf build node_modules/.vite && \
    DISABLE_ESLINT_PLUGIN='true' VITE_APP_VERSION=$(cat VERSION) yarn build

FROM golang:1.24.2 AS builder2

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux

WORKDIR /build

# 复制Go依赖文件
ADD go.mod go.sum ./
RUN go mod download

# 复制所有源码
COPY . .

# 从第一阶段复制构建好的前端文件
COPY --from=builder /build/build ./web/build

# 构建Go应用
RUN go build -ldflags "-s -w -X 'one-api/common.Version=$(cat VERSION)' -extldflags '-static'" -o one-api

FROM alpine

# 安装必要的系统依赖
RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

# 复制构建好的应用
COPY --from=builder2 /build/one-api /

# 暴露端口
EXPOSE 3000

# 设置工作目录
WORKDIR /data

# 启动应用
ENTRYPOINT ["/one-api"]