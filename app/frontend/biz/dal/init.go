package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/frontend/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
