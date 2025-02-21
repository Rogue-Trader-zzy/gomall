package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/payment/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
