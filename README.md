# Go GRPC [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-grpc?status.svg)](https://godoc.org/github.com/micro/go-grpc) [![Travis CI](https://api.travis-ci.org/micro/go-grpc.svg?branch=master)](https://travis-ci.org/micro/go-grpc) [![Go Report Card](https://goreportcard.com/badge/micro/go-grpc)](https://goreportcard.com/report/github.com/micro/go-grpc)

Go GRPC is a go-micro compatible library which leverages gRPC as the client, server and transport. Go GRPC shares the go-micro code base, making it 
a pluggable GRPC framework for microservices. Everything is the same as go-micro except for the rpc, transport and encoding. It still works with 
micro generated protobufs and defaults to consul for service discovery.

Find an example greeter service in [examples/greeter](https://github.com/micro/go-grpc/tree/master/examples/greeter).

## Write services identical to Micro

Initialisation of a go-grpc service is identical to a go-micro service. Which means you can swap out `micro.NewService` for `grpc.NewService` 
with zero other code changes. 

```go
package main

import (
	"log"
	"time"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	hello "github.com/micro/go-grpc/examples/greeter/server/proto/hello"

	"golang.org/x/net/context"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```

## Use client/server directly

This library is actually only 30 lines of code that simplifies initialisation of a go-micro.Service. 

You can do the same yourself.

```go
package main

import (
	"log"

	"github.com/micro/go-micro"
	client "github.com/micro/go-plugins/client/grpc"
	server "github.com/micro/go-plugins/server/grpc"

	hello "github.com/micro/go-grpc/examples/greeter/server/proto/hello"

	"golang.org/x/net/context"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// new micro service with grpc client/server
	service := micro.NewService(
		micro.Client(client.NewClient()),
		micro.Server(server.NewServer()),
		micro.Name("go.micro.srv.greeter"),
	)

	// parse command line flags
	service.Init()

	// register handler
	hello.RegisterSayHandler(service.Server(), new(Say))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```
