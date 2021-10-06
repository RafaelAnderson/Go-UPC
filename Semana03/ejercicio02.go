// Algoritmo 3.6
package main

import "fmt"

// Iniciaria cualquiera
var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("Linea01 SNC P")
		fmt.Println("Linea02 SNC P")
		for wantq {
			// Espera
		}
		wantp = true
		fmt.Println("Linea01 SC P")
		fmt.Println("Linea02 SC P")
		wantp = false
	}
}

func q() {
	for {
		fmt.Println("Linea01 SNC Q")
		fmt.Println("Linea02 SNC Q")
		for wantp {
			// Espera
		}
		wantq = true
		fmt.Println("Linea01 SC Q")
		fmt.Println("Linea02 SC Q")
		wantq = false
	}
}

func main() {
	go p()
	q()
}
