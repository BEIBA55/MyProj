package models

type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	TeacherID *uint  `json:"teacher_id"`
}
