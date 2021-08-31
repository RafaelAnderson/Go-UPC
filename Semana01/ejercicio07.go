package main

import "fmt"

func main() {
	// Declaración de un arreglo
	arreglo := [7]int{5, 3, 2, 7, 8, 9, 1}

	// Manejo del for
	for i, v := range arreglo {
		// Print con formato
		fmt.Printf("El valor de v es %d en la posición #%d\n", v, i)
	}

	// For infinito como while(true)
	i := 0
	for {
		fmt.Println(i)
		if i == 4 {
			break // salir del bucle
		}
		i++
	}

	// For con condición, while i<10
	for i < 10 {
		fmt.Println("continua")
		fmt.Println(i)
		i++
	}

	// For clásico
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
