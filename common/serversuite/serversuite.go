package serversuite

import (
	"github.com/Rogue-Trader-zzy/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/remote/transmeta"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

type CommonServerSuite struct {
	CurrentServerName string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServerName,
		}),
		server.WithTracer(prometheus.NewServerTracer("",
			"",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry)),
		),
	}
	return opts
}
