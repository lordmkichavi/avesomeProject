package pruebasunitarias

import (
	"fmt"
)

func Sumar(numero1, numero2 int) int {
	return numero1 + numero2
}

func SumarAcomulativa(numbers ...int) int {
	var resultado int = 0
	for _, number := range numbers {
		resultado = resultado + number
	}
	return resultado
}

func PruebaUnitarioa1() {
	fmt.Println("caty")
}
