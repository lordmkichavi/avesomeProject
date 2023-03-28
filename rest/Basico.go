package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RegistrarHolaMundo() {
	fmt.Println("inicio")
	http.HandleFunc("/saludar", saludar)

	server := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
	fmt.Println("fin")
}

func IniciarServidorWeb() {
	//http://localhost:8080/inicio.html
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("subiendo el servidor web")
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}
