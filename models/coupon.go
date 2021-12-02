package models

import "time"

type Coupon struct {
	CouponId     int    `json:"categoryId" gorm:"primaryKey"`
	CouponName   string `json:"CouponName"`
	PercentDisc  int
	MinimumPrice int
	ValidUntil   time.Time
}

func (c Coupon) TableName() string {
	return "hasycoffee.coupon"
}
