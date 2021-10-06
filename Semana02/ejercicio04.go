package main

import "fmt"

func saludar(p int) {
	fmt.Println("Hola", p)
}

func pausa() {
	var input string
	fmt.Scanln(&input) // Espera que se escriba para finalizar
}

func main() {
	for i := 1; i <= 5; i++ {
		go saludar(i)
	}

	fmt.Println("Procesando ...")
	pausa()
}
