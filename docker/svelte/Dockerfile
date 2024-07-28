# ベースイメージの指定
FROM node:18-alpine

# 作業ディレクトリの設定
WORKDIR /app/svelte

# 依存関係をインストール
COPY frontend/package*.json ./
RUN npm install

# ソースコードをコピー
COPY frontend .

# デフォルトのポート
EXPOSE 5173

# アプリケーションを起動
CMD ["npm", "run", "dev"]