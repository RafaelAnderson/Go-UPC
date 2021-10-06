// Algoritmo 3.7
package main

import "fmt"

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("Linea01 SNC P")
		fmt.Println("Linea02 SNC P")
		for wantp {
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
		for wantq {
			// Espera
		}
		wantp = true
		fmt.Println("Linea01 SC Q")
		fmt.Println("Linea02 SC Q")
		wantp = false
	}
}

func main() {

}
