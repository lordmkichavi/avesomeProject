package model

type Comunidad struct {
	Nombre string `json:"nombre"`
}

type Comunidades []Comunidad

type Persona struct {
	Nombre      string      `json:"nombre"`
	Edad        uint8       `json:"edad"`
	Comunidades Comunidades `json:"comunidades"`
}

type Personas []Persona
