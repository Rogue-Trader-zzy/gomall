package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func middleware(next endpoint.Endpoint) endpoint.Endpoint {
	// Finish your middleware logic.
	return func(ctx context.Context, req, resp interface{}) (err error) {
		begin := time.Now()
		err = next(ctx, req, resp)
		fmt.Println(time.Since(begin))
		return err
	}
}
