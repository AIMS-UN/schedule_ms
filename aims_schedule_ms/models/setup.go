package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	_ "github.com/lib/pq"
	//"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	// Enable VIPER to read Environment Variables
	//viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env
	// viper_user := os.Getenv("POSTGRES_USER")
	// viper_password := os.Getenv("POSTGRES_PASSWORD")
	// viper_db := os.Getenv("POSTGRES_DB")
	// viper_host := os.Getenv("POSTGRES_HOST")
	// viper_port := os.Getenv("POSTGRES_PORT")
	viper_user := "postgres"
	viper_password := "postgres"
	viper_db := "aims_enrollment_db"
	viper_host := "localhost"
	viper_port := 5434
	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Enrrollment{})
	// Initialise value
	m := Enrrollment{Enrrolment_id: 1, Final_grade: 3.6, Semester: "2021-1", User_id: "U2342UDTdsa3254sfs3423", Gruop_id: 3, Subject_id: 56}
	db.Create(&m)
	return db
}
