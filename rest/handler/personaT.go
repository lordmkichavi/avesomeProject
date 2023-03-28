package handler

import (
	"awesomeProject/rest/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type personaT struct {
	almacenamiento Almacenamiento
}

func newPersonaT(almacenamiento Almacenamiento) personaT {
	return personaT{almacenamiento}
}

func (p *personaT) create(c echo.Context) error {
	data := model.Persona{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.almacenamiento.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Mensaje, "persona creada correctamente.", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *personaT) getAll() {

}

func (p *personaT) update() {

}

func (p *personaT) delete() {

}
