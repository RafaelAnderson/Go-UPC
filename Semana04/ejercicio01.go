// Algoritmo 3.9
package main

import "fmt"

func p(c chan int) {
	// IN
	c <- 5
}

func main() {
	// sincrono ( <- Entrada, -> Salida)
	c := make(chan int)

	go p(c)
	go p(c)
	go p(c)

	n := <-c // OUT
	f := <-c
	m := <-c

	fmt.Println(n, f, m)
}
