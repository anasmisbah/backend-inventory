package models

type Product struct {
	Id uint			`json:"id"`
	Name string		`json:"name"`
	Photo string	`json:"photo"`
	Price int64		`json:"price"`
}
