# examples

jzero examples

## Install

```shell
go install github.com/jzero-io/jzero@latest
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

## cli

```shell
jzero new simplecli --branch cli
cd simplecli
go mod tidy
go run main.go
```

## api-goctl

```shell
jzero new simpleapi-goctl --branch api-goctl
cd simpleapi-goctl
jzero gen
go mod tidy
go run main.go server
```

## rpc-goctl

```shell
jzero new simplerpc-goctl --branch rpc-goctl
cd simplerpc-goctl
jzero gen
go mod tidy
go run main.go server
```
