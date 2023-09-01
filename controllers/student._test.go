package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"go-gin-student/database"
	"go-gin-student/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var studentController *StudentController = InitStudentContoller()

func TestGetAllStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"data\":",
	}

	InitGin()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	studentController.GetAll(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
}

// func TestGetAllStudent_Failed(t *testing.T) {
// 	testcase := TestCase{
// 		name:                   "failed",
// 		path:                   "/students",
// 		expectedStatus:         http.StatusBadRequest,
// 		expectedBodyStartsWith: "{\"message\":",
// 	}

// 	InitGinFailed()

// 	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
// 	rec := httptest.NewRecorder()

// 	ctx, _ := gin.CreateTestContext(rec)
// 	ctx.Request = req

// 	studentController.GetAll(ctx)

// 	assert.Equal(t, testcase.expectedStatus, rec.Code)
// 	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
// }

func TestGetByIdStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"data\":",
	}

	InitGin()

	student, _ := database.SeedStudent()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(student.ID), 10)})

	studentController.GetById(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
}

func TestGetByIdStudent_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "failed",
		path:                   "/students/:id",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"message\":",
	}

	InitGin()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(0), 10)})

	studentController.GetById(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
}

func TestCreateStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"data\":",
	}

	InitGin()
	
	student := models.RequestStudent{
		FirstName: "test",
		LastName:  "test",
		StudentID: "test",
		Major:     "test",
	}

	studentJSON, _ := json.Marshal(student)

	req := httptest.NewRequest(http.MethodPost, testcase.path, strings.NewReader(string(studentJSON)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(studentJSON)))

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	studentController.Create(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))

	_ = database.CleanSeeders()
}

// func TestCreateStudent_Failed(t *testing.T) {
// 	testcase := TestCase{
// 		name:                   "failed",
// 		path:                   "/students",
// 		expectedStatus:         http.StatusBadRequest,
// 		expectedBodyStartsWith: "{\"message\":",
// 	}

// 	gin.SetMode(gin.ReleaseMode)
// 	g := gin.Default()
// 	g.Run(":9000")
	
// 	student := models.RequestStudent{
// 		FirstName: "test",
// 		LastName:  "test",
// 		StudentID: "test",
// 		Major:     "test",
// 	}

// 	studentJSON, _ := json.Marshal(student)

// 	req := httptest.NewRequest(http.MethodPost, testcase.path, strings.NewReader(string(studentJSON)))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(studentJSON)))

// 	rec := httptest.NewRecorder()

// 	ctx, _ := gin.CreateTestContext(rec)
// 	ctx.Request = req

// 	studentController.Create(ctx)

// 	assert.Equal(t, testcase.expectedStatus, rec.Code)
// 	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))

// 	_ = database.CleanSeeders()
// }

func TestDeleteStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	InitGin()

	student, _ := database.SeedStudent()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(student.ID), 10)})

	studentController.Delete(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
}

func TestDeleteStudent_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "failed",
		path:                   "/students/:id",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"message\":",
	}

	InitGin()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(0), 10)})

	studentController.Delete(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))
}

func TestUpdateStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"data\":",
	}

	InitGin()

	student, _ := database.SeedStudent()
	
	studentUpdate := models.RequestStudent{
		FirstName: "test",
		LastName:  "test",
		StudentID: "test",
		Major:     "test",
	}

	studentJSON, _ := json.Marshal(studentUpdate)

	req := httptest.NewRequest(http.MethodPost, testcase.path, strings.NewReader(string(studentJSON)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(studentJSON)))

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(student.ID), 10)})

	studentController.Update(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))

	_ = database.CleanSeeders()
}

func TestUpdateStudent_Failed(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students/:id",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"message\":",
	}

	InitGin()
	
	studentUpdate := models.RequestStudent{
		FirstName: "test",
		LastName:  "test",
		StudentID: "test",
		Major:     "test",
	}

	studentJSON, _ := json.Marshal(studentUpdate)

	req := httptest.NewRequest(http.MethodPost, testcase.path, strings.NewReader(string(studentJSON)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(studentJSON)))

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(0), 10)})

	studentController.Update(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)
	assert.True(t, strings.HasPrefix(rec.Body.String(), testcase.expectedBodyStartsWith))

	_ = database.CleanSeeders()
}
