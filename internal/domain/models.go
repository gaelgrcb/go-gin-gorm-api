package domain

import "time"

type Category struct {
	ID          int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime:true" json:"created_at"`
}

type Product struct {
	ID         int64     `gorm:"type:bigint;primaryKey" json:"id"`
	SKU        string    `gorm:"type:varchar(255);unique;not null" json:"sku"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Price      int64     `gorm:"type:bigint;not null" json:"price"`
	Stock      int       `gorm:"type:int;not null;check:stock >= 0" json:"stock"`
	CategoryID int64     `gorm:"type:bigint;not null" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type StockMovement struct {
	ID        int64     `gorm:"type:bigint;primaryKey" json:"id"`
	ProductID int64     `gorm:"type:bigint;not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Type      string    `gorm:"type:varchar(255);not null" json:"type"`
	Quantity  int       `gorm:"type:int;not null" json:"quantity"`
	Reason    string    `gorm:"type:varchar(255);not null" json:"reason"`
	Date      time.Time `gorm:"autoCreateTime:true" json:"date"`
}
