package funciones

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func PruebaInterfaces() {
	/*var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	a = &v
	a = v
	fmt.Println(a.Abs())*/

	var i I = T{"hello"}
	var H I = T{"cayt"}
	i.M()
	H.M()
	//fmt.Printf(t)
}
