package main

import "fmt"

// global
const igv float64 = 18.0

func main() {
	// ambito local
	var x string = "Hola a todos" // forma 1

	fmt.Println(x)

	dato := 20 // forma 2: corta de declarar e inicializar
	fmt.Println("Valor de la variable es", dato)

	var (
		nombre string
		edad   int
	)

	nombre = "Rafael"
	edad = 24

	fmt.Println("Mi nombre es", nombre, "y mi edad es", edad, "años")
}
