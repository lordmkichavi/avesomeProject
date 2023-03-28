package ejercicio

import "fmt"

type Cliente struct {
	nombre    string
	direccion string
	telefono  string
}

type Empleado struct {
	Cliente
	salario float64
}

func NewCliente(nombre string, direccion string, telefono string) Cliente {
	return Cliente{
		nombre:    nombre,
		direccion: direccion,
		telefono:  telefono,
	}
}

func (empleado Cliente) Saludar(saludo string) string {
	fmt.Println("hola cliente " + saludo)
	return "mensaje"
}

func NewEmpleado(nombre string, direccion string, telefono string, salario float64) Empleado {
	return Empleado{NewCliente(nombre, direccion, telefono), salario}
}

func (empleado *Empleado) CalcularNomina() {
	empleado.salario = empleado.salario * 0.9
}
