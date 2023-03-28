package handler

import (
	"awesomeProject/rest/model"
	"errors"
)

type AlmacenamientoMockError struct {
}

func (ame *AlmacenamientoMockError) Create(person *model.Persona) error {
	return errors.New("error en Mock")
}

func (ame *AlmacenamientoMockError) Update(ID int, person *model.Persona) error {
	return errors.New("error en Mock")
}

func (ame *AlmacenamientoMockError) Delete(ID int) error {
	return errors.New("error en Mock")
}

func (ame *AlmacenamientoMockError) GetByID(ID int) (model.Persona, error) {
	return model.Persona{}, errors.New("error en Mock")
}

func (ame *AlmacenamientoMockError) GetALL() (model.Personas, error) {
	return model.Personas{}, errors.New("error en Mock")
}

type AlmacenamientoMockOK struct {
}

func (ame *AlmacenamientoMockOK) Create(person *model.Persona) error {
	return nil
}

func (ame *AlmacenamientoMockOK) Update(ID int, person *model.Persona) error {
	return nil
}

func (ame *AlmacenamientoMockOK) Delete(ID int) error {
	return nil
}

func (ame *AlmacenamientoMockOK) GetByID(ID int) (model.Persona, error) {
	return model.Persona{}, nil
}

func (ame *AlmacenamientoMockOK) GetALL() (model.Personas, error) {
	return model.Personas{}, nil
}
