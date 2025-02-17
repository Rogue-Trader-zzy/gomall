package utils

import "context"

func GetUserIDFromCtx(ctx context.Context) int32 {
	UserId := ctx.Value(SessionUserId)
	if UserId == nil {
		return 0
	}
	return UserId.(int32)
}
