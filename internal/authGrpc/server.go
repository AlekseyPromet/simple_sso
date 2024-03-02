package authGrpc

import (
	ssov1 "AlekseyPromet/simple_sso/internal/gen/ssov1/contracts"
	"AlekseyPromet/simple_sso/internal/models"
	"net"

	"google.golang.org/grpc"
)

const emptyValue = 0

type GrpcServerApiV1 struct {
	ssov1.UnimplementedAuthServer
	lisner  net.Listener
	storage models.IStorage
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &GrpcServerApiV1{})
}
