package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPersona_Create_Estructura_Erronea(t *testing.T) {
	datos := []byte(`{"Nombre" : 123, "Edad" : 18 }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(datos))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)
	p := newPersonaT(&AlmacenamientoMockOK{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba %v", err)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("codigo estado %d, se esperaba se obtuvo %d", http.StatusBadRequest, w.Code)
	}
	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo hacer unmarshall")
	}
	wantMessage := "La persona no tiene una estructura"
	if resp.Mensaje != wantMessage {
		t.Errorf("no se pudo hacer unmarshall, se obtuvo %q, se esperaba %q", resp.Mensaje, wantMessage)
	}
}

func TestPersona_Create_Memoria_Erronea(t *testing.T) {
	datos := []byte(`{"Nombre" : "javier", "Edad" : 18, "communities": [] }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(datos))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)
	p := newPersonaT(&AlmacenamientoMockError{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba %v", err)
	}
	if w.Code != http.StatusInternalServerError {
		t.Errorf("codigo estado %d, se esperaba se obtuvo %d", http.StatusInternalServerError, w.Code)
	}
	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo hacer unmarshall")
	}
	wantMessage := "Hubo un problema al crear la persona"
	if resp.Mensaje != wantMessage {
		t.Errorf("no se pudo hacer unmarshall, se obtuvo %q, se esperaba %q", resp.Mensaje, wantMessage)
	}
}
