package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/checkout/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
