package concurrencia

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cuenta struct {
	nombre  string
	balance int
}

func Tranferencia(monto int, origen *Cuenta, destino *Cuenta) {
	if origen.balance < monto {
		return
	}
	//time.Sleep(time.Second)
	time.Sleep(5 * time.Second)

	destino.balance = destino.balance + monto
	origen.balance = origen.balance - monto

	fmt.Printf("origen: %s , destino %s \n", origen, destino)
}

func EjecutarMovimientosSecuencial() {
	log.Printf("Inicial ")
	inicial := time.Now()
	fmt.Println("inicio")
	c1 := Cuenta{"javier", 500}
	c2 := Cuenta{"corinna", 900}

	for _, i2 := range []int{50, 50, 50, 50, 60, 50, 50, 50, 50, 50} {
		Tranferencia(i2, &c1, &c2)
	}

	elapsed := time.Since(inicial)
	log.Printf("Transcurrido %s", elapsed)
	fmt.Println("fin")
}

func EjecutarMovimientosConcurrente() {
	log.Printf("Inicial ")
	var pagos = []int{50, 50, 50, 50, 60, 50, 50, 50, 50, 50}
	inicial := time.Now()
	fmt.Println("inicio")
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(pagos))

	c1 := Cuenta{"javier", 500}
	c2 := Cuenta{"corinna", 900}

	for _, valor := range pagos {
		go func(monto int) {
			mu.Lock()
			Tranferencia(monto, &c1, &c2)
			mu.Unlock()
			wg.Done()
		}(valor)
	}

	wg.Wait()

	elapsed := time.Since(inicial)
	log.Printf("Transcurrido %s", elapsed)
	fmt.Println("fin")
}
