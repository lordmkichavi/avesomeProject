package abstraccion

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) ImprimirClasses() {
	valoresUno := course{
		Price:   45.67,
		IsFree:  true,
		UserIDs: []uint{2, 3, 4},
		Classes: map[uint]string{
			1: "Calculo",
			2: "Fisica",
			3: "Catedra",
		},
	}
	fmt.Println(c.Price)

	for u, s := range valoresUno.Classes {
		fmt.Printf("valor %v is %c\n", u, s)
	}
}

func (c course) CambiarPrecio(nuevoPrecio float64) {
	c.Price = nuevoPrecio
}
