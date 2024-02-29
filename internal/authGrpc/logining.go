package authGrpc

import (
	ssov1 "AlekseyPromet/authorization/internal/gen/ssov1/contracts"
	"context"
	"net/mail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *GrpcServerApiV1) Logining(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {

	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	_, err := mail.ParseAddress(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "it is'n valid email")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app id is required")
	}

	resp := &ssov1.LoginResponse{
		Token: "",
	}
	return resp, nil
}
