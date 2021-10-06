package main

import "fmt"

func sumar(arr []int, c chan int) {
	suma := 0
	// range devuelve indice y valor
	for _, v := range arr {
		suma += v
	}
	c <- suma // envio de la suma al canal
}

func main() {
	// Arreglo
	arreglo := []int{6, 2, 8, 9, 4, 1}
	c := make(chan int)

	go sumar(arreglo[:len(arreglo)/2], c) // Primera Mitad
	go sumar(arreglo[len(arreglo)/2:], c) // Segunda Mitad

	a, b := <-c, <-c // Lecturas

	fmt.Println("Suma primera mitad: ", a)
	fmt.Println("Suma segunda mitad: ", b)
	fmt.Println("Suma total", a+b)
}
