# 构建阶段
FROM golang:1.23-alpine AS builder

WORKDIR /build

# 设置 Go 代理（使用国内镜像）
ENV GOPROXY=https://goproxy.cn,direct

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译 Go 程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o KrillinAI ./cmd/server

# 运行阶段
FROM ubuntu:latest

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates ffmpeg && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p bin && \
    ARCH=$(uname -m) && \
    case "$ARCH" in \
    x86_64) \
    YT_DLP_URL="https://github.com/yt-dlp/yt-dlp/releases/download/2025.01.15/yt-dlp_linux"; \
    EDGE_TTS_URL="https://github.com/puji4810/edge-tts-pkg/releases/download/v0.0.1/edge-tts-linux-amd64"; \
    ;; \
    armv7l) \
    YT_DLP_URL="https://github.com/puji4810/edge-tts-pkg/releases/download/v0.0.1/edge-tts-linux-armv7"; \
    EDGE_TTS_URL="https://github.com/puji4810/edge-tts-pkg/releases/download/v0.0.1/edge-tts-linux-armv7"; \
    ;; \
    aarch64) \
    YT_DLP_URL="https://github.com/yt-dlp/yt-dlp/releases/download/2025.01.15/yt-dlp_linux_aarch64"; \
    EDGE_TTS_URL="https://github.com/puji4810/edge-tts-pkg/releases/download/v0.0.1/edge-tts-linux-arm64"; \
    ;; \
    *) \
    echo "Unsupported architecture: $ARCH" && exit 1; \
    ;; \
    esac && \
    wget -O bin/yt-dlp "$YT_DLP_URL" && \
    wget -O bin/edge-tts "$EDGE_TTS_URL" && \
    chmod +x bin/yt-dlp bin/edge-tts

# 从构建阶段复制编译好的可执行文件
COPY --from=builder /build/KrillinAI ./

RUN mkdir -p /app/models && \
    chmod +x ./KrillinAI

VOLUME ["/app/bin", "/app/models"]

ENV PATH="/app/bin:${PATH}"

EXPOSE 8888/tcp

ENTRYPOINT ["./KrillinAI"]
