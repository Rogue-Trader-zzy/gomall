package service

import (
	"context"
	"testing"

	"github.com/Rogue-Trader-zzy/gomall/app/user/biz/dal/mysql"
	user "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "dem@admin.com",
		Password:        "12456",
		PasswordConfirm: "12456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
