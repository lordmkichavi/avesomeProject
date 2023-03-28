package handler

import (
	"awesomeProject/rest/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type persona struct {
	almacenamiento Almacenamiento
}

func newPersona(almacenamiento Almacenamiento) persona {
	return persona{almacenamiento}
}

func (p *persona) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respuesta := newResponse(Error, "message:no encontrada", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data := model.Persona{}
	error := json.NewDecoder(r.Body).Decode(&data)
	if error != nil {
		respuesta := newResponse(Error, "la persona no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	err := p.almacenamiento.Create(&data)
	if err != nil {
		respuesta := newResponse(Error, "la persona no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, respuesta)
		return
	}

	respuesta := newResponse(Mensaje, "la persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, respuesta)
}

func (p *persona) createMultipart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respuesta := newResponse(Error, "message:no encontrada", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data := model.Persona{}
	error := json.NewDecoder(r.Body).Decode(&data)
	if error != nil {
		respuesta := newResponse(Error, "la persona no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	err := p.almacenamiento.Create(&data)
	if err != nil {
		respuesta := newResponse(Error, "la persona no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, respuesta)
		return
	}

	respuesta := newResponse(Mensaje, "la persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, respuesta)
}

func (p *persona) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respuesta := newResponse(Error, "Error en el metodo", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	res, error := p.almacenamiento.GetALL()
	if error != nil {
		respuesta := newResponse(Error, "Error en el llamado total", nil)
		responseJSON(w, http.StatusInternalServerError, respuesta)
		return
	}

	respuesta := newResponse(Mensaje, "OK", res)
	responseJSON(w, http.StatusOK, respuesta)
}

func (p *persona) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respuesta := newResponse(Error, "Metodo no encontrado", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		respuesta := newResponse(Error, "El id debe ser un numero", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data := model.Persona{}
	error := json.NewDecoder(r.Body).Decode(&data)
	if error != nil {
		respuesta := newResponse(Error, "La persona no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	error = p.almacenamiento.Update(ID, &data)
	if error != nil {
		respuesta := newResponse(Error, "La persona no tiene la estructura correct", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	respuesta := newResponse(Mensaje, "persona actualizada", nil)
	responseJSON(w, http.StatusOK, respuesta)
}

func (p *persona) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respuesta := newResponse(Error, "Metodo no encontrado", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		respuesta := newResponse(Error, "El id debe ser un numero", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	error := p.almacenamiento.Delete(ID)
	if errors.Is(error, model.ErrIDPersonaNoExiste) {
		respuesta := newResponse(Error, "id no existe", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
	}

	if error != nil {
		respuesta := newResponse(Error, "La persona no tiene la estructura correct", nil)
		responseJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	respuesta := newResponse(Mensaje, "persona eliminada", nil)
	responseJSON(w, http.StatusOK, respuesta)
}
