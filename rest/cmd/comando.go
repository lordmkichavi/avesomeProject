package cmd

import (
	"awesomeProject/rest/almacenamiento"
	"awesomeProject/rest/handler"
	"log"
	"net/http"
)

func PruebaCMD() {
	almacen := almacenamiento.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePersona(mux, &almacen)
	err := http.ListenAndServe(":8080", mux)
	log.Println("iniciando desde el puerto 8080")
	if err != nil {
		log.Printf("Error %v", err)
	}
}
