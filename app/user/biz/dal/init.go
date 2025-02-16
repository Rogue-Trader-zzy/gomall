package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/user/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
