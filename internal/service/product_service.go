package service

import (
	"context"
	"fmt"
	"mapps_gateway/internal/data"
	"mapps_gateway/internal/service/variables"

	productv1 "mapps_gateway/api/generated/proto/products/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// ProductService — прокси-сервис для mapps/product_service.
type ProductService struct {
	productv1.UnimplementedProductsServer
	data   *data.Data
	logger log.Logger
}

func NewProductService(d *data.Data, logger log.Logger) *ProductService {
	return &ProductService{
		data:   d,
		logger: logger,
	}
}

func (s *ProductService) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	if s.data.ProductClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "product")
	}
	return s.data.ProductClient.GetProduct(ctx, req)
}

func (s *ProductService) ListProducts(ctx context.Context, req *productv1.ListProductsRequest) (*productv1.ListProductsResponse, error) {
	if s.data.ProductClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "product")
	}
	return s.data.ProductClient.ListProducts(ctx, req)
}

func (s *ProductService) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	if s.data.ProductClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "product")
	}
	return s.data.ProductClient.CreateProduct(ctx, req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (*productv1.UpdateProductResponse, error) {
	if s.data.ProductClient == nil {
		return nil, fmt.Errorf(variables.ServiceIsDown, "product")
	}
	return s.data.ProductClient.UpdateProduct(ctx, req)
}
