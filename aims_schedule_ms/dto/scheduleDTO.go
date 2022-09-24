package dto

type Scheduledto struct {
	Enrrolment_id uint    `json:"enrrolment_id" gorm:"primary_key"`
	Final_grade   float32 `json:"final_grade"`
	Semester      string  `json:"semester"`
	User_id       string  `json:"user_id"`
	Gruop_id      uint    `json:"gruop_id"`
	Subject_id    uint    `json:"subject_id"`
}
