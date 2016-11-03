package grpc

import (
	"testing"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/mock"

	hello "github.com/micro/go-grpc/examples/greeter/server/proto/hello"

	"golang.org/x/net/context"
)

type testHandler struct{}

func (t *testHandler) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func TestGRPC(t *testing.T) {
	// create mock registry
	r := mock.NewRegistry()

	// create GRPC service
	service := NewService(
		micro.Name("test.service"),
		micro.Registry(r),
	)

	// register test handler
	hello.RegisterSayHandler(service.Server(), &testHandler{})

	// run service
	go func() {
		if err := service.Run(); err != nil {
			t.Fatal(err)
		}
	}()

	// artificial delay?
	time.Sleep(time.Millisecond * 10)

	// create client
	say := hello.NewSayClient("test.service", service.Client())

	// call service
	rsp, err := say.Hello(context.Background(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		t.Fatal(err)
	}

	// check message
	if rsp.Msg != "Hello John" {
		t.Fatal("unexpected response %s", rsp.Msg)
	}
}
