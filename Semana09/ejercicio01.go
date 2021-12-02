// Algorithm 8.2: Conway's problem

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX = 9
	K   = 4
)

func compress(inC, pipe chan rune) {
	n := 0
	previous := <-inC
	for {
		// Continua la lectura
		c := <-inC
		if c == previous && n < MAX-1 {
			n++
		} else {
			if n > 0 {
				// Compress
				pipe <- rune(n + 49)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
}

func output(pipe, outC chan rune) {
	m := 0
	for {
		//c := <-pipe
		//outC <- c
		outC <- <-pipe // pipe gateway
		m++

		if m >= K {
			outC <- '\n'
			m = 0
		}
	}
}

func main() {
	inC := make(chan rune) // Canal de tipo string
	pipe := make(chan rune)
	outC := make(chan rune)

	// Goroutines
	go compress(inC, pipe)
	go output(pipe, outC)

	// Generador de datos del flujo
	go func() {
		// Semilla
		rand.Seed(time.Now().UTC().UnixNano())
		for {
			inC <- rune(rand.Intn(26) + 65) // Valor máximo 25
		}
	}()

	// Lectura del canal de salida
	for {
		fmt.Printf("%c", <-outC)
	}
}
