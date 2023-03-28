package almacenamiento

import (
	"awesomeProject/rest/model"
	"testing"
)

func TestCreate(t *testing.T) {
	datos := []struct {
		nombre    string
		model     *model.Persona
		wantError error
	}{
		{"vacio", nil, model.ErrNoPuedeSerNulo},
		{"nueva persona", &model.Persona{
			Nombre: "Corinna",
			Edad:   2,
		}, nil},
		{"nueva persona 2", &model.Persona{
			Nombre: "Corinna",
			Edad:   2,
		}, nil},
	}

	m := NewMemory()
	for _, dato := range datos {
		t.Run(dato.nombre, func(t *testing.T) {
			gotError := m.Create(dato.model)
			if gotError != dato.wantError {
				t.Errorf("se esperaba %d, y se obtuvo %d", gotError, dato.wantError)
			}
		})
	}
	wantCount := len(datos) - 1
	if m.idActual != wantCount {
		t.Errorf("se obtuvo %d y se obtuvo %d ", m.idActual, wantCount)
	}
}

func TestCreate_persona_vacia(t *testing.T) {
	m := NewMemory()
	err := m.Create(nil)
	if err == nil {
		t.Error("se obtuvo nuil, se esperaba nil")
	}
	if err != model.ErrNoPuedeSerNulo {
		t.Errorf("se esperaba el wantError %v y se obtuvo %v", model.ErrNoPuedeSerNulo, err)
	}
}

func TestCreate_contador_elementos(t *testing.T) {
	m := NewMemory()
	p := model.Persona{
		Nombre: "corinna",
		Edad:   2,
	}
	err := m.Create(&p)
	if err != nil {
		t.Errorf("se esperaba mayor el wantError y se obtuvo %v", err)
	}
	if m.idActual != 1 {
		t.Errorf("se esperaba mayor a 0, %d ", m.idActual)
	}
}
