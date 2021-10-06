// Algoritmo 2.4
/*
				integer n <- 0
	p								q
integer temp				integer temp
p1: temp <- n				q1: temp <- n
p2: n <- temp + 1			q2: n <- temp + 1
*/
package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	var temp1 int
	temp1 = n
	n = temp1 + 1

	fmt.Println("P Temp = ", temp1)
}

func q() {
	var temp2 int
	temp2 = n
	n = temp2 + 1

	fmt.Println("Q Temp = ", temp2)
}

/*
func pq() {
	var temp int
	temp = n
	time.Sleep(time.Second) // 2da forma
	n = temp + 1
}
*/
func main() {
	go p()
	go q()
	time.Sleep(time.Second)
	fmt.Println(n)
}
