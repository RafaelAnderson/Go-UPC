package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Ingrese un número:")
	// Buffer utiliza un espacio en la memoria
	// Stdin - entrada por consola
	bufferIn := bufio.NewReader(os.Stdin)
	ingreso, er := bufferIn.ReadString('\n')

	if er != nil {
		fmt.Println("error:", er.Error())
		os.Exit(1)
	}

	//fmt.Print(ingreso)

	ingreso = strings.TrimRight(ingreso, "\r\n")
	// Atoi - string a entero
	numero, err := strconv.Atoi(ingreso)

	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	double := numero * 2

	fmt.Println("El doble del número es", double)
}
