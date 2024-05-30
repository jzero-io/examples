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

## grpc + grpc-gateway

```shell
jzero new simplegateway --branch gateway
cd simplegateway
jzero gen --home ~/.jzero/templates/gateway
go mod tidy
go run main.go server
```

## api

```shell
jzero new simpleapi--branch api
cd simpleapi
jzero gen --home ~/.jzero/templates/api
go mod tidy
go run main.go server
```

## rpc

```shell
jzero new simplerpc--branch rpc
cd simplerpc
jzero gen --home ~/.jzero/templates/rpc
go mod tidy
go run main.go server
```

## cli

```shell
jzero new simplerpc--branch cli
cd simplecli
go mod tidy
go run main.go server
```