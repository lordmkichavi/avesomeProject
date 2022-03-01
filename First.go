package main

import "fmt"

func main() {

	/*var dog string
	dog = "perro"

	var cat string = "cat"

	var mico, dedo string = "ffff", "ggggg"

	fmt.Println("Hello, world")
	fmt.Println(dog)
	fmt.Println(cat)
	fmt.Println(mico)
	fmt.Println(dedo)

	var cosa, cosa2 = "dddd", 1
	declarada := "mouse"
	declarada = "dddddd"
	cosa = "jjj"

	fmt.Println(cosa)
	fmt.Println(cosa2)
	fmt.Println(declarada)

	fmt.Println(cosa, cosa2, declarada)

	const pi = 3.1

	fmt.Println(pi)

	var a bool = true
	var bc string = "cadena"
	var num uint8 = 255
	var num2 int8 = -120
	var bb rune = 'a'
	var nnn float64 = 23.56
	var nn2 float32 = 98.78

	var ggg uint8 = 100
	var ggg2 uint16 = 100

	c := uint16(ggg) + ggg2

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v\n", bc, bc)
	fmt.Printf("Tipo: %T, Valor: %v\n", num, num)
	fmt.Printf("Tipo: %T, Valor: %v\n", num2, num2)
	fmt.Printf("Tipo: %T, Valor: %v\n", bb, bb)
	fmt.Printf("Tipo: %T, Valor: %v\n", bb, bb)
	fmt.Printf("Tipo: %T, Valor: %v\n", nnn, nnn)
	fmt.Printf("Tipo: %T, Valor: %v\n", nn2, nn2)
	fmt.Println("TOTAL " + string(c))

	var nnnn uint

	fmt.Println(nnnn)

	fmt.Println("TOTAL ", 70)

	var fruta = "manzana"
	var memoria *string
	memoria = &fruta
	*memoria = "gallo"
	fmt.Println("Tipo: %T, Valor: %s, Direccion:%v\n", fruta, fruta, &fruta)
	fmt.Println("Tipo: %T, Valor: %s\n", memoria, memoria, *memoria)


	set := [8]string{"a", "b", "c", "d", "e", "g", "h"}
	animals := set[0:5]
	fly := set[3:7]
	fly[0] = "k"

	fmt.Println(set)
	fmt.Println(animals)
	fmt.Println(fly)
	fmt.Println(set[:])

	animals := make(map[string]string)
	animals["cat"] = "cat"
	animals["dog"] = "dogiiii"
	fmt.Println(animals)
	fruits := map[string]string{
		"gato":  "gato",
		"perro": "mico",
	}
	fmt.Println(fruits)
	delete(animals, "cat")
	fmt.Println(animals)
	if _, ok := animals["perro"]; !ok {
		animals["perro"] = "miau"
	}

	fmt.Println(animals)
	*/

	/*type course struct {
		Name     string
		Profesor string
		Country  string
	}
	db := course{
		Name:     "Juan",
		Profesor: "Alexis",
		Country:  "Colombia",
	}
	fmt.Printf("%+v\n\n", db)
	gota := course{"hugo", "paco", "luis"}
	fmt.Printf("%+v\n\n", gota)
	nene := course{Profesor: "marco"}
	fmt.Printf("%+v\n\n", nene)*/
	//c := db
	//p := &db
	//fmt.Println(db)
	//fmt.Println(p)
	//p.Profesor = "javier"
	//fmt.Println(p)
	//fmt.Println(db)
	//fmt.Println(c)

	// hello()
	var fff = "ffff"
	var bb = change(&fff)
	fmt.Println(*bb)
	fmt.Println(fff)
}

func hello() {
	fmt.Println("hello")
}

func change(gato *string) *string {
	*gato = "pez"
	return gato
}
