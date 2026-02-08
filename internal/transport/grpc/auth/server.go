package auth

import (
	"context"
	"proto-auth/internal/domain/auth"

	ssov1 "github.com/belskirill/proto-api-auth/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func RegisterAuthAPIServer(srv *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(srv, &authAPI{
		auth: auth,
	})
}

func (a *authAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := a.auth.Login(ctx, auth.Login{
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
	}

	return &ssov1.LoginResponse{}, nil
}

func (a *authAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}
