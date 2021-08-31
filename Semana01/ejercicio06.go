package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Ingrese un número: ")
	bufferIn := bufio.NewReader(os.Stdin)
	dato, _ := bufferIn.ReadString('\n')
	dato = strings.TrimRight(dato, "\r\n")
	num, _ := strconv.Atoi(dato)

	evaluar(num)
}

func evaluar(numero int) {
	if numero%2 == 0 {
		fmt.Println("El número", numero, "es divisible por 2")
	} else if numero%3 == 0 {
		fmt.Println("El número", numero, "es divisible por 3")
	}
}
