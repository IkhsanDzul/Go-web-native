package entities

import "time"

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Category) Create(category Category) any {
	panic("unimplemented")
}
