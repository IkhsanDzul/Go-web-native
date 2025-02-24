package entities

import "time"

type Product struct {
	Id        uint
	Name      string
	CategoryID  uint
    Category    Category `gorm:"foreignKey:CategoryID"`
	Stock     int64
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}

