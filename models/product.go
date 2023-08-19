package models

type Product struct {
	Id uint			`gorm:"primaryKey" form:"id"`
	Name string		`form:"name" validate:"required"`
	Photo string	`form:"photo" validate:"required"`
	Price int64		`form:"price" validate:"required"`
}
