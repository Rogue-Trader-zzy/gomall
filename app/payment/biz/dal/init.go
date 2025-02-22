package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
