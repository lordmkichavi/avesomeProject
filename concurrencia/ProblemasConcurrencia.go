package concurrencia

import (
	"fmt"
	"sync"
)

func PruebaLlamadaGoRutina() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	data := 1

	go func() {
		mu.Lock()
		data++
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(data)
}

func PruebaLlamadaGoRutina2() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(numero int) {
			fmt.Println(numero)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func PruebaLlamadaGoRutina3() {
	cursos := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	go func() {
		mu.Lock()
		cursos["go desde cero"] = "basico"
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		cursos["go concurrencia"] = "avanzado"
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(cursos)
}
