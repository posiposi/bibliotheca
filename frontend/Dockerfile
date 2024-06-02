# ベースイメージとしてnodeを使用
FROM node:latest

# 作業ディレクトリを作成
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install

# プロジェクトの全てのファイルをコピー
COPY . .

# 開発サーバーを起動
CMD ["npm", "run", "dev"]