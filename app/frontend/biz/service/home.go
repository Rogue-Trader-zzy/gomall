package service

import (
	"context"

	common "github.com/Rogue-Trader-zzy/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	resp := make(map[string]any)
	items := []map[string]any{
		{"Name": "饼图", "Price": 100, "Picture": "/static/image/饼图.png"},
		{"Name": "气泡图", "Price": 110, "Picture": "/static/image/气泡图.png"},
		{"Name": "条形图", "Price": 120, "Picture": "/static/image/条形图.png"},
		{"Name": "相关性图", "Price": 130, "Picture": "/static/image/相关性图.png"},
		{"Name": "直方图", "Price": 140, "Picture": "/static/image/直方图.png"},
		{"Name": "饼图", "Price": 150, "Picture": "/static/image/饼图.png"},
	}
	resp["Title"] = "BioCoder"
	resp["Items"] = items
	return resp, nil
}
