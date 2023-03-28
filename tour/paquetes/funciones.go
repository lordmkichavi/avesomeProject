package paquetes

import (
	"fmt"
	"time"
)

func EstablecerValorCeroSinPuntero(ival int) {
	ival = 0
}

func EstablecerValorCeroConPunteros(iptr *int) {
	*iptr = 0
}

func EstablecerRetornoCeroConPunteros(iptr *int) *int {
	return iptr
}

func EstablecerRetornoCeroSinPunteros(iptr *int) int {
	return *iptr
}

//callback func(entero int) bool
func MultiplesFuncioesConRetornos(callback func(valor int) bool) func() bool {

	fmt.Println("iiiii")
	if callback((5)) {
		fmt.Println("ingresocallback")
	}

	fmt.Println("iii22")

	return devolverVerdadero
}

func devolverVerdadero() bool {
	return false
}

type MyFuncion func(string)

func MiddlewareLog(f MyFuncion) MyFuncion {
	return func(param string) {
		fmt.Println(time.Now().Format("2006-01-02"))
		f(param)
	}
}
