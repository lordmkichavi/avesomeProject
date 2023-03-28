package model

import "errors"

var (
	ErrNoPuedeSerNulo    = errors.New("La persona no puede ser nula")
	ErrIDPersonaNoExiste = errors.New("La persona no existe")
)
