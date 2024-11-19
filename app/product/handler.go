package main

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	"github.com/jichenssg/tikmall/app/product/dal/model"
	product "github.com/jichenssg/tikmall/gen/kitex_gen/product"
	"gorm.io/gorm"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	klog.Infof("ListProducts service")

	page := int(req.Page)
	pageSize := int(req.PageSize)
	catrgoryId := req.CategoryId

	klog.Debugf("ListProducts service page: %v, pageSize: %v, catrgoryId: %v", page, pageSize, catrgoryId)

	products, err := model.ListProducts(mysql.GetDB(), ctx, page, pageSize, catrgoryId)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		err = kerrors.NewBizStatusError(500, "list products failed")
		return
	}

	productList := make([]*product.Product, 0, len(products))
	for _, p := range products {
		productList = append(productList, product2resp(&p))
	}

	resp = &product.ListProductsResp{
		Products: productList,
	}

	return
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	klog.Infof("GetProduct service")

	id := req.Id
	klog.Debugf("GetProduct service id: %v", id)

	p, err := model.GetProduct(mysql.GetDB(), ctx, id)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "product not exists")
		} else {
			err = kerrors.NewBizStatusError(500, "get product failed")
		}

		return
	}

	resp = &product.GetProductResp{
		Product: product2resp(p),
	}

	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	klog.Infof("SearchProducts service")

	keyword := req.Query
	klog.Debugf("SearchProducts service keyword: %v", keyword)

	products, err := model.SearchProducts(mysql.GetDB(), ctx, keyword)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		err = kerrors.NewBizStatusError(500, "search products failed")
	}

	productList := make([]*product.Product, 0, len(products))
	for _, p := range products {
		productList = append(productList, product2resp(&p))
	}

	resp = &product.SearchProductsResp{
		Results: productList,
	}

	return
}

// CreateProduct implements the ProductCatalogServiceImpl interface.

func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	klog.Infof("CreateProduct service")

	name := req.Name
	description := req.Description
	price := req.Price
	picture := req.Picture
	categoryIds := req.CategoryIds

	klog.Debugf("CreateProduct service name: %v, description: %v, price: %v, picture: %v, categoryIds: %v", name, description, price, picture, categoryIds)

	categories := make([]model.Category, 0, len(categoryIds))
	for _, id := range categoryIds {
		categories = append(categories, model.Category{ID: id})
	}

	p := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Picture:     picture,
		Categories:  categories,
	}

	err = model.CreateProduct(mysql.GetDB(), ctx, p)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "product already exists")
		} else {
			err = kerrors.NewBizStatusError(500, "create product failed")
		}
	}

	resp = &product.CreateProductResp{
		Id: p.ID,
	}

	return
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	klog.Infof("UpdateProduct service")

	id := req.Id
	name := req.Name
	description := req.Description
	price := req.Price
	picture := req.Picture

	categoryIds := req.CategoryIds

	klog.Debugf("UpdateProduct service id: %v, name: %v, description: %v, price: %v, picture: %v, categoryIds: %v", id, name, description, price, picture, categoryIds)

	categories := make([]model.Category, 0, len(categoryIds))
	for _, id := range categoryIds {
		categories = append(categories, model.Category{ID: id})
	}

	p := &model.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Picture:     picture,
		Categories:  categories,
	}

	err = model.UpdateProduct(mysql.GetDB(), ctx, p)

	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "product not exists")
		} else {
			err = kerrors.NewBizStatusError(500, "update product failed")
		}
	}

	// detect if product is deleted
	if p.DeletedAt.Valid {
		err = kerrors.NewBizStatusError(400, "product not exists")
	}

	resp = &product.UpdateProductResp{}

	return
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	klog.Infof("DeleteProduct service")

	id := req.Id
	klog.Debugf("DeleteProduct service id: %v", id)

	err = model.DeleteProduct(mysql.GetDB(), ctx, id)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "product not exists")
		} else {
			err = kerrors.NewBizStatusError(500, "delete product failed")
		}

		return
	}

	resp = &product.DeleteProductResp{}

	return
}

// CreateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	klog.Infof("CreateCategory service")

	name := req.Name
	klog.Debugf("CreateCategory service name: %v", name)

	category := &model.Category{
		Name: name,
	}

	err = model.CreateCategory(mysql.GetDB(), ctx, category)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "category already exists")
		} else {
			err = kerrors.NewBizStatusError(500, "create category failed")
		}

		return
	}

	resp = &product.CreateCategoryResp{
		Id: category.ID,
	}

	return
}

// ListCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListCategories(ctx context.Context, req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	klog.Infof("ListCategories service")

	categories, err := model.ListCategories(mysql.GetDB(), ctx)

	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)
		err = kerrors.NewBizStatusError(500, "list categories failed")
		return
	}

	productCategories := make([]*product.Category, 0, len(categories))
	for _, category := range categories {
		productCategories = append(productCategories, &product.Category{
			Id:   category.ID,
			Name: category.Name,
		})
	}

	resp = &product.ListCategoriesResp{
		Categories: productCategories,
	}

	return
}

// DeleteCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	klog.Infof("DeleteCategory service")

	id := req.Id
	klog.Debugf("DeleteCategory service id: %v", id)

	err = model.DeleteCategory(mysql.GetDB(), ctx, id)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "category not exists")
		} else {
			err = kerrors.NewBizStatusError(500, "delete category failed")
		}

		return
	}

	resp = &product.DeleteCategoryResp{}

	return
}

// UpdateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	klog.Infof("UpdateCategory service")

	id := req.Id
	name := req.Name
	klog.Debugf("UpdateCategory service id: %v, name: %v", id, name)

	category := &model.Category{
		ID:   id,
		Name: name,
	}

	err = model.UpdateCategory(mysql.GetDB(), ctx, category)
	if err != nil {
		klog.Errorf("UpdateCategory service err: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = kerrors.NewBizStatusError(400, "category not exists")
		} else {
			err = kerrors.NewBizStatusError(500, "update category failed")
		}
	}

	// detect if category is deleted
	if category.DeletedAt.Valid {
		err = kerrors.NewBizStatusError(400, "category not exists")
	}

	resp = &product.UpdateCategoryResp{}

	return
}

func product2resp(p *model.Product) (res *product.Product) {
	res = &product.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Picture:     p.Picture,
		Categories: func() (categories []*product.Category) {
			for _, c := range p.Categories {
				categories = append(categories, &product.Category{
					Id:   c.ID,
					Name: c.Name,
				})
			}

			return
		}(),
	}

	return
}
