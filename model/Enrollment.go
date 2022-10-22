package model

type Enrollment struct {
	Enrollment_id string  `json:"enrollment_id" gorm:"primary_key"`
	Final_grade   float32 `json:"final_grade"`
	Semester      string  `json:"semester"`
	User_id       string  `json:"user_id"`
	Group_id      string  `json:"group_id"`
	Subject_id    string  `json:"subject_id"`
}
