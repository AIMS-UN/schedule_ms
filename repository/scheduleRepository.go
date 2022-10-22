package repository

import (
	model "aims_schedule_ms/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	//connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	fmt.Println("conname is\t\t", connStr)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&model.Enrollment{})

	// check if migration is successful
	// if db.HasTable(&model.Enrollment{}) {
	// 	fmt.Println("Enrollment table created")
	// } else {
	// 	fmt.Println("Enrollment table not created")
	// }

	// // Initialise value
	// m := model.Enrollment{Enrollment_id: 1, Final_grade: 3.6, Semester: "2021-1", User_id: "U2342UDTdsa3254sfs3423", Group_id: "3", Subject_id: "56"}
	// db.Create(&m)
	return db
}
