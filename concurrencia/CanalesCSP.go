package concurrencia

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func PruebaCanal1() {
	mensaje := make(chan string)
	go func() {
		mensaje <- "Hola"
	}()
	var gg = <-mensaje
	fmt.Println(gg)
}

func PruebaCanal2() {
	senal := make(chan struct{})
	mensaje := make(chan int, 1)
	go recibir(senal, mensaje)
	enviar(mensaje)

	senal <- struct{}{}
}

func enviar(numero chan<- int) {
	numero <- 1
	numero <- 2
	numero <- 3
	numero <- 7
	numero <- 2
	numero <- 3
	numero <- 7
	numero <- 2
	numero <- 3
	numero <- 7
	close(numero)
}

func recibir(senal <-chan struct{}, numero <-chan int) {
	for {
		select {
		case v := <-numero:
			fmt.Println(v)
		case <-senal:
			return
		default:
			fmt.Println("pensando")
		}
	}
}

func SimularProcesoConcurrenteCSP() {
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
