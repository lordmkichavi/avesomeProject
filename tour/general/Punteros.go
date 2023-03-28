package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (c *circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func info(s forma) {
	fmt.Println("area", s.area())
}

func main() {
	/*x := 5
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)*/
	//fmt.Println("gato")
	//fmt.Println("punteros")

	//c := circulo{4}
	//info(c)
	//info(&c)

	nunerosMultinivel()

	//validarManejoPunterosFucnciones()
}

func nunerosMultinivel() {
	x := 10
	px := &x
	ppx := &px

	fmt.Println("x =", x, ",px =", *px, ",ppx =", **ppx)
	*px = *px * 2
	fmt.Println("x =", x, ",px =", *px, ",ppx =", **ppx)
	**ppx = **ppx + 5
	fmt.Println("x =", x, ",px =", *px, ",ppx =", **ppx)
}

func localSwap(this, that string) {
	hold := this
	this = that
	that = hold
}

func realSwap(this, that *string) {
	hold := *this // hold -> this
	*this = *that // this -> that
	*that = hold  // that -> this
}

func validarManejoPunterosFucnciones() {
	this := "this"
	that := "that"

	fmt.Println("this = ", this, "that = ", that)
	localSwap(this, that)
	fmt.Println("After swap: this = ", this, "that = ", that) //no change
	realSwap(&this, &that)
	fmt.Println("After real swap: this = ", this, "that = ", that)
}
