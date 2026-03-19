package service

import (
	"context"
	"fmt"
	"time"

	v1 "mapps_gateway/api/viability"
	"mapps_gateway/internal/service/variables"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthService struct {
	v1.UnimplementedViabilityServer
	uptime time.Time
}

func NewHealthService() *HealthService {
	return &HealthService{
		uptime: time.Now(),
	}
}

func (s *HealthService) Health(context.Context, *emptypb.Empty) (*v1.HealthReply, error) {
	return &v1.HealthReply{
		GatewayStatus: fmt.Sprintf(variables.ServiceIsUp, "mApps Gateway"),
		GatewayUptime: s.uptime.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *HealthService) Ready(context.Context, *emptypb.Empty) (*v1.ReadinessReply, error) {
	// TODO: реализовать проверку downstream сервисов (auth, product, asset_manager)
	return &v1.ReadinessReply{
		Status:             "ok",
		AuthStatus:         "TODO: check auth service",
		ProductStatus:      "TODO: check product service",
		AssetManagerStatus: "TODO: check asset_manager service",
	}, nil
}
