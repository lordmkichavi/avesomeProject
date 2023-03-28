package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Persona struct {
	Nombre string
	Edad   int
}

func Get(w http.ResponseWriter, r *http.Request) {
	p := Persona{
		"corinna", 2,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetEcho(c echo.Context) error {
	p := Persona{
		"corinna", 2,
	}

	return c.JSON(http.StatusOK, p)
}
