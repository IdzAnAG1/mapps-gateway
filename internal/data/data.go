package data

import (
	"fmt"
	"mapps_gateway/internal/conf"

	assetv1 "mapps_gateway/api/generated/proto/asset_manager/v1"
	authv1 "mapps_gateway/api/generated/proto/auth/v1"
	productv1 "mapps_gateway/api/generated/proto/products/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data содержит typed gRPC клиенты для downstream сервисов.
type Data struct {
	AuthClient         authv1.AuthClient
	ProductClient      productv1.ProductsClient
	AssetManagerClient assetv1.AssetManagerClient
}

// NewData создаёт gRPC соединения с downstream сервисами.
func NewData(c *conf.Data) (*Data, func(), error) {
	var (
		authConn    *grpc.ClientConn
		productConn *grpc.ClientConn
		assetConn   *grpc.ClientConn
		err         error
	)

	if c.Auth != nil && c.Auth.Addr != "" {
		authConn, err = grpc.NewClient(
			c.Auth.GetAddr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect to auth service: %w", err)
		}
	}

	if c.Product != nil && c.Product.Addr != "" {
		productConn, err = grpc.NewClient(
			c.Product.GetAddr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect to product service: %w", err)
		}
	}

	if c.AssetManager != nil && c.AssetManager.Addr != "" {
		assetConn, err = grpc.NewClient(
			c.AssetManager.GetAddr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect to asset_manager service: %w", err)
		}
	}

	cleanup := func() {
		closeConn := func(conn *grpc.ClientConn, name string) {
			if conn == nil {
				return
			}
			if err := conn.Close(); err != nil {
				log.Error(fmt.Sprintf("failed to close %s gRPC client: %v", name, err))
			}
		}
		closeConn(authConn, "auth")
		closeConn(productConn, "product")
		closeConn(assetConn, "asset_manager")
		log.Info("data resources closed")
	}

	d := &Data{}
	if authConn != nil {
		d.AuthClient = authv1.NewAuthClient(authConn)
	}
	if productConn != nil {
		d.ProductClient = productv1.NewProductsClient(productConn)
	}
	if assetConn != nil {
		d.AssetManagerClient = assetv1.NewAssetManagerClient(assetConn)
	}

	return d, cleanup, nil
}
