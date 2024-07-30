# repoapi

repoapiはレポートとユーザーを管理するためのシンプルなREST APIです。Go言語とGin Webフレームワークを使用して開発され、データベースにはMySQLを使用しています。

## 目次

- [インストールと実行](#インストールと実行)
- [APIの使用](#APIの使用)
- [テスト](#テスト)
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

#### ユーザーの取得

- メソッド: `GET /user?id={userId}`
- 概要: 指定したIDのユーザー情報をデータベースから取得します。

リクエストの例:

```bash
curl -X GET "localhost:8080/user?id=ymd333"
```

#### ユーザー情報の更新

- メソッド: `PUT /user`
- 概要: 指定したIDのユーザーの情報を更新します。
- リクエストボディ: JSON形式で、`id`と`name`のフィールドを含めることができます。

リクエストの例:

```bash
curl -X PUT "localhost:8080/user" -d '{"id":"ymd333","name":"山田 花子"}' -H "Content-Type: application/json"
```

### レポートエンドポイント

#### レポートの登録

- メソッド: `POST /report`
- 概要: 新しいレポートをデータベースに登録します。
- リクエストボディ: JSON形式で、`author_id`、`count`、`title`、`style`、`language`のフィールドを含める必要があります。`style`は`polite`または`definite`、`language`は`jp`または`en`である必要があります。

リクエストの例:

```bash
curl -X POST localhost:8080/report -d '{"author_id":"ymd333", "count":300, "title":"レイヤードアーキテクチャについて", "style":"polite", "language":"jp"}' -H "Content-Type: application/json"
```

#### レポートの削除

- メソッド: `DELETE /report?id={reportId}`
- 概要: 指定したIDのレポートをデータベースから削除します。

リクエストの例:

```bash
curl -X DELETE "localhost:8080/report?id=858e6581-7b63-f39a-10c8-be32ad6aafb5"
```

#### レポートの取得

- メソッド: `GET /report`
- 概要: レポートを取得します。`id`または`author_id`と`title`、`style`、`language`のクエリパラメータを使用してフィルタリングできます。

リクエストの例:

- IDでレポートを取得:

```bash
curl -X GET "localhost:8080/report?id=858e6581-7b63-f39a-10c8-be32ad6aafb5"
```

- 作成者IDでレポートを取得:

```bash
curl -X GET "localhost:8080/report?author_id=ymd333"
```

- 複数のクエリパラメータを指定してレポートを取得:

```bash
curl -X GET "localhost:8080/report?author_id=ymd333&title=レイヤードアーキテクチャについて"
curl -X GET "localhost:8080/report?author_id=ymd333&style=definite&language=en"
```

#### レポート情報の更新

- メソッド: `PUT /report`
- 概要: 指定したIDのレポートの情報を更新します。
- リクエストボディ: JSON形式で、`id`、`count`、`title`、`style`、`language`のフィールドを含めることができます。

リクエストの例:

```bash
curl -X PUT localhost:8080/report -d '{"id":"858e6581-7b63-f39a-10c8-be32ad6aafb5","count":400, "title":"クリーンアーキテクチャについて", "style":"definite", "language":"jp"}' -H "Content-Type: application/json"
```

## テスト

repoapiは、単体テストと統合テストの両方を含んでいます。テストを実行するには、プロジェクトのルートディレクトリで次のコマンドを実行します。

```bash
go test ./...
```

## ディレクトリ構造

```
repoapi/
├── cmd/
│   └── main.go                                   # アプリケーションのエントリポイント
├── src/
│   ├── application/                             # アプリケーション層
│   │   ├── report.go                            # レポートに関するアプリケーションロジック
│   │   ├── report_test.go                       # レポートに関するアプリケーションロジックのテスト
│   │   ├── user.go                             # ユーザーに関するアプリケーションロジック
│   │   └── user_test.go                        # ユーザーに関するアプリケーションロジックのテスト
│   ├── domain/                                  # ドメイン層
│   │   ├── model/                              # データモデル
│   │   │   ├── report.go                       # レポートのデータモデル
│   │   │   └── user.go                          # ユーザーのデータモデル
│   │   └── repository/                         # リポジトリのインターフェース
│   │       ├── report.go                       # レポートリポジトリのインターフェース
│   │       └── user.go                          # ユーザーリポジトリのインターフェース
│   ├── infra/                                   # インフラストラクチャ層
│   │
│   │   ├── database.go                         # データベース接続
│   │   └── persistence/                        # データベースとのやり取り
│   │       ├── report.go                       # レポートに関するデータベース操作
│   │       └── user.go                          # ユーザーに関するデータベース操作
│   └── presentation/                            # プレゼンテーション層
│       └── rest/                               # REST API
│           ├── report.go                       # レポートに関するREST APIハンドラ
│           └── user.go                         # ユーザーに関するREST APIハンドラ
├── go.mod                                       # Goモジュール定義ファイル
├── go.sum                                       # Goモジュールチェックサムファイル
└── docker-compose.yml                           # Docker Compose構成ファイル
```
