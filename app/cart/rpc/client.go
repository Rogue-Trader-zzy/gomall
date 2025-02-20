package rpc

import (
	"sync"

	"github.com/Rogue-Trader-zzy/gomall/app/cart/conf"
	cartUtils "github.com/Rogue-Trader-zzy/gomall/app/cart/utils"
	productcatalogservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		iniProductClient()
	})
}

func iniProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
