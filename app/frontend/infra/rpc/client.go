package rpc

import (
	"sync"

	"github.com/Rogue-Trader-zzy/gomall/app/frontend/conf"
	frontendUtils "github.com/Rogue-Trader-zzy/gomall/app/frontend/utils"
	userservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
