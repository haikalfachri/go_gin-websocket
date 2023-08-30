package controllers

import (
	"go-gin-student/models"
	"go-gin-student/services"
	"net/http"

	ws "go-gin-student/websocket"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service services.StudentService
}

func InitStudentContoller() StudentController {
	return StudentController{
		service: services.InitStudentService(),
	}
}

// @Summary Get all students
// @Description Get a list of all students
// @Tags students
// @Produce json
// @Success 200 {object} models.ResponseSuccessGetAll
// @Failure 400 {object} models.ResponseMessage
// @Failure 500 {object} models.ResponseMessage
// @Router /students [get]
func (sc *StudentController) GetAll(c *gin.Context) {
	students, err := sc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ws.SendWebSocketUpdate("get all student")
	c.JSON(http.StatusOK, gin.H{"message": "success fetch all students", "data": students})
}

// @Summary Get a student by ID
// @Description Get a student by their ID
// @Tags students
// @Param id path string true "Student ID"
// @Produce json
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseMessage
// @Failure 500 {object} models.ResponseMessage
// @Router /students/{id} [get]
func (sc *StudentController) GetById(c *gin.Context) {
	id := c.Param("id")
	student, err := sc.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ws.SendWebSocketUpdate("get student by id")
	c.JSON(http.StatusOK, gin.H{"message": "success fetch a student", "data": student})
}

// @Summary Create a new student
// @Description Create a new student
// @Tags students
// @Accept json
// @Param student body models.RequestStudent true "Student data"
// @Produce json
// @Success 201 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseMessage
// @Failure 500 {object} models.ResponseMessage
// @Router /students [post]
func (sc *StudentController) Create(c *gin.Context) {
	var studentInput models.RequestStudent

	if err:=c.ShouldBindJSON(&studentInput);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	 }

	student, err := sc.service.Create(studentInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ws.SendWebSocketUpdate("student created")
	c.JSON(http.StatusCreated, gin.H{"message": "student created", "data": student})
}

// @Summary Delete a student by ID
// @Description Delete a student by their ID
// @Tags students
// @Param id path string true "Student ID"
// @Produce json
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ResponseMessage
// @Failure 500 {object} models.ResponseMessage
// @Router /students/{id} [delete]
func (sc *StudentController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := sc.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ws.SendWebSocketUpdate("student deleted")
	c.JSON(http.StatusOK, gin.H{"message": "success delete a student"})
}

// @Summary Update a student by ID
// @Description Update a student by their ID
// @Tags students
// @Param id path string true "Student ID"
// @Accept json
// @Param student body models.RequestStudent true "Updated student data"
// @Produce json
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseMessage
// @Failure 500 {object} models.ResponseMessage
// @Router /students/{id} [put]
func (sc *StudentController) Update(c *gin.Context) {
	id := c.Param("id")

	var studentInput models.RequestStudent

	if err:=c.ShouldBindJSON(&studentInput);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	 }

	student, err := sc.service.Update(studentInput, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ws.SendWebSocketUpdate("student updated")
	c.JSON(http.StatusOK, gin.H{"message": "student updated", "data": student})
}
