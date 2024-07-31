package schema

type User struct {
	Id   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
