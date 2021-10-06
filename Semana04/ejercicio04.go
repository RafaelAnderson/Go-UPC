package main

import (
	"fmt"
	"time"
)

// Equivalencia a c3
func proceso3(c3 chan string) {
	for {
		c3 <- "Datos enviados desde el proceso 3"
		// time.Sleep(time.Second * 4) // Verificar tiempo de espera
	}
}

func main() {
	// Crear canales
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// Funcion anonima
	go func() {
		for {
			c1 <- "Datos enviados desde el proceso 1"
			// time.Sleep(time.Second * 3) // Verificar tiempo de espera
		}
	}()

	go func() {
		for {
			c2 <- "Datos enviados desde el proceso 2"
			// time.Sleep(time.Second * 2) // Verificar tiempo de espera
		}
	}()

	go proceso3(c3) // Función con nombre

	// Interceptar
	func() {
		for {
			select {
			case msg1 := <-c1: // Recepcionar del canal 1 y capturar en la variable
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-c3:
				fmt.Println("Mensaje 3") // No se guarda, se ignora
				// Manejar tiempos de respuesta
			case <-time.After(time.Second): // Si en un segundo no detecta datos
				fmt.Println("Tiempo de espera caducado")
			}
		}
	}()

	// Para y fin del programa
	var input string
	fmt.Scanln(&input)
}
