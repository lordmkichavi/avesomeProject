package concurrencia

import (
	"log"
	"sync"
	"time"
)

/*var tiempos = []time.Duration{
	3, 4, 7, 15, 2, 4, 5, 9, 9,
	7, 7, 7, 7, 9, 6, 9, 10, 14,
	9, 6, 7, 6, 2, 5, 5, 9, 13,
	2, 5, 7, 3, 3, 7, 8, 11, 11,
	1, 4, 7, 4, 2, 6, 5, 9, 10,
	4, 2, 7, 5, 1, 4, 7, 13, 9,
	8, 1, 7, 1, 2, 2, 7, 9, 10,
}*/

var tiempos = []time.Duration{
	9, 4, 7,
}

func EjecutarDuracion(segundos time.Duration) {
	log.Printf("Inicial ", segundos)
	inicial := time.Now()
	time.Sleep(segundos * time.Second)
	elapsed := time.Since(inicial)
	log.Printf("Transcurrido %s", elapsed)
}

func SimularProcesoSecuencial() {
	log.Printf("INICIAL")
	inicial := time.Now()
	for _, tiempo := range tiempos {
		EjecutarDuracion(tiempo)
	}
	elapsed := time.Since(inicial)
	log.Printf("TOTAL %s", elapsed)
}

func SimularProcesoConcurrente() {
	var wg sync.WaitGroup
	wg.Add(len(tiempos))
	log.Printf("INICIAL")
	inicial := time.Now()
	for _, tiempo := range tiempos {
		go func(t time.Duration) {
			EjecutarDuracion(t)
			wg.Done()
		}(tiempo)
	}
	wg.Wait()
	elapsed := time.Since(inicial)
	log.Printf("TOTAL %s", elapsed)
}
