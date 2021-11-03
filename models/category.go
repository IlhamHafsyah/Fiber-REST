package models

type Category struct {
	CategoryId   int    `json:"categoryId" gorm:"primaryKey"`
	CategoryDesc string `json:"categoryDesc"`
}

func (c Category) TableName() string {
	return "hasycoffee.category"
}
