# h23s_01

# h23s_01_practice

ハッカソンの本番用リポジトリ（Vue + Echo）

## バックエンド

[backend/README.md](backend/README.md)を見てください。

## 開発環境の起動（フロントエンド）

### 初回 or 不足しているフロントエンドのパッケージがある場合

```shell
cd frontend
npm install
npm run dev
```

### 2 回目以降（不足しているフロントエンドのパッケージがない場合）

```shell
cd frontend && npm run dev
```

### 新しいフロントエンドのパッケージを追加する場合

`package.json`に新しいパッケージの情報が追加され、他の人が`npm install`でインストールできるようになります。

```shell
cd frontend
npm install <パッケージ名>
```

## Git 開発の基本

### 自分の作業ブランチを作成する

`main`ブランチを直接変更するのは基本的に避けましょう。

```shell
git switch -c <新しい作業ブランチ名>
```

### 作業ブランチで加えた変更内容をリモートリポジトリにコミット・プッシュする

```shell
git add .
git commit -m "<コミットメッセージ>"
git push -u origin
```

### 最新のリモートリポジトリの内容をローカルに取り込む

```
git pull
```

### リモートリポジトリの`main`ブランチを作業ブランチにマージする

`main`ブランチを取り込んだ際に、**コンフリクト**が発生することがあります。その際は、コンフリクトを解消した上で変更内容をプッシュすることになります。

```shell
git merge origin main
```

## 開発環境構築（フロントエンド）

> 注：@aya-se 用のメモです。通常この作業を行う必要はありません。

簡単のため、（TypeScript ではなく）JavaScript を採用。また、Docker を使用せずに開発環境を構築する。

```
npm create vite@latest
cd frontend
npm install
npm install sass
```
