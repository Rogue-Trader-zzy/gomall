package rpc

import (
	"sync"

	"github.com/Rogue-Trader-zzy/gomall/app/frontend/conf"
	frontendUtils "github.com/Rogue-Trader-zzy/gomall/app/frontend/utils"
	"github.com/Rogue-Trader-zzy/gomall/common/clientsuite"
	cartservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/cart/cartservice"
	checkoutservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	orderservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/order/orderservice"
	productcatalogservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	userservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	OrderClient    orderservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	ServiceName    = frontendUtils.ServiceName
	err            error
)

func Init() {
	once.Do(func() {
		iniUserClient()
		iniProductClient()
		iniCartClient()
		iniCheckoutClient()
		iniOrderClient()
	})
}

func iniUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
