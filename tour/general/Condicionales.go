package main

import (
	saludo2 "awesomeProject/paquetes"
	"errors"
	"fmt"
	"io/ioutil"
)
import "strings"

func main() {
	var p = "cat"

	if p == "cat" {
		fmt.Println("catito")
	} else if p == "dog" {
		fmt.Println("dog")
	}

	imprimitNombre("javier", "fajardo")
	var g = "gato"
	cambiar(&g)
	fmt.Println(g)
	var ff, dd = cambiarNombres("gAo", "vFrde")
	fmt.Println(ff, dd)

	manejoError()

	r, err := division(6, 0)

	if err != nil {
		fmt.Printf("error numero %v %v \n", err, r)
	} else {
		fmt.Printf("error numero %v %v \n", err, r)
	}

	nums := []int{2, 3, 4, 50, 8, 12, 25, 70, 3, 2}

	gat, err := filter(nums, func(valor int) bool {
		var valorR = 5
		if valorR < valor {
			return true
		} else {
			return false
		}
	})

	fmt.Printf("valores %v %v \n", gat, err)

	x := retornaFuncion("gato")("dedo")
	fmt.Println(x)

	rr := func(valor string) string {
		return "hello " + valor
	}("cat")

	fmt.Println(rr)

	defer fmt.Println("1")

	var b = "dedo"

	defer fmt.Println(b)

	b = "kitty"

	defer fmt.Println(b)

	divisionPanic(12, 3)
	divisionPanic(12, 0)
	divisionPanic(12, 6)

	panico2()

	var saludoEs = saludo2.SaludarSpanish("gato")
	fmt.Println(saludoEs)

	var saludoEn = saludo2.SaludarIngles("gato")
	fmt.Println(saludoEn)
}

func panico2() {
	slice := []int{3, 4, 5, 6, 8}
	fmt.Println(slice[3]) // 6
	fmt.Println(slice[7]) // panic: runtime error: index out of range

	fmt.Println("Fin") // no se ejecuta
}

func divisionPanic(dividendo int, divisor int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("recuperandome %v \n", r)
		}
	}()
	valorarNumerador(divisor)
	var div = dividendo / divisor
	fmt.Printf("la division es: %v\n", div)
}

func valorarNumerador(numerador int) {
	if numerador == 0 {
		panic("hay un problema!")
	}
}

func retornaFuncion(nombre string) func(cat string) string {
	return func(cat string) string {
		var gato = "hello " + nombre + " cat " + cat
		return gato
	}
}

func filter(valores []int, callback func(entero int) bool) ([]int, error) {
	var s []int
	if len(valores) == 0 {
		return nil, errors.New("cadena vacia")
	} else {
		for _, valor := range valores {
			if callback(valor) {
				s = append(s, valor)
			}
		}
		return s, nil
	}
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		errorG := errors.New("no puede ser 0")
		return 0, errorG
	} else {
		return dividendo / divisor, nil
	}
}

func manejoError() {
	var archivo, err = ioutil.ReadFile("ato.txt")

	if err != nil {
		fmt.Printf("no encontro el archivo %v \n", err)
	} else {
		fmt.Println(string(archivo))
	}

}

func cambiar(valor *string) {
	*valor = *valor + "kitty"
}

func imprimitNombre(nombre string, apelligo string) {
	fmt.Printf("El nombre es %s de %s \n", nombre, apelligo)
}

func cambiarNombres(nombre string, apellido string) (string, string) {
	var minusculas = strings.ToLower(nombre + apellido)
	var mayusculas = strings.ToUpper(nombre + apellido)
	return minusculas, mayusculas
}
