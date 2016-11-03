# Go GRPC [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-grpc?status.svg)](https://godoc.org/github.com/micro/go-grpc) [![Travis CI](https://api.travis-ci.org/micro/go-grpc.svg?branch=master)](https://travis-ci.org/micro/go-grpc) [![Go Report Card](https://goreportcard.com/badge/micro/go-grpc)](https://goreportcard.com/report/github.com/micro/go-grpc)

Go GRPC is a go-micro compatible service which leverages gRPC as the client, server and transport. Go GRPC shares the go-micro code base, making it 
a pluggable GRPC framework for microservices. Everything is the same as go-micro except for the rpc, transport and encoding. It still works with 
micro generated protobufs and defaults to consul for service discovery.

Find an example greeter service in [examples/greeter](https://github.com/micro/go-grpc/tree/master/examples/greeter).
