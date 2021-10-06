/*
				integer n <- 0
p								q
integer k1 <- 1				integer k2 <- 2
p1: n <- k1					q1: n <- k2
*/

package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	k1 := 1
	n = k1
}

func q() {
	k2 := 2
	n = k2
}

func main() {
	go p()
	go q()

	time.Sleep(time.Second)

	fmt.Println(n) // Debería salir 2
}
