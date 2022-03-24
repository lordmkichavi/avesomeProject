package main

import (
	saludo2 "awesomeProject/saludo"
	"fmt"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println("hello")

	var g = saludo2.SaludarIngles("gatogato")

	var bbb = quoteV3.Concurrency()

	fmt.Println(bbb)
	fmt.Println(g)

	var bb = quote.Hello()

	fmt.Println(bb)
}
