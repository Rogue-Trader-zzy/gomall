package rpc

import (
	"sync"

	"github.com/Rogue-Trader-zzy/gomall/app/cart/conf"
	cartUtils "github.com/Rogue-Trader-zzy/gomall/app/cart/utils"
	clientsuite "github.com/Rogue-Trader-zzy/gomall/common/clientsuite"
	productcatalogservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	ServiceName   = conf.GetConf().Kitex.Service
	err           error
)

func InitClient() {
	once.Do(func() {
		iniProductClient()
	})
}

func iniProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
