# TODO管理アプリ

## アプリについて

### リンク
[https://todo-vue-frontend.web.app/](https://todo-vue-frontend.web.app/)

### 目的
1. バックエンドとフロントエンドを切り分ける
1. バックエンドではクリーンアーキテクチャで構成する
1. Dockerによる開発環境
1. デプロイまで行う

### 内容
- TODOの登録
- TODOの更新
- TODOの削除
- TODOの一覧表示

### 構成
#### バックエンド
- 開発環境: Docker
- フォルダ: `/backend`
- 言語: Go
- フレームワーク: なし
- データベース: Postgres
- デプロイ先: Heroku

#### フロントエンド
- 開発環境: Docker
- フォルダ: `/frontend`
- 言語: html, css, javascript
- フレームワーク: Vue.js
- デプロイ先: Firebase

## 環境構築について

### 環境構築
必須: 
    - `docker`
    - `docker-compose`
    - `./firebase_deploy.env`に`FIREBASE_TOKEN`を保存
    - `./frontend/vue-demo/.env.development`に`VUE_APP_API_BASE_URL={開発環境のバックエンドAPI}`を保存
    - `./frontend/vue-demo/.env.local`に`VUE_APP_API_BASE_URL={ローカル環境のバックエンドAPI}`を保存

```
$ make build
```

### バックエンドコンテナ、フロントエンドコンテナ、データベースコンテナを立ち上げ

```
$ make up
```

### バックエンドコンテナ、フロントエンドコンテナ、データベースコンテナを終了

```
$ make down
```

### バックエンドをHerokuにデプロイ

```
$ make deploy_backend
```

### フロントエンドをFirebaseにデプロイ

```
$ make deploy_frontend
```