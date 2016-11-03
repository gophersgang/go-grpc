package grpc

import (
	"github.com/micro/go-micro"

	client "github.com/micro/go-plugins/client/grpc"
	server "github.com/micro/go-plugins/server/grpc"
)

// NewService returns a grpc service compatible with go-micro.Service
func NewService(opts ...micro.Option) micro.Service {
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return micro.NewService(options...)
}
