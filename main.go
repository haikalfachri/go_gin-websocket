package main

import (
	"go-gin-student/database"
	"go-gin-student/routes"
	"go-gin-student/utils"
	ws "go-gin-student/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Student API
// @version 1.0
// @description API endpoints for students CRUD
// @host localhost:9000
// @BasePath /
func main() {
	configDB := database.Config{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	configDB.ConnectDB()

	database.MigrateDB()

	r := gin.Default()

	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = []string{"http://localhost:" + utils.GetConfig("APP_PORT")}
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(configCORS))

	routes.SetUpRoutes(r)

	go ws.StartWebSocketServer()

	r.Run(":" + utils.GetConfig("APP_PORT"))
}
