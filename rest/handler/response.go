package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	TipoMensaje string      `json:"message-type"`
	Mensaje     string      `json:"message"`
	Data        interface{} `json:"data"`
}

const (
	Error   = "error"
	Mensaje = "message"
)

func newResponse(tipoMensaje string, mensaje string, dato interface{}) response {
	return response{
		TipoMensaje: tipoMensaje,
		Mensaje:     mensaje,
		Data:        dato,
	}
}

func responseJSON(w http.ResponseWriter, codigo int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codigo)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
