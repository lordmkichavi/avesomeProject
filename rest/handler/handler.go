package handler

import (
	"awesomeProject/rest/model"
)

type Almacenamiento interface {
	Create(person *model.Persona) error
	Update(ID int, person *model.Persona) error
	Delete(ID int) error
	GetByID(ID int) (model.Persona, error)
	GetALL() (model.Personas, error)
}
