package model

type User struct {
	ID     int     `gorm:"primaryKey"`
	Name   string  `gorm:"type:varchar(100);not null"`
	Orders []Order `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *User) TableName() string {
	return "users"
}
