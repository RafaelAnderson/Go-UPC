package main

import "fmt"

func main() {
	// Salto de línea
	fmt.Print("Ingrese un número:")
	var numero float64

	// Direcciona el dato a leer, & es el comando de dirección
	fmt.Scanf("%f", &numero) // lectura de consola

	doble := numero * 2

	fmt.Println("El doble de", numero, "es", doble)
}
