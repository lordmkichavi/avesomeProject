package testmockuniotarios

import (
	"errors"
)

//go:generate mockgen -destination=/Users/jafajardo/GolandProjects/awesomeProject/testmockuniotarios/MultipleTest_Mock.go -mock_names=Client=MockMultipleClient -source=/Users/jafajardo/GolandProjects/awesomeProject/testmockuniotarios/MultipleTest.go -package=testmockuniotarios

var (
	ErrMissingArgs   = errors.New("FirstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Searcher interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type Phonebook struct {
	People []*Person
}

func (p *Phonebook) Find(searcher Searcher, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := searcher.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
