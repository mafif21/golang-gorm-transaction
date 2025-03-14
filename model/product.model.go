package model

type Product struct {
	ID     int     `gorm:"primaryKey"`
	Name   string  `gorm:"type:varchar(100);not null"`
	Amount int     `gorm:"not null"`
	Price  float64 `gorm:"type:decimal(10,2);not null"`
}

func (p *Product) TableName() string {
	return "products"
}
