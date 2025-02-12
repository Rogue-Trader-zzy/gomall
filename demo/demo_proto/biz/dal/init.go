package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/demo/demo_proto/biz/dal/mysql"
	// "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
