package tour

import "fmt"

// Address has a City, State and a Country.
type Address struct {
	City    string
	State   string
	Country string
}

// Person has a Name, Age and Address.
type Person struct {
	Name string
	Age  uint
	Addr *Address
}

func (p Person) GoString() string {
	if p.Addr != nil {
		return fmt.Sprintf("PersonCta{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
	}
	return fmt.Sprintf("PersonCat{Name: %q, Age: %d}", p.Name, int(p.Age))
}
