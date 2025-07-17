# --- build stage ---
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Go モジュールを先にコピー＆取得（キャッシュ効かせる）
COPY go.mod .
COPY go.sum .

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN go mod download

# ソースをコピーしてビルド
COPY . .
# 環境変数 GOOS/GOARCH はホストと同じなら不要
RUN go build -o server ./main.go

# --- final stage ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# builder からビルド成果物をコピー
COPY --from=builder /app/server .

# 必要に応じて環境変数を指定
# ENV PORT=8080

EXPOSE 8081
CMD ["./server"]
