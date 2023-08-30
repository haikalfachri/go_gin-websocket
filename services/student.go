package services

import (
	"go-gin-student/repositories"
	"go-gin-student/models"
)

type StudentService struct {
	repository repositories.StudentRepository
}

func InitStudentService() StudentService {
	return StudentService{
		repository: &repositories.StudentRepositoryImp{},
	}
}

func (sr *StudentService) GetAll() ([]models.Student, error){
	return sr.repository.GetAll()
}

func (sr *StudentService) GetById(id string) (models.Student, error){
	return sr.repository.GetById(id)
}

func (sr *StudentService) Create(studentInput models.RequestStudent) (models.Student, error){
	return sr.repository.Create(studentInput)
}

func (sr *StudentService) Update(studentInput models.RequestStudent, id string) (models.Student, error){
	return sr.repository.Update(studentInput, id)
}

func (sr *StudentService) Delete(id string) (error){
	return sr.repository.Delete(id)
}