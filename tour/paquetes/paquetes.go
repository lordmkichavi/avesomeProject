package paquetes

import (
	"fmt"
	strftime "github.com/itchyny/timefmt-go"
	"math/rand"
	"time"
)

func PruebaPaquete() {

	fmt.Println("My favorite number is", rand.Intn(10))

	t := time.Now()
	fmt.Println(t)
	b := strftime.Format(t, "%Y%m%d%H%M%S")

	fmt.Println(b)

}
