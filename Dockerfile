FROM node:22 AS builder

# 可选：从外部传入版本号，未传入时回退到 VERSION 文件
ARG VERSION_STR

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
    DISABLE_ESLINT_PLUGIN='true' VITE_APP_VERSION=${VERSION_STR:-$(cat VERSION)} yarn build

FROM golang:1.24.2 AS builder2

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    TIKTOKEN_CACHE_DIR=/opt/tiktoken-cache

WORKDIR /build

# 可选：从外部传入版本号，未传入时回退到 VERSION 文件
ARG VERSION_STR

# 复制Go依赖文件
ADD go.mod go.sum ./
RUN go mod download

# 复制所有源码
COPY . .

# 生成 Swagger 文档
RUN go generate ./...

# 从第一阶段复制构建好的前端文件
COPY --from=builder /build/build ./web/build

# 预热 tiktoken 词表缓存，避免运行时再下载
RUN mkdir -p ${TIKTOKEN_CACHE_DIR} && \
    go run ./hack/scripts/prefetch_tiktoken/main.go

# 构建Go应用（注入版本号：优先使用传入版本号，否则读取 VERSION 文件）
RUN VERSION_VALUE="${VERSION_STR:-$(cat VERSION)}" && \
    go build -ldflags "-s -w -X one-api/common.Version=${VERSION_VALUE} -extldflags '-static'" -o one-api

FROM alpine

# 安装必要的系统依赖
RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

# 复制构建好的应用
COPY --from=builder2 /build/one-api /
# 复制预热完成的 tiktoken 缓存
COPY --from=builder2 /opt/tiktoken-cache /opt/tiktoken-cache

# 暴露端口
EXPOSE 3000

# 设置工作目录
WORKDIR /data

# 设置运行阶段的 tiktoken 缓存路径
ENV TIKTOKEN_CACHE_DIR=/opt/tiktoken-cache

# 启动应用
ENTRYPOINT ["/one-api"]
