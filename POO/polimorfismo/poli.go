package polimorfismo

import "fmt"

type IMetodosPago interface {
	Pagar()
}

type PagoPaypal struct {
}

func (p PagoPaypal) Pagar() {
	fmt.Println("pagado por paypal")
}

type PagoEfectivo struct {
}

func (p PagoEfectivo) Pagar() {
	fmt.Println("pagado por efectivo")
}

type PagoTarjetaCredito struct {
}

func (p PagoTarjetaCredito) Pagar() {
	fmt.Println("pagado por tarejta de credito")
}

func MetodoFabrica(indiceMetodo int) IMetodosPago {
	switch indiceMetodo {
	case 1:
		return PagoPaypal{}
	case 2:
		return PagoEfectivo{}
	case 3:
		return PagoTarjetaCredito{}
	default:
		return nil
	}
}
