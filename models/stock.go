package models

type Stock struct {
	Id 			uint		`gorm:"primaryKey" form:"id"`
	Amount 		int64		`form:"amount" validate:"required"`
	UpdatedAt 	int			`gorm:"autoUpdateTime"` 
	ProductId 	uint		`form:"product_id" validate:"required"`
}