// Para ingresar a la sección crítica, será de manera exclusiva

package main

import "fmt"

var turno int = 1

func p() {
	for {
		fmt.Println("Linea01 SNC P") // SNC: Seccion no critica
		fmt.Println("Linea02 SNC P")
		// Condicion PreProtocol
		for turno != 1 {
			// Espera
		}
		// Seccion Critica
		fmt.Println("Linea 01 SC P") // SC: Seccion critica
		fmt.Println("Linea 02 SC P")
		// PostProtocol
		turno = 2
	}
}

func q() {
	for {
		fmt.Println("Linea01 SNC Q")
		fmt.Println("Linea02 SNC Q")
		for turno != 2 {
			// Espera
		}
		fmt.Println("Linea 01 SC Q")
		fmt.Println("Linea 02 SC Q")

		turno = 1
	}
}

func main() {
	go p() // goroutine
	q()
}
