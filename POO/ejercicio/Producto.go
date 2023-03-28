package ejercicio

type Producto struct {
	id       uint
	producto string
	valor    float64
}

type Productos []Producto

func NewProducto(id uint, producto string, valor float64) Producto {
	return Producto{
		id:       id,
		producto: producto,
		valor:    valor,
	}
}

func (listaProducto Productos) Total() float64 {
	var total float64 = 0
	for _, producto := range listaProducto {
		total = total + producto.valor
	}
	return total
}
