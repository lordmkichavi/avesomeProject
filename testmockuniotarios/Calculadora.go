package testmockuniotarios

//go:generate mockgen -destination=/Users/jafajardo/GolandProjects/awesomeProject/testmockuniotarios/Calculadora_Mock.go -mock_names=Client=MockCalculatorClient -source=/Users/jafajardo/GolandProjects/awesomeProject/testmockuniotarios/Calculadora.go -package=testmockuniotarios
type FiguraGeometrica interface {
	CalcularArea() float64
}

type Rectangulo struct {
	Base, Altura float64
}

type Trapecio struct {
	Base_mayor, Base_menor, Altura float64
}

func (figura Rectangulo) CalcularArea() float64 {
	return figura.Base * figura.Altura
}

func (figura Trapecio) CalcularArea() float64 {
	return (figura.Base_mayor + figura.Base_menor) * figura.Altura / 2
}

func ObtenerArea(fig FiguraGeometrica) float64 {
	return retornoCondicional(fig.CalcularArea())
}

func retornoCondicional(valor float64) float64 {
	if valor > 10 {
		return 999
	} else {
		return 222
	}
}
