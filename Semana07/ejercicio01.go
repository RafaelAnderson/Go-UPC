// Algoritmo 6.8 Producer-Consumer - Estructura Mutex
package main

import (
	"fmt"
	"sync"
)

var buffer []int
var tope = 10 // buffer finito

func productor(m *sync.Mutex) {

	var d int
	for {
		m.Lock()
		if len(buffer) < tope {
			// Producir
			d++
			buffer = append(buffer, d) // Agregar el valor d al slide
			fmt.Printf("Produciendo %d\n", d)
		} else {
			fmt.Println("Buffer lleno!!")
		}
		m.Unlock()
	}
}

func consumidor(m *sync.Mutex) {
	var d int
	for {
		m.Lock()
		if len(buffer) > 0 {
			// Producir
			d = buffer[0]
			buffer = buffer[1:] // Agregar el valor d al slide
			fmt.Printf("Consumiendo %d\n", d)
		} else {
			fmt.Println("Buffer vacio, esperando!!")
		}
		m.Unlock()
	}
}

func main() {
	m := new(sync.Mutex)
	go productor(m)  // 1 proceso
	go consumidor(m) // 1 proceso

	// pausa
	var in string
	fmt.Scanln(&in)
}
