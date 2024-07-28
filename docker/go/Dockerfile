# ベースイメージの指定
FROM golang:1.22-alpine

# 必要なパッケージをインストール
RUN apk add --no-cache git

# Airをインストール
RUN go install github.com/air-verse/air@latest

# 作業ディレクトリの設定
WORKDIR /app/go/src

# Goのモジュールとソースコードをコピー
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download

COPY backend ./

# ポートの設定
EXPOSE 8080

# Airを使ってアプリケーションを起動
CMD ["air"]