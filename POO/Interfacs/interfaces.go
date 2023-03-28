package Interfacs

import "fmt"

type ISaludar interface {
	Saludar(nombre string) string
}

type IDespedir interface {
	Despedir()
}

type Persona struct {
	Name string
}

type Texto string

type Multiple interface {
	ISaludar
	IDespedir
}

func (p Persona) Saludar(nombre string) string {
	return "paquetes desde persona " + nombre + " " + p.Name
}

func (p Persona) String() string {
	return "prueba " + p.Name + " vamos a ver "
}

func (p Persona) GoString() string {
	fmt.Println("hello kitty")
	if &p != nil {
		return fmt.Sprintf("gatooo Persona{Name: %q}", p.Name)
	}
	return fmt.Sprintf("gatooo2 Person{Name: %q}", p.Name)
}

func (p Texto) Saludar(nombre string) string {
	return "paquetes desde texto " + nombre + " dedos"
}

func (p Persona) Despedir() {
	fmt.Println("despedir perdona")
}

func (p Texto) Despedir() {
	fmt.Println("despedir texto")
}

func SaludarTodos(saludos ...Multiple) {
	for _, sal := range saludos {
		fmt.Println(sal.Saludar("car"))
		sal.Despedir()
	}
}
