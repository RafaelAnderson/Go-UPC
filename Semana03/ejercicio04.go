// Estructura de Canal - Sincrona o asincrona
package main

import "fmt"

func envio(c chan int) {
	// envio de datos
	c <- 5
}

func main() {
	// Variable que representa el canal de enteros
	c := make(chan int) // sincrona

	go envio(c)
	fmt.Println(<-c) // Leo lo que llega por el canal, la instruccion espera
}
