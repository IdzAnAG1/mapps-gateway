package service

import (
	"context"
	"fmt"
	"mapps_gateway/internal/data"
	"mapps_gateway/internal/service/variables"

	authv1 "mapps_gateway/api/generated/proto/auth/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// AuthService — прокси-сервис для mapps/auth.
type AuthService struct {
	authv1.UnimplementedAuthServer
	data   *data.Data
	logger log.Logger
}

func NewAuthService(d *data.Data, logger log.Logger) *AuthService {
	return &AuthService{
		data:   d,
		logger: logger,
	}
}

func (s *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	if s.data.AuthClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "auth")
	}
	return s.data.AuthClient.Register(ctx, req)
}

func (s *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	if s.data.AuthClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "auth")
	}
	return s.data.AuthClient.Login(ctx, req)
}
