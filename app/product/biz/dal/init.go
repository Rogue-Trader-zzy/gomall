package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/product/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
