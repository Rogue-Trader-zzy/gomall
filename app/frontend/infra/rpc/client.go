package rpc

import (
	"log"
	"sync"

	userservice "github.com/Rogue-Trader-zzy/gomall/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	userClient = userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {})
}

func iniUserClient(client userservice.UserService) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	client, err = userservice.NewClient("greet.server", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
