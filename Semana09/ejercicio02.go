// Canal Síncrono
package main

import "fmt"

func enviarNumero(chNum chan int) {
	// Comportamiento limitado
	for i := 0; i < 10; i++ {
		chNum <- i
	}
	close(chNum) // cerrar el canal
}

func main() {
	chNum := make(chan int) // Modelo de canal síncrono

	go enviarNumero(chNum)

	// Constante - el programa principal depende de este proceso
	for num := range chNum {
		fmt.Println(num)
	}
}
