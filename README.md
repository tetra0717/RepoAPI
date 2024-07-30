# repoapi

repoapiはレポートとユーザーを管理するためのシンプルなREST APIです。Go言語とGin Webフレームワークを使用して開発され、データベースにはMySQLを使用しています。

## 目次

- [インストールと実行](#インストールと実行)
- [APIの使用](#APIの使用)
- [ディレクトリ構造](#ディレクトリ構造)

## インストールと実行

### 依存関係

repoapiはMySQLデータベースに依存しています。まだインストールしていない場合は、MySQLをインストールし、データベースサーバーを起動してください。

### リポジトリのクローン

```bash
git clone https://github.com/tetra0717/RepoAPI.git
cd RepoAPI
```

### Dockerコンテナのビルドと実行

このコマンドは、アプリケーションに必要なDockerコンテナをバックグラウンドでビルドして起動します。具体的には、MySQLデータベースコンテナとrepoapiアプリケーションコンテナの2つが起動します。

```bash
docker-compose up -d
```

- `db`サービスは、MySQLデータベースイメージを使用してデータベースコンテナをビルドします。データベースの接続情報は`docker-compose.yml`ファイルで定義されています。
- `app`サービスは、アプリケーションイメージを使用してアプリケーションコンテナをビルドします。このコンテナは、ポート`8080`でホストマシンのポートにバインドされます。


## APIの使用

repoapiは、HTTPリクエストを使用して操作します。APIエンドポイントは`/user`と`/report`の2つです。

### ユーザーエンドポイント

#### ユーザーの登録

- メソッド: `POST /user`
- 概要: 新しいユーザーをデータベースに登録します。
- リクエストボディ: JSON形式で、`id`と`name`のフィールドを含める必要があります。

リクエストの例:

```bash
curl -X POST localhost:8080/user -d '{"id":"ymd333", "name":"山田 太郎"}' -H "Content-Type: application/json"
```
![image](https://github.com/user-attachments/assets/dac3aa04-75a5-477a-bd6b-b59e7e7aada6)


#### ユーザーの取得

- メソッド: `GET /user?id={userId}`
- 概要: 指定したIDのユーザー情報をデータベースから取得します。

リクエストの例:

```bash
curl -X GET "localhost:8080/user?id=ymd333"
```
![image](https://github.com/user-attachments/assets/f24d2c03-8e7a-4536-bdef-c84a6c259228)



#### ユーザー情報の更新

- メソッド: `PUT /user`
- 概要: 指定したIDのユーザーの情報を更新します。
- リクエストボディ: JSON形式で、`id`と`name`のフィールドを含めることができます。

リクエストの例:

```bash
curl -X PUT "localhost:8080/user" -d '{"id":"ymd333","name":"山田 花子"}' -H "Content-Type: application/json"
```
![image](https://github.com/user-attachments/assets/92c40d5f-5e4a-44e6-9381-4f3f59c007ee)


### レポートエンドポイント

#### レポートの登録

- メソッド: `POST /report`
- 概要: 新しいレポートをデータベースに登録します。
- リクエストボディ: JSON形式で、`author_id`、`count`、`title`、`style`、`language`のフィールドを含める必要があります。`style`は`polite`または`definite`、`language`は`jp`または`en`である必要があります。

リクエストの例:

```bash
curl -X POST localhost:8080/report -d '{"author_id":"ymd333", "count":300, "title":"レイヤードアーキテクチャについて", "style":"polite", "language":"jp"}' -H "Content-Type: application/json"
```
![image](https://github.com/user-attachments/assets/5f6ebabe-b113-4f2f-a4bf-fb96bf4631bb)


#### レポートの削除

- メソッド: `DELETE /report?id={reportId}`
- 概要: 指定したIDのレポートをデータベースから削除します。

リクエストの例:

```bash
curl -X DELETE "localhost:8080/report?id=30b61e17-eca3-4312-b141-878de36a70d1"
```
![image](https://github.com/user-attachments/assets/ddc16ad2-78b9-4956-b154-8c679dfd52e8)


#### レポートの取得

- メソッド: `GET /report`
- 概要: レポートを取得します。`id`または`author_id`と`title`、`style`、`language`のクエリパラメータを使用してフィルタリングできます。

リクエストの例:

- IDでレポートを取得:

```bash
curl -X GET "localhost:8080/report?id=30b61e17-eca3-4312-b141-878de36a70d1"
```

- 作成者IDでレポートを取得:

```bash
curl -X GET "localhost:8080/report?author_id=ymd333"
```
![image](https://github.com/user-attachments/assets/f48acd37-cafd-448d-8179-37869606a607)


- 複数のクエリパラメータを指定してレポートを取得:

```bash
curl -X GET "localhost:8080/report?author_id=ymd333&title=レイヤードアーキテクチャについて"
```
![image](https://github.com/user-attachments/assets/76b0ca38-8ea1-463e-ac63-d0fcf37c6eda)


#### レポート情報の更新

- メソッド: `PUT /report`
- 概要: 指定したIDのレポートの情報を更新します。
- リクエストボディ: JSON形式で、`id`、`count`、`title`、`style`、`language`のフィールドを含めることができます。

リクエストの例:

```bash
curl -X PUT localhost:8080/report -d '{"id":"30b61e17-eca3-4312-b141-878de36a70d1","count":400, "title":"クリーンアーキテクチャについて", "style":"definite", "language":"jp"}' -H "Content-Type: application/json"
```
![image](https://github.com/user-attachments/assets/6f3df378-2203-4bdc-8011-d4043a565a8c)

## ディレクトリ構造

```
repoapi/
├── cmd/
│   └── main.go                                   # アプリケーションのエントリポイント
├── src/
│   ├── application/                             # アプリケーション層
│   │   ├── report.go                            # レポートに関するアプリケーションロジック
│   │   └── user.go                              # ユーザーに関するアプリケーションロジック
│   ├── domain/                                  # ドメイン層
│   │   ├── model/                               # データモデル
│   │   │   ├── report.go                        # レポートのデータモデル
│   │   │   └── user.go                          # ユーザーのデータモデル
│   │   └── repository/                          # リポジトリのインターフェース
│   │       ├── report.go                        # レポートリポジトリのインターフェース
│   │       └── user.go                          # ユーザーリポジトリのインターフェース
│   ├── infra/                                   # インフラストラクチャ層
│   │   ├── database.go                          # データベース接続
│   │   └── persistence/                         # データベースとのやり取り
│   │       ├── report.go                        # レポートに関するデータベース操作
│   │       └── user.go                          # ユーザーに関するデータベース操作
│   └── presentation/                            # プレゼンテーション層
│       └── rest/                                # REST API
│           ├── report.go                        # レポートに関するREST APIハンドラ
│           └── user.go                          # ユーザーに関するREST APIハンドラ
├── go.mod                                       # Goモジュール定義ファイル
├── go.sum                                       # Goモジュールチェックサムファイル
└── docker-compose.yml                           # Docker Compose構成ファイル
```
