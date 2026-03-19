package service

import (
	"context"
	"mapps_gateway/internal/data"

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
	return s.data.ProductClient.GetProduct(ctx, req)
}

func (s *ProductService) ListProducts(ctx context.Context, req *productv1.ListProductsRequest) (*productv1.ListProductsResponse, error) {
	return s.data.ProductClient.ListProducts(ctx, req)
}

func (s *ProductService) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	return s.data.ProductClient.CreateProduct(ctx, req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (*productv1.UpdateProductResponse, error) {
	return s.data.ProductClient.UpdateProduct(ctx, req)
}
