package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := new(sync.Mutex) // Estructura que permite aplicar aspectos de exclusion (Método bloquear y método desbloquear)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock() // bloquear el acceso de algún otro proceso al actual
			fmt.Println(i, " SC line 1")
			time.Sleep(time.Second)
			fmt.Println(i, " SC line 2")
			m.Unlock() // Desbloqueamos dando la posibilidad a que otro proceso ingrese a la sección crítica
		}(i)
	}

	// Pausa
	var input string
	fmt.Scanln(&input)
}
