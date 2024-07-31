package schema

type User struct {
	UserId string `json:"userId" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null"`
}
