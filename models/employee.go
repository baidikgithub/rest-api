package models

type Employee struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
