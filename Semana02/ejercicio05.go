// Algoritmo 2.9
/*
				integer n <- 0
	p								q
integer temp				integer temp
p1: do 10 times				q1: do 10 times
p2: temp <- n				q2: temp <- n
p3: n <- temp + 1			q3: n <- temp + 1
*/
package main

import (
	"fmt"
	"time"
)

var n int

func pq() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		// duracion := time.Duration(rand.Intn(251))
		// time.Sleep(time.Millisecond + duracion)
		n = temp + 1
	}
}

func main() {
	go pq() // Proceso 1
	go pq() // Proceso 2

	time.Sleep(time.Second)
	fmt.Println(n)
}
