package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/email/biz/dal/mysql"
	"github.com/Rogue-Trader-zzy/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
