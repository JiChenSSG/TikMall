package utils

import (
	"github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/product"
	productrpc "github.com/jichenssg/tikmall/gen/kitex_gen/product"
)

func Product2resp(p *productrpc.Product) (res *product.Product) {
	res = &product.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Categories: func() (categories []*product.Category) {
			for _, c := range p.Categories {
				categories = append(categories, Category2resp(c))
			}
			return
		}(),
	}

	return
}

func Category2resp(c *productrpc.Category) (res *product.Category) {
	res = &product.Category{
		Id:   c.Id,
		Name: c.Name,
	}

	return
}
