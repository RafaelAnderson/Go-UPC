package main

import "fmt"

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func imprimir(c chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string)

	go ping(c)
	go imprimir(c)

	// Paro y termino
	var input string
	fmt.Scanln(&input)
}
