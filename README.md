# examples

jzero examples

## Install

```shell
go install github.com/jzero-io/jzero@latest
jzero check
```

## Quick Start

```shell
jzero new quickstart
cd quickstart
jzero gen
go mod tidy
go run main.go server
```

## gateway

```shell
jzero new simplegateway --branch gateway
cd simplegateway
jzero gen
go mod tidy
go run main.go server
```

## api

```shell
jzero new simpleapi --branch api
cd simpleapi
jzero gen
go mod tidy
go run main.go server
```

## api-goctl

```shell
jzero new simpleapi-goctl --branch api
cd simpleapi-goctl
jzero gen
go mod tidy
go run main.go server
```

## rpc

```shell
jzero new simplerpc --branch rpc
cd simplerpc
jzero gen
go mod tidy
go run main.go server
```

## rpc-goctl

与原生 goctl 的 rpc 模板保持一致

```shell
jzero new simplerpc-goctl --branch rpc-goctl
cd simplerpc-goctl
jzero gen
go mod tidy
go run main.go server
```


## cli

```shell
jzero new simplecli --branch cli
cd simplecli
go mod tidy
go run main.go server
```
