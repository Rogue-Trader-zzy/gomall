package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11);not null"`
	OrderIdRefer string  `gorm:"type:varchar(100);index"`
	Quantity     uint32  `gorm:"type:int(11);not null"`
	Cost         float32 `gorm:"type:decimal (10,2);not null"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
