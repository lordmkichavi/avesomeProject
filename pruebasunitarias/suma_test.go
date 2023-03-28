package pruebasunitarias

import (
	"reflect"
	"testing"
)

type AnimalT8 struct {
	nombre string
	edad   int
}

func TestSuma(t *testing.T) {
	what := 8
	got := Sumar(5, 3)
	t.Logf("cat %d", got)
	if what != got {
		t.Errorf("se esperaba %d, y se obtuvo %d", what, got)
	}
}

func TestCompararEstructuras(t *testing.T) {
	gato := &AnimalT8{
		nombre: "firulo",
		edad:   3,
	}

	perro := &AnimalT8{
		nombre: "firulo",
		edad:   3,
	}

	if !reflect.DeepEqual(gato, perro) {
		t.Errorf("son diferentes uno %v y 2 %v ", perro, gato)
	}
}

func TestSumarAcomulativa(t *testing.T) {
	what := 9
	got := SumarAcomulativa(2, 3, 4)
	if what != got {
		t.Errorf("se esperaba %d, y se obtuvo %d", what, got)
	}
}

func TestSumarAcomulativaMultiple(t *testing.T) {
	table := []struct {
		num1 int
		num2 int
		what int
		name string
	}{
		{1, 1, 2, "1+1"},
		{1, 2, 7, "1+2"},
		{2, 1, 3, "2+1"},
	}
	for _, i2 := range table {
		t.Run(i2.name, func(t *testing.T) {
			got := Sumar(i2.num1, i2.num2)
			if got != i2.what {
				t.Errorf("se esperaba %d, y se obtuvo %d", got, i2.what)
			}
		})
	}
}

func BenchmarkSumarAcomulativa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumarAcomulativa(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
