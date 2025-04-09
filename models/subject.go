package models

type Subject struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
