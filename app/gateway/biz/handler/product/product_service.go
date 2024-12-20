// Code generated by hertz generator.

package product

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jichenssg/tikmall/app/common/client"
	common "github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/common"
	product "github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/product"
	"github.com/jichenssg/tikmall/app/gateway/utils"
	productrpc "github.com/jichenssg/tikmall/gen/kitex_gen/product"
)

// ProductList .
// @router /product/list [GET]
func ProductList(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("ListProducts service")

	var err error
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "ListProducts service page: %v, pageSize: %v, catrgoryId: %v", req.Page, req.PageSize, req.CategoryId)

	productclient := client.ProductClient
	products, err := productclient.ListProducts(ctx, &productrpc.ListProductsReq{
		Page:       req.Page,
		PageSize:   req.PageSize,
		CategoryId: req.CategoryId,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	productList := make([]*product.Product, 0, len(products.Products))
	for _, p := range products.Products {
		productList = append(productList, utils.Product2resp(p))
	}

	resp := &product.ListProductsResp{
		Message:  "list products success",
		Products: productList,
	}

	c.JSON(consts.StatusOK, resp)
}

// ProductGet .
// @router /product/get [GET]
func ProductGet(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("GetProduct service")

	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "GetProduct service id: %v", req.Id)

	productclient := client.ProductClient
	p, err := productclient.GetProduct(ctx, &productrpc.GetProductReq{
		Id: req.Id,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &product.GetProductResp{
		Message: "get product success",
		Product: utils.Product2resp(p.Product),
	}

	c.JSON(consts.StatusOK, resp)
}

// ProductSearch .
// @router /product/search [GET]
func ProductSearch(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("SearchProducts service")

	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "SearchProducts service query: %v", req.Query)

	productclient := client.ProductClient
	products, err := productclient.SearchProducts(ctx, &productrpc.SearchProductsReq{
		Query: req.Query,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	productList := make([]*product.Product, 0, len(products.Results))
	for _, p := range products.Results {
		productList = append(productList, utils.Product2resp(p))
	}

	resp := &product.SearchProductsResp{
		Message: "search products success",
		Results: productList,
	}

	c.JSON(consts.StatusOK, resp)
}

// ProductCreate .
// @router /product/create [POST]
func ProductCreate(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("CreateProduct service")

	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "CreateProduct service name: %v, description: %v, price: %v, categories: %v", req.Name, req.Description, req.Price, req.CategoryIds)

	productclient := client.ProductClient
	createresp, err := productclient.CreateProduct(ctx, &productrpc.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Picture:     req.Picture,
		CategoryIds: req.CategoryIds,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &product.CreateProductResp{
		Message: "create product success",
		Id:      createresp.Id,
	}

	c.JSON(consts.StatusOK, resp)
}

// ProductUpdate .
// @router /product/update [POST]
func ProductUpdate(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("UpdateProduct service")

	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "UpdateProduct service id: %v, name: %v, description: %v, price: %v, categories: %v", req.Id, req.Name, req.Description, req.Price, req.CategoryIds)

	productclient := client.ProductClient
	_, err = productclient.UpdateProduct(ctx, &productrpc.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Picture:     req.Picture,
		CategoryIds: req.CategoryIds,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &common.Response{
		Message: "update product success",
	}

	c.JSON(consts.StatusOK, resp)
}

// ProductDelete .
// @router /product/delete [POST]
func ProductDelete(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("DeleteProduct service")

	var err error
	var req product.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "DeleteProduct service id: %v", req.Id)

	productclient := client.ProductClient
	_, err = productclient.DeleteProduct(ctx, &productrpc.DeleteProductReq{
		Id: req.Id,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &common.Response{
		Message: "delete product success",
	}

	c.JSON(consts.StatusOK, resp)
}

// CategoryCreate .
// @router /category/create [POST]
func CategoryCreate(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("CreateCategory service")

	var err error
	var req product.CreateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "CreateCategory service name: %v", req.Name)

	productclient := client.ProductClient
	createresp, err := productclient.CreateCategory(ctx, &productrpc.CreateCategoryReq{
		Name: req.Name,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &product.CreateCategoryResp{
		Message: "create category success",
		Id:      createresp.Id,
	}

	c.JSON(consts.StatusOK, resp)
}

// CategoryList .
// @router /category/list [GET]
func CategoryList(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("ListCategories service")

	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	productclient := client.ProductClient
	categories, err := productclient.ListCategories(ctx, &productrpc.ListCategoriesReq{})
	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	categoryList := make([]*product.Category, 0, len(categories.Categories))
	for _, c := range categories.Categories {
		categoryList = append(categoryList, utils.Category2resp(c))
	}

	resp := &product.ListCategoriesResp{
		Message:    "list categories success",
		Categories: categoryList,
	}

	c.JSON(consts.StatusOK, resp)
}

// CategoryDelete .
// @router /category/delete [POST]
func CategoryDelete(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("DeleteCategory service")

	var err error
	var req product.DeleteCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "DeleteCategory service id: %v", req.Id)

	productclient := client.ProductClient
	_, err = productclient.DeleteCategory(ctx, &productrpc.DeleteCategoryReq{
		Id: req.Id,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &common.Response{
		Message: "delete category success",
	}

	c.JSON(consts.StatusOK, resp)
}

// CategoryUpdate .
// @router /category/update [POST]
func CategoryUpdate(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("UpdateCategory service")

	var err error
	var req product.UpdateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxDebugf(ctx, "UpdateCategory service id: %v, name: %v", req.Id, req.Name)

	productclient := client.ProductClient
	_, err = productclient.UpdateCategory(ctx, &productrpc.UpdateCategoryReq{
		Id:   req.Id,
		Name: req.Name,
	})

	if err != nil {
		c.JSON(utils.ParseRpcError(err))
		return
	}

	resp := &common.Response{
		Message: "update category success",
	}

	c.JSON(consts.StatusOK, resp)
}
