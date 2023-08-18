package models

type Stock struct {
	Id 			uint		`json:"id"`
	Amount 		int64		`json:"amount"`
	UpdatedAt 	string		`json:"updated_at"`
	ProductId 	int64		
}