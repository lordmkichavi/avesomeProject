package cmd

import (
	"awesomeProject/rest/almacenamiento"
	"awesomeProject/rest/handler"
	"net/http"
)

func PruebaCMD() {
	almacen := almacenamiento.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePersona(mux, &almacen)
	http.ListenAndServe(":8080", mux)
}
