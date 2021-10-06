package main

import (
	"fmt"
	"time"
)

func saludar() {
	fmt.Println("Hola")
}

func main() {

	/* Programa Secuencial
	saludar()*/

	// Agregando concurrencia para que el proceso maneje su propia linea de tiempo
	go saludar()

	time.Sleep(time.Second) // Espera antes de terminar la ejecución
}
