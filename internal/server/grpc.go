package server

import (
	assetv1 "mapps_gateway/api/generated/proto/asset_manager/v1"
	authv1 "mapps_gateway/api/generated/proto/auth/v1"
	productv1 "mapps_gateway/api/generated/proto/products/v1"
	viabilityv1 "mapps_gateway/api/viability"
	"mapps_gateway/internal/conf"
	"mapps_gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer создаёт gRPC сервер.
func NewGRPCServer(
	c *conf.Server,
	health *service.HealthService,
	authService *service.AuthService,
	productService *service.ProductService,
	assetManagerService *service.AssetManagerService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.GetNetwork() != "" {
		opts = append(opts, grpc.Network(c.Grpc.GetNetwork()))
	}
	if c.Grpc.GetAddr() != "" {
		opts = append(opts, grpc.Address(c.Grpc.GetAddr()))
	}
	srv := grpc.NewServer(opts...)

	viabilityv1.RegisterViabilityServer(srv, health)
	authv1.RegisterAuthServer(srv, authService)
	productv1.RegisterProductsServer(srv, productService)
	assetv1.RegisterAssetManagerServer(srv, assetManagerService)

	return srv
}
