package repositories

import (
	"go-gin-student/database"
	"go-gin-student/models"
)

type StudentRepositoryImp struct {
}

func InitStudentRepository() StudentRepository {
	return &StudentRepositoryImp{}
}

func (sr *StudentRepositoryImp) GetAll() ([]models.Student, error) {
	var students []models.Student

	if err := database.DB.Find(&students).Error; err != nil {
		return students, err
	}
	return students, nil
}

func (sr *StudentRepositoryImp) GetById(id string) (models.Student, error) {
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (sr *StudentRepositoryImp) Create(request models.RequestStudent) (models.Student, error) {

	var student models.Student = models.Student{
		FirstName: request.FirstName,
		LastName: request.LastName,
		StudentID: request.StudentID,
		Major: request.Major,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		return models.Student{}, err
	}

	if err := database.DB.Last(&student).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (sr *StudentRepositoryImp) Update(request models.RequestStudent, id string) (models.Student, error) {

	student, err := sr.GetById(id)

	if err != nil {
		return models.Student{}, err
	}

	if student.FirstName != request.FirstName {
		student.FirstName = request.FirstName
	}

	if student.LastName != request.LastName {
		student.LastName = request.LastName
	}

	if student.StudentID != request.StudentID {
		student.StudentID = request.StudentID
	}

	if student.Major != request.Major {
		student.Major = request.Major
	}

	if err := database.DB.Save(&student).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (sr *StudentRepositoryImp) Delete(id string) error {
	student, err := sr.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Unscoped().Delete(&student).Error; err != nil {
		return err
	}

	return nil
}