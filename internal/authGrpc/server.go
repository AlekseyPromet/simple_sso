package authGrpc

import (
	ssov1 "AlekseyPromet/authorization/internal/gen/v1/contracts"
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServerApiV1 struct {
	ssov1.UnimplementedAuthServer
	lisner net.Listener
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &GrpcServerApiV1{})
}

func (ssov1 *GrpcServerApiV1) Register(ctx context.Context, req *ssov1.RegisterRequest) (resp *ssov1.RegisterResponse, err error) {
	return resp, nil
}

func (ssov1 *GrpcServerApiV1) Logining(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	return nil, nil
}

func (ssov1 *GrpcServerApiV1) Run(cfg models.GrpcConfig) error {
	const op = "authGrpc.Run"

	listner, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("%s %v", op, err)
	}
	ssov1.lisner = listner

	return nil
}

func (ssov1 *GrpcServerApiV1) Shutdown(ctx context.Context) error {
	return ssov1.lisner.Close()
}
