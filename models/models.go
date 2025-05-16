package models

import "gorm.io/gorm"

// Product Model (Represents the 'products' table in MySQL)
type Product struct {
	gorm.Model
	Name     string  `gorm:"unique;not null" json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	OwnerID  uint    `json:"owner_id"`
}

// Order Model (Represents the 'orders' table in MySQL)
type Order struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
