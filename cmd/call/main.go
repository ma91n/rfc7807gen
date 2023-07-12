package main

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/ma91n/rfc7807gen/client"
	"github.com/ma91n/rfc7807gen/client/example"
)

func main() {

	rt := httptransport.New(
		"localhost:8081",
		client.DefaultBasePath,
		client.DefaultSchemes)
	rt.Consumers["application/problem+json"] = runtime.JSONConsumer()

	exampleClient := client.New(rt, nil)

	resp, err := exampleClient.Example.Hello((&example.HelloParams{}).WithContext(context.Background()))
	if err != nil {
		fmt.Println("!!!", err)
		return
	}
	payload := resp.GetPayload()
	fmt.Println(payload.Message)

}
