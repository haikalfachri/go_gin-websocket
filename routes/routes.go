package routes

import (
	_"go-gin-student/docs" //kunci nya biar work
	"go-gin-student/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	
)

func SetUpRoutes(g *gin.Engine) {
	studentCtrl := controllers.InitStudentContoller()
	
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.GET("/students", studentCtrl.GetAll)
	g.GET("/students/:id", studentCtrl.GetById)
	g.POST("/students", studentCtrl.Create)
	g.PUT("/students/:id", studentCtrl.Update)
	g.DELETE("/students/:id", studentCtrl.Delete)
}