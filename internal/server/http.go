package server

import (
	assetv1 "mapps_gateway/api/generated/proto/asset_manager/v1"
	authv1 "mapps_gateway/api/generated/proto/auth/v1"
	productv1 "mapps_gateway/api/generated/proto/products/v1"
	viabilityv1 "mapps_gateway/api/viability"
	"mapps_gateway/internal/conf"
	"mapps_gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// NewHTTPServer создаёт HTTP сервер.
func NewHTTPServer(
	c *conf.Server,
	healthChecker *service.HealthService,
	authService *service.AuthService,
	productService *service.ProductService,
	assetManagerService *service.AssetManagerService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
		http.ResponseEncoder(func(writer http.ResponseWriter, request *http.Request, a any) error {
			if m, ok := a.(interface{ proto.Message }); ok {
				mo := protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				}
				b, err := mo.Marshal(m)
				if err != nil {
					return err
				}
				writer.Header().Set("Content-Type", "application/json")
				_, err = writer.Write(b)
				return err
			}
			return http.DefaultResponseEncoder(writer, request, a)
		}),
	}
	if c.Http.GetNetwork() != "" {
		opts = append(opts, http.Network(c.Http.GetNetwork()))
	}
	if c.Http.GetAddr() != "" {
		opts = append(opts, http.Address(c.Http.GetAddr()))
	}

	srv := http.NewServer(opts...)

	_ = logger.Log(log.LevelInfo, "msg", "HTTP server initialized")

	viabilityv1.RegisterViabilityHTTPServer(srv, healthChecker)
	authv1.RegisterAuthHTTPServer(srv, authService)
	productv1.RegisterProductsHTTPServer(srv, productService)
	assetv1.RegisterAssetManagerHTTPServer(srv, assetManagerService)

	return srv
}
