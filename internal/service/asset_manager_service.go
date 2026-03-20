package service

import (
	"context"
	"fmt"
	"mapps_gateway/internal/data"
	"mapps_gateway/internal/service/variables"

	assetv1 "mapps_gateway/api/generated/proto/asset_manager/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// AssetManagerService — прокси-сервис для mapps/asset_manager.
type AssetManagerService struct {
	assetv1.UnimplementedAssetManagerServer
	data   *data.Data
	logger log.Logger
}

func NewAssetManagerService(d *data.Data, logger log.Logger) *AssetManagerService {
	return &AssetManagerService{
		data:   d,
		logger: logger,
	}
}

func (s *AssetManagerService) GetModelUploadURL(ctx context.Context, req *assetv1.GetModelUploadURLRequest) (*assetv1.GetModelUploadURLResponse, error) {
	if s.data.AssetManagerClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "asset_manager")
	}
	return s.data.AssetManagerClient.GetModelUploadURL(ctx, req)
}

func (s *AssetManagerService) GetModel(ctx context.Context, req *assetv1.GetModelRequest) (*assetv1.GetModelResponse, error) {
	if s.data.AssetManagerClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "asset_manager")
	}
	return s.data.AssetManagerClient.GetModel(ctx, req)
}

func (s *AssetManagerService) GetAssetUploadURL(ctx context.Context, req *assetv1.GetAssetUploadURLRequest) (*assetv1.GetAssetUploadURLResponse, error) {
	if s.data.AssetManagerClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "asset_manager")
	}
	return s.data.AssetManagerClient.GetAssetUploadURL(ctx, req)
}

func (s *AssetManagerService) GetAsset(ctx context.Context, req *assetv1.GetAssetRequest) (*assetv1.GetAssetResponse, error) {
	if s.data.AssetManagerClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "asset_manager")
	}
	return s.data.AssetManagerClient.GetAsset(ctx, req)
}
