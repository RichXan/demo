BuildVersion=0.1.0
BUILD_DATE=$(shell date  +%y%m%d%H%M%S)
TARGET_NAME=x-project

all: x64

x64: env
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -X core.Version=$(BuildVersion) -X core.BuildDate=$(BUILD_DATE)" -o $(TARGET_NAME) cmd/server/main.go
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -X core.Version=$(BuildVersion) -X core.BuildDate=$(BUILD_DATE)" -o $(TARGET_NAME)-db cmd/db/main.go

clean:
	go clean

up:
	docker-compose up -d

down:
	docker-compose down

tidy: env
	go mod tidy
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golang/mock/mockgen@v1.6.0
	go get github.com/google/wire/cmd/wire

#	newgrp docker
shell:
	docker exec -it -w /$(TARGET_NAME) $(TARGET_NAME) bash

build: tidy
	docker build --progress=plain -t $(TARGET_NAME) $(pwd) .

run: env
	go run cmd/server/main.go -c config/development.yml http

nats: env
	go run cmd/server/main.go -c config/development.yml nats

env:
	go env -w GO111MODULE=auto
	go env -w GOPROXY=https://goproxy.cn,direct
	go env -w GOPRIVATE="dev.azure.com"
	
dbup:
	go run ./cmd/db/main.go -c ./config/development.yml up

dbdw:
	go run ./cmd/db/main.go -c ./config/development.yml down

# 进入postgre中seanet数据库
postgres:
	docker exec -it sn-postgres psql -h sn-postgres -Uadmin -d seanet

syncdbdata:
	docker exec -it sn-postgres psql -U seanet -f /var/lib/postgresql/data/seanet_pgdump-20240811.dmp

restorepsql: dropseanetdb
	docker exec -it sn-postgres pg_restore -d seanet -Uadmin --clean --if-exists --no-owner --no-privileges --exit-on-error --verbose /var/lib/postgresql/data/seanet_pgdump-20241117.sql

dropseanetdb:
	docker exec -it sn-postgres psql -h sn-postgres -Uadmin -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'seanet' AND pid <> pg_backend_pid();"
	docker exec -it sn-postgres psql -h sn-postgres -Uadmin -d postgres -c "DROP DATABASE IF EXISTS seanet;"
	docker exec -it sn-postgres psql -h sn-postgres -Uadmin -d postgres -c "CREATE DATABASE seanet;"
	

ping:
	apt-get update && apt-get install iputils-ping

protocol:
	protoc -I C:\Users\admin\go\src --proto_path=./proto/ --go_out=./ ./proto/q.proto