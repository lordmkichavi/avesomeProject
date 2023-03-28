package concurrencia

import (
	"log"
	"time"
)

func PruebaCanalCancelacionSegura() {
	hecho := make(chan struct{})

	for _, tiempo := range tiempos {
		go func(t time.Duration) {
			EjecutarDuracionCSP(t)
			select {
			case <-hecho:
				return
			}
		}(tiempo)
	}

	select {
	case <-time.After(time.Second * 10):
		close(hecho)
	}
}

func EjecutarDuracionCSP(segundos time.Duration) {
	log.Printf("Inicial ", segundos)
	inicial := time.Now()
	time.Sleep(segundos * time.Second)
	elapsed := time.Since(inicial)
	log.Printf("Transcurrido %s", elapsed)
}

func iterarProcesos() {

}
