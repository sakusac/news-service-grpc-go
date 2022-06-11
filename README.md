# news-service-grpc-go

## requirement
- Go
- docker, docker-compose

## plugin install
protoc-gen-go
https://grpc.io/docs/languages/go/quickstart/

## init project
```:bash
go mod tidy

go get gopkg.in/ini.v1
```

## compile proto file
```
cd <project root dir>

protoc -I. --go_out=. --go-grpc_out=. proto/*.proto

# install liblary
go mod tidy
```

## psql driver import
```
go get gitgub.com/jackc/pgx/v4
```
