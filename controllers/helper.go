package controllers

import (
	"go-gin-student/database"

	"github.com/gin-gonic/gin"
)

var configDB database.Config = database.Config{
	DB_USERNAME: "root",
	DB_PASSWORD: "",
	DB_HOST:     "localhost",
	DB_PORT:     "3306",
	DB_NAME:     "student_crud_db",
}

var configDBFailed database.Config = database.Config{
	DB_USERNAME: "root123",
	DB_PASSWORD: "",
	DB_HOST:     "db",
	DB_PORT:     "3306",
	DB_NAME:     "student_crud_db",
}

var g = gin.Default()

type TestCase struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

func InitGin() {
	gin.SetMode(gin.TestMode)
	configDB.ConnectDB()
	database.MigrateDB()
	g := gin.Default()
	g.Run(":9000")
}

// func InitGinFailed() {
// 	gin.SetMode(gin.TestMode)
// 	configDBFailed.ConnectDB()
// 	database.MigrateDB()
// 	g := gin.Default()
// 	g.Run(":9000")
// }