PROJECT_ID ?= test-project
INSTANCE_ID ?= test-instance
DATABASE_ID ?= test-database

SPANNER_EMULATOR_URL ?= http://spanner:9020/
SPANNER_EMULATOR_HOST ?= localhost:9010

LOCAL ?= SPANNER_PROJECT_ID=$(PROJECT_ID) SPANNER_INSTANCE_ID=$(INSTANCE_ID) SPANNER_DATABASE_ID=$(DATABASE_ID) SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST)

setup: install up run

# ローカル実行
run:
	# air -d
	$(LOCAL) go run src/main.go --host 0.0.0.0 --port 8080

mod:
	go mod tidy
	go mod vendor

build:
	go build -mod=vendor

# install lib
install:
	brew tap go-swagger/go-swagger
	brew install go-swagger

	go mod tidy
	go mod vendor

install-win:
	.\swagger.bat

	go mod tidy
	go mod vendor

## generate go-swagger & yo
gen: gen-src yo mod

## generate gen source by go-swagger
gen-src:
	@mkdir -p ./src/gen
	swagger generate server -A routing -t ./src/gen -f ./swagger.yaml --exclude-main

## generate entity source by yo
# export GOPATH=$(go env GOPATH)
# export SPANNER_EMULATOR_HOST="localhost:9010"
yo:
	@mkdir -p ./src/entities
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) yo $(PROJECT_ID) $(INSTANCE_ID) $(DATABASE_ID) -o ./src/entities

## Go Test
test:
	go test ./src/handler/... -cover -coverprofile=coverage.out -covermode=set

	# カバレッジ確認
	go tool cover -html=coverage.out

testd:
	go test ./src/... -v


## docker up
up:
	docker-compose up -d

## docker up --build
upb:
	docker-compose up -d --build

## docker container kill
kill:
	docker-compose kill

clean:
	docker-compose rm -fv --all
	docker-compose down --rmi local --remove-orphans

## docker reload
reload: kill up




# ==========
# local Spanner用コマンド
# ==========

# 基本的にはdocker-composeのwrench-create、wrench-applyでテーブル作成とデータ作成は実施されるため作り直したい場合用

## Cloud Config Spannerエミュレータ用config作成
create-emu:
	gcloud config configurations create emulator
	gcloud config set auth/disable_credentials true
	gcloud config set project $(PROJECT_ID)
	gcloud config set api_endpoint_overrides/spanner $(SPANNER_EMULATOR_URL)
	gcloud config configurations list

## Cloud Config エミュレータに切り替え	
set-emu:
	@echo 'configをemulatorに切り替え'
	gcloud config configurations activate emulator
ifndef $(SPANNER_EMULATOR_HOST)
	set SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST)
endif

## Cloud Config デフォルトに切り替え	
set-def:
	@echo 'configをdefaultに切り替え'
	gcloud config configurations activate default

## Cloud Configリスト表示
config-list:
	gcloud config configurations list

# local環境用　インスタンス作成、データベース作成
local-db-create: set-emu
	@echo 'インスタンス作成'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner instances create $(INSTANCE_ID) --config=regional-asia-northeast1 --description=$(INSTANCE_ID) --nodes=1

	@echo 'Database作成'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner databases create $(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID)

# local環境用　インスタンス削除
local-db-delete: set-emu
	@echo 'Database削除'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner databases delete $(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID)

	@echo 'インスタンス削除'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner instances delete $(INSTANCE_ID)

## Spanner テーブル作成（ローカル用）
migrate:
	@echo 'テーブル削除'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner databases ddl update $(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID) \
	--ddl-file='./migrations/schema_drop.sql'

	@echo 'テーブル作成'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) gcloud spanner databases ddl update $(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID) \
	--ddl-file='./migrations/ddl/schema.sql'

	@echo 'データ登録'
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) wrench apply --dml migrations/dml/dml.sql --database=$(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID)

## SpannerのDB定義の確認 前テーブルのCreate文が返ってくる
desc:
	gcloud spanner databases ddl describe $(DATABASE_ID) --instance=$(INSTANCE_ID) --project=$(PROJECT_ID) --format=flattened
	@echo "インスタンスリスト"
	gcloud spanner instances list --project=$(PROJECT_ID)
	@echo "データベースリスト"
	gcloud spanner databases list --instance=$(INSTANCE_ID) --project=$(PROJECT_ID)


# ==========
# おまけ
# ==========

# ソースからSwaggr定義を作成
gen-spec:
	swagger generate spec -o ./swagger_spec.yaml

# Api Design Document Generator Setup
gen-spec-setup: 	
	brew install npm
	npm install -g redoc-cli

# Swaggr定義から仕様書HTMLを作成
gen-spec-html:
	redoc-cli bundle ./swagger.yaml -o ./swagger.html

# bash:
# 	docker-compose exec spanner bash


# ==========
# dev用
# ==========

## dev用に環境変数
PROJECT_ID_DEV = <GCPプロジェクトID>
INSTANCE_ID_DEV = test-instance
DATABASE_ID_DEV = test-database

# dev環境用　configの作成
create-dev:
	gcloud config configurations create dev
	gcloud config set account <GCPアカウントのメールアドレス>
	gcloud config set project $(PROJECT_ID_DEV)
	gcloud config configurations list

set-dev:
	@echo 'configをdevに切り替え'
	gcloud config configurations activate dev
ifdef $(SPANNER_EMULATOR_HOST)
	unset SPANNER_EMULATOR_HOST
endif

# dev環境用　インスタンス作成、データベース作成
dev-db-create:
	@echo 'インスタンス作成'
	# betaのコマンドであればノードではなく処理ユニット数（processing-units）の指定が可能
	gcloud beta spanner instances create $(INSTANCE_ID_DEV) --config=regional-asia-northeast1 --description=xml-parser --processing-units=100

	@echo 'Database作成'
	gcloud spanner databases create $(DATABASE_ID_DEV) --instance=$(INSTANCE_ID_DEV) --project=$(PROJECT_ID_DEV)

# dev環境用　インスタンス削除
dev-db-delete:
	@echo 'Database削除'
	gcloud spanner databases delete $(DATABASE_ID_DEV) --instance=$(INSTANCE_ID_DEV) --project=$(PROJECT_ID_DEV)

	@echo 'インスタンス削除'
	gcloud spanner instances delete $(INSTANCE_ID_DEV)


# dev環境用　DML実行用のライブラリインストール
migrate-dev-setup:
	go get -u github.com/cloudspannerecosystem/wrench

# dev環境用　DDL実行、DML実行
migrate-dev: set-dev
	@echo 'TODO テーブル削除'
	gcloud spanner databases ddl update $(DATABASE_ID_DEV) --instance=$(INSTANCE_ID_DEV) --project=$(PROJECT_ID_DEV) \
	--ddl-file='./migrations/schema_drop.sql'

	@echo 'テーブル作成'
	gcloud spanner databases ddl update $(DATABASE_ID_DEV) --instance=$(INSTANCE_ID_DEV) --project=$(PROJECT_ID_DEV) \
	--ddl-file='./migrations/ddl/schema.sql'

	@echo 'データ登録'
	wrench apply --dml migrations/dml/dml.sql --database=$(DATABASE_ID_DEV) --instance=$(INSTANCE_ID_DEV) --project=$(PROJECT_ID_DEV)
