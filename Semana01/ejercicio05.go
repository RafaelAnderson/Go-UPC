package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese su nombre:")
	bufferIn := bufio.NewReader(os.Stdin)

	// _ es para ignorar el elemento
	nombre, _ := bufferIn.ReadString('\n')
	nombre = strings.TrimRight(nombre, "\r\n")

	menu :=
		`
		**** MENU ****
		[1] Pizza
		[2] Empanada
		¿Cuál es tu pedido?
	`
	println("Bienvenido", nombre)
	println(menu)
	opc, _ := bufferIn.ReadString('\n')
	opc = strings.TrimRight(opc, "\r\n") //limpiar

	switch opc {
	case "1":
		fmt.Println("Su pizza estará en 30 minutos")
	case "2":
		fmt.Println("Su empanada estará en 5 minutos")
	}

	fmt.Println("Gracias por su preferencia")
}
