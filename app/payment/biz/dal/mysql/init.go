package mysql

import (
	"fmt"
	"os"

	"github.com/Rogue-Trader-zzy/gomall/app/payment/biz/model"
	"github.com/Rogue-Trader-zzy/gomall/app/payment/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		if err := DB.AutoMigrate(&model.PaymentLog{}); err != nil {
			klog.Error(err)
		}
	}
}
