package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/order/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
