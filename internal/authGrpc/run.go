package authGrpc

import (
	"AlekseyPromet/simple_sso/internal/models"
	"context"
	"fmt"
	"net"
)

func (srv *GrpcServerApiV1) Run(cfg models.GrpcConfig, storage models.IStorage) error {
	const op = "authGrpc.Run"

	listner, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("%s %v", op, err)
	}
	srv.lisner = listner
	srv.storage = storage

	return nil
}

func (srv *GrpcServerApiV1) Shutdown(ctx context.Context) error {
	return srv.lisner.Close()
}
