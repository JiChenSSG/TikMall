package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	product "github.com/jichenssg/tikmall/gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	klog.Infof("ListProducts service")

	return
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// TODO: Your code here...
	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// TODO: Your code here...
	return
}

// CreateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// TODO: Your code here...
	return
}

// CreateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	// TODO: Your code here...
	return
}

// ListCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListCategories(ctx context.Context, req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	// TODO: Your code here...
	return
}
