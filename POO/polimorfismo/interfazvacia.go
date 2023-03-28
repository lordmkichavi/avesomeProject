package polimorfismo

import (
	"fmt"
)

type Envoltorio interface {
	x()
}

func Encapsulador(inter interface{}) {
	fmt.Printf("valor %v, Tipo %T  \n", inter, inter)
	switch v := inter.(type) {
	case string:
		fmt.Println("cadena ", v)
	case int:
		fmt.Println("entero ", v)
	case PagoPaypal:
		v.Pagar()
	}
}
