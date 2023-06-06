package testmockuniotarios

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestObtenerArea(t *testing.T) {
	type args struct {
		fig FiguraGeometrica
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		times int
	}{
		/*{
			name: "calculando el area",
			args: args{
				retornarFiguraGeometrica(),
			},
			want: 999,
		},*/
		{
			name: "calculando el area",
			args: args{
				retornarFiguraGeometrica(),
			},
			want:  999,
			times: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockTT := NewMockFiguraGeometrica(ctrl)
			mockTT.EXPECT().CalcularArea().
				Return(999.00).MaxTimes(2)

			bbb := mockTT.EXPECT().CalcularArea()

			fmt.Println(bbb)

			fmt.Println(tt.want)
			fmt.Println(mockTT.CalcularArea())
			fmt.Println(mockTT.CalcularArea())

			//	t.Errorf("ObtenerArea() = %v, want %v", got, tt.want)

			//if got := ObtenerArea(tt.args.fig); got != tt.want {
			//	t.Errorf("ObtenerArea() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func retornarFiguraGeometrica() FiguraGeometrica {
	return Rectangulo{
		Base:   5,
		Altura: 8,
	}
}

func TestRectangulo_CalcularArea(t *testing.T) {
	type fields struct {
		Base   float64
		Altura float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			figura := Rectangulo{
				Base:   tt.fields.Base,
				Altura: tt.fields.Altura,
			}
			if got := figura.CalcularArea(); got != tt.want {
				t.Errorf("CalcularArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrapecio_CalcularArea(t *testing.T) {
	type fields struct {
		Base_mayor float64
		Base_menor float64
		Altura     float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			figura := Trapecio{
				Base_mayor: tt.fields.Base_mayor,
				Base_menor: tt.fields.Base_menor,
				Altura:     tt.fields.Altura,
			}
			if got := figura.CalcularArea(); got != tt.want {
				t.Errorf("CalcularArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retornoCondicional(t *testing.T) {
	type args struct {
		valor float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retornoCondicional(tt.args.valor); got != tt.want {
				t.Errorf("retornoCondicional() = %v, want %v", got, tt.want)
			}
		})
	}
}
