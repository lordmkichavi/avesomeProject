package almacenamiento

import (
	"awesomeProject/rest/model"
	"fmt"
)

type Memoria struct {
	idActual int
	Personas map[int]model.Persona
}

func NewMemory() Memoria {
	personas := make(map[int]model.Persona)
	return Memoria{
		idActual: 0,
		Personas: personas,
	}
}

func (m *Memoria) Create(person *model.Persona) error {
	if person == nil {
		return model.ErrNoPuedeSerNulo
	}
	m.idActual++
	m.Personas[m.idActual] = *person
	return nil
}

func (m *Memoria) Update(ID int, person *model.Persona) error {
	if person == nil {
		return model.ErrIDPersonaNoExiste
	}
	_, existePersona := m.Personas[ID]
	if !existePersona {
		fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonaNoExiste)
	}
	m.Personas[ID] = *person
	return nil
}

func (m *Memoria) Delete(ID int) error {
	_, existePersona := m.Personas[ID]
	if !existePersona {
		fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonaNoExiste)
	}
	delete(m.Personas, ID)
	return nil
}

func (m *Memoria) GetByID(ID int) (model.Persona, error) {
	existePersona, ok := m.Personas[ID]
	if !ok {
		return existePersona, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonaNoExiste)
	}
	return existePersona, nil
}

func (m *Memoria) GetALL() (model.Personas, error) {
	var listado model.Personas
	for _, persona := range m.Personas {
		listado = append(listado, persona)
	}
	return listado, nil
}
