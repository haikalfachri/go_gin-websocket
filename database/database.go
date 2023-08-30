package database

import (
	"errors"
	"fmt"
	"go-gin-student/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}
  
func (c *Config) ConnectDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  c.DB_USERNAME,
	  c.DB_PASSWORD,
	  c.DB_HOST,
	  c.DB_PORT,
	  c.DB_NAME,
	)
	
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
	  panic(err)
	}

	log.Printf("successfully connected to database\n")
}

func MigrateDB() {

	err := DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

	log.Printf("successfully database migration\n")
}

func SeedStudent() (models.Student, error) {
	var student models.Student = models.Student{
		FirstName: "Haikal",
		LastName: "Fachri",
		StudentID: "123456789",
		Major: "Computer Science",
	}

	result := DB.Create(student)

	if err := result.Error; err != nil {
		return models.Student{}, err
	}

	if err := result.Last(&student).Error; err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func CleanSeeders() error {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	studentErr := DB.Exec("DELETE FROM students").Error

	var isFailed bool = studentErr != nil

	if isFailed {
		return errors.New("cleaning failed")
	}

	log.Println("seeders are cleaned up successfully")

	return nil
}