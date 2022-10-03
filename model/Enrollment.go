package model

type Enrollment struct {
	Enrolment_id uint    `json:"enrolment_id" gorm:"primary_key"`
	Final_grade   float32 `json:"final_grade"`
	Semester      string  `json:"semester"`
	User_id       string  `json:"user_id"`
	Group_id      uint    `json:"group_id"`
	Subject_id    uint    `json:"subject_id"`
}