package dal

import (
	"github.com/Rogue-Trader-zzy/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
