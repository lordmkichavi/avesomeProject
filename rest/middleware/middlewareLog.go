package middleware

import (
	"log"
	"net/http"
)

type MyLog func(http.ResponseWriter, *http.Request)

func Log(f MyLog) MyLog {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion%q, metodo%q ", r.URL.Path, r.Method)
		f(w, r)
	}
}

func Authorization(f MyLog) MyLog {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "no valido koken" {
			validarAcceso(w, r)
			return
		}
		f(w, r)
	}
}

func validarAcceso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("no tiene autorizacion"))
}
