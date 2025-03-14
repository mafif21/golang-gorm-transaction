package model

import "time"

type Order struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"index;not null"`
	ProductID int       `gorm:"index;not null"`
	Amount    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User    User    `gorm:"foreignKey:UserID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

func (o *Order) TableName() string {
	return "orders"
}
