package Interfacs

import "fmt"

type Almacenador interface {
	Get() string
	Set(nombre string)
}

type Animal struct {
	Nombre string
}

func NewAnimal(nombre string) *Animal {
	return &Animal{
		nombre,
	}
}

func Exec(almacenador Almacenador, nombre string) {
	almacenador.Set(nombre)
	fmt.Println("here", almacenador.Get())
}

func (p Animal) Get() string {
	return p.Nombre
}

func (p *Animal) Set(nombre string) {
	p.Nombre = nombre
}
