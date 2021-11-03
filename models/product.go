package models

import "time"

type Product struct {
	ProductID  uint      `json:"ProductId" gorm:"primaryKey"` // productId
	Name       string    `json:"Name"`                        // name
	Price      int64     // price
	AvailHours time.Time // availHours
	EndHours   time.Time // availHours
	IDCategory int64     // idCategory
}

func (c Product) TableName() string {
	return "hasycoffee.product"
}
