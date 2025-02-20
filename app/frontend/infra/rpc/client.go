package rpc

import (
	"sync"

	"github.com/Rogue-Trader-zzy/gomall/app/frontend/conf"
	frontendUtils "github.com/Rogue-Trader-zzy/gomall/app/frontend/utils"
	cartservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/cart/cartservice"
	productcatalogservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	userservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
		iniProductClient()
		iniCartClient()
	})
}

func iniUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func iniProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func iniCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}
