package authGrpc

import (
	ssov1 "AlekseyPromet/authorization/internal/gen/ssov1/contracts"
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"
	"net"
	"net/mail"
	"unicode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

type GrpcServerApiV1 struct {
	ssov1.UnimplementedAuthServer
	lisner net.Listener
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &GrpcServerApiV1{})
}

func (srv *GrpcServerApiV1) Register(ctx context.Context, req *ssov1.RegisterRequest) (resp *ssov1.RegisterResponse, err error) {
	return resp, nil
}

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

	sevenOrMore, number, upper, special := verifyPassword(req.GetPassword())
	if !sevenOrMore {
		return nil, status.Error(codes.InvalidArgument, "password is short,it length must be 7 or large ")
	}
	if !number {
		return nil, status.Error(codes.InvalidArgument, "password not correct, it is not conteined number")
	}
	if !upper {
		return nil, status.Error(codes.InvalidArgument, "password not correct, it is not conteined upper case character")
	}
	if !special {
		return nil, status.Error(codes.InvalidArgument, "password not correct, it is not conteined special character")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app id is required")
	}

	resp := &ssov1.LoginResponse{
		Token: "",
	}
	return resp, nil
}

func (srv *GrpcServerApiV1) Run(cfg models.GrpcConfig) error {
	const op = "authGrpc.Run"

	listner, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("%s %v", op, err)
	}
	srv.lisner = listner

	return nil
}

func (srv *GrpcServerApiV1) Shutdown(ctx context.Context) error {
	return srv.lisner.Close()
}

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	sevenOrMore = letters >= 7
	return
}
