package authGrpc

import (
	ssov1 "AlekseyPromet/authorization/internal/gen/ssov1/contracts"
	"AlekseyPromet/authorization/internal/models"
	"context"
	"net/mail"
	"unicode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *GrpcServerApiV1) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {

	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	email, err := mail.ParseAddress(req.GetEmail())
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

	userId, err := srv.storage.Register(ctx, models.User{
		Email:    email.String(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, "register failed")
	}

	resp := &ssov1.RegisterResponse{
		UserId: userId.String(),
	}

	return resp, nil
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
