package main

import (
	"context"
	"fmt"
	quickstart "quickstart-go"
	"quickstart-go/model/quickstart/pb/hellopb"
	"quickstart-go/rest"
)

func main() {
	clientset, err := quickstart.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithAddr("127.0.0.1"),
		rest.WithPort("8001"),
	)
	if err != nil {
		panic(err)
	}

	response, err := clientset.Hello().SayHello(context.Background(), &hellopb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
