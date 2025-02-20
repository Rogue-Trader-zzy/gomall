package service

import (
	"context"

	cart "github.com/Rogue-Trader-zzy/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/Rogue-Trader-zzy/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/Rogue-Trader-zzy/gomall/app/frontend/infra/rpc"
	rpccart "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpc.CartClient.AddItem(h.Context, &rpccart.AddCartItemReq{
		ProductId: req.ProductId,
	})
	return
}
