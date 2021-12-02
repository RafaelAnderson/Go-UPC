// Canal Asíncrono
package main

import "fmt"

func main() {
	//ch := make(chan int) // Síncrono
	ch := make(chan int, 1) // Asíncrono
	ch <- 5

	fmt.Println(<-ch)
}
