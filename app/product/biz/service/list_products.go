package service

import (
	"context"

	"github.com/Rogue-Trader-zzy/gomall/app/product/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/product/biz/model"
	product "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategorynName)
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       float32(v.Price),
			})
		}
	}
	return resp, nil
}
