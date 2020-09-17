package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	PoructName  string  `gorm:"column:productname" json:"productname"`
	Description string
}

func (Product) TableName() string {
	return "products"
}
