package repositories

import "go-gin-student/models"

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetById(id string) (models.Student, error)
	Create(studentInput models.RequestStudent) (models.Student, error)
	Update(studentInput models.RequestStudent, id string) (models.Student, error)
	Delete(id string) error
}