package handler

import (
	"awesomeProject/rest/middleware"
	"net/http"
)

func RoutePersona(mux *http.ServeMux, almacenamiento Almacenamiento) {
	h := newPersona(almacenamiento)
	mux.HandleFunc("/v1/personas/create", middleware.Log(middleware.Authorization(h.create)))
	mux.HandleFunc("/v1/personas/get-all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/personas/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/personas/delete", middleware.Log(h.delete))
}
