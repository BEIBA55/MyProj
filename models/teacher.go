package models

type Teacher struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	SubjectID uint    `json:"subject_id"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
}
