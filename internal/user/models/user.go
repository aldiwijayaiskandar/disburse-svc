package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Name 	 string `json:"name"`
}