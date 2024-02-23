package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/contract"
	"github.com/zeddy-go/zeddy/errx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

func NewModule() contract.IModule {
	return &Module{}
}

type Module struct {
	grpcServer *grpc.Server
}

func (m *Module) Name() string {
	return "grpc"
}

func (m *Module) Init() (err error) {
	m.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(simpleInterceptor),
	)

	container.Invoke(func(c *viper.Viper) {
		if c.GetBool("grpc.reflection") {
			reflection.Register(m.grpcServer)
		}
	})

	err = container.Bind[*grpc.Server](m.grpcServer, container.AsSingleton())
	if err != nil {
		return
	}

	return
}

func (m *Module) Start() {
	var lis net.Listener
	err := container.Invoke(func(c *viper.Viper) (err error) {
		lis, err = net.Listen("tcp", fmt.Sprintf(":%d", c.GetInt("grpc.port")))
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		panic(errx.Wrap(err, "tcp listen port failed"))
	}

	err = m.grpcServer.Serve(lis)
	if err != nil {
		slog.Error("grpc server shutdown", "error", err)
	}
}
