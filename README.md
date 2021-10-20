# go-xml-perser


## package構成

``` shell
tree -d -L 2 -a
.
├── .git
├── .github
│   ├── ISSUE_TEMPLATE
│   └── workflows
├── migrations
│   ├── ddl
│   └── dml
├── src
│   ├── entities    (yo自動生成先)
│   ├── gen         (go-swagger自動生成先)
│   ├── handler
│   ├── service
│   ├── testUtil
│   ├── util
|   └── main.go
├── docker-compose.yml
├── swagger.yaml

```

## バージョン
$ go version
go version go1.16.5 darwin/amd64

gcloud 360.0.0
``` shell
$ gcloud components list  

Your current Cloud SDK version is: 360.0.0
The latest available version is: 360.0.0

┌────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                                 Components                                                 │
├───────────────┬──────────────────────────────────────────────────────┬──────────────────────────┬──────────┤
│     Status    │                         Name                         │            ID            │   Size   │
├───────────────┼──────────────────────────────────────────────────────┼──────────────────────────┼──────────┤
│ Installed     │ BigQuery Command Line Tool                           │ bq                       │  < 1 MiB │
│ Installed     │ Cloud Datastore Emulator                             │ cloud-datastore-emulator │ 18.4 MiB │
│ Installed     │ Cloud SDK Core Libraries                             │ core                     │ 20.4 MiB │
│ Installed     │ Cloud Storage Command Line Tool                      │ gsutil                   │  8.5 MiB │
│ Installed     │ Minikube                                             │ minikube                 │ 26.6 MiB │
│ Installed     │ Skaffold                                             │ skaffold                 │ 19.9 MiB │
│ Installed     │ gcloud Alpha Commands                                │ alpha                    │  < 1 MiB │
│ Installed     │ gcloud Beta Commands                                 │ beta                     │  < 1 MiB │
│ Installed     │ gcloud app Python Extensions                         │ app-engine-python        │  7.8 MiB │
│ Installed     │ kubectl                                              │ kubectl                  │  < 1 MiB │
└───────────────┴──────────────────────────────────────────────────────┴──────────────────────────┴──────────┘
```

## セットアップ
``` shell
# 必要なライブラリのインストール、DBコンテナ起動、API起動
$ make setup

```

## ローカル実行
``` shell
$ make run

```

## swagger生成
``` shell
make gen-src
```

## yoモデル生成
``` shell
make yo
```

## docker起動

``` shell
$ make up
```

## APIサーバ起動

``` shell
$ make run
```

## 環境変数
export PROJECT_ID=test-project
export INSTANCE_NAME=test-instance
export DATABASE_NAME=test-database

export SPANNER_EMULATOR_URL=http://spanner:9020/
export SPANNER_EMULATOR_HOST=localhost:9010

## lib
spannerエミュレータ
https://github.com/tarosaiba/docker-compose-spanner

go-swagger(swagger定義を用いたAPI I/Fジェネレータ)
https://github.com/go-swagger/go-swagger

wrench(spannerデータCRUD)
https://github.com/cloudspannerecosystem/wrench

yo(spannerデータCRUDジェネレータ)
https://github.com/cloudspannerecosystem/yo
