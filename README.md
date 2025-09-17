# examples

jzero examples

## Install

```shell
go install github.com/jzero-io/jzero/cmd/jzero@latest
jzero check
```

## api

```shell
jzero new simpleapi --frame api
cd simpleapi
jzero gen
go mod tidy
go run main.go server
```

## gateway

```shell
jzero new simplegateway --frame gateway
cd simplegateway
jzero gen
go mod tidy
go run main.go server
```

## rpc

```shell
jzero new simplerpc --frame rpc
cd simplerpc
jzero gen
go mod tidy
go run main.go server
```