package models

type Wallet struct {
	UserId       string  `json:"userId" gorm:"primaryKey"`
	Balance 	 float64 `json:"balance"`
}