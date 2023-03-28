package ejercicio

type Factura struct {
	pais      string
	ciudad    string
	total     float64
	cliente   Cliente
	productos Productos
}

func NewFactura(pais string, ciudad string, total float64, cliente Cliente, productos Productos) Factura {
	return Factura{
		pais:      pais,
		ciudad:    ciudad,
		total:     total,
		cliente:   cliente,
		productos: productos,
	}
}

func NewProductos(producto ...Producto) Productos {
	var valores Productos
	for _, prod := range producto {
		valores = append(valores, prod)
	}
	return valores
}

func (factura *Factura) SetTotal() {
	factura.total = factura.productos.Total()
}
