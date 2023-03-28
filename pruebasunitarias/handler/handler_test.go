package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	Get(w, r)
	if w.Code != http.StatusOK {

	}
	t.Log(w.Body.String())
	got := Persona{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("no se pudo leer %v", err)

	}
	want := Persona{
		"corinna", 2,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("se esperaba %v y se obtuvo  %v", want, got)
	}
}

func TestGetEcho(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	e := echo.New()
	ctx := e.NewContext(r, w)
	GetEcho(ctx)
	if w.Code != http.StatusOK {
		t.Errorf("Codigo esperado %d, se obtuvo %d", http.StatusOK, w.Code)
	}
	got := Persona{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("no se pudo leer %v", err)

	}
	want := Persona{
		"corinna", 2,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("se esperaba %v y se obtuvo  %v", want, got)
	}
}
