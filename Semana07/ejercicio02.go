package main

import (
	"fmt"
	"sync"
)

func filosofo(id int, tenedor1, tenedor2 sync.Mutex) {
	for {
		fmt.Printf("Filosofo %d, pensando!!\n", id)
		tenedor1.Lock()
		tenedor2.Lock()
		fmt.Printf("Filosofo %d, comiendo!!!\n", id)
		tenedor1.Unlock()
		tenedor2.Unlock()
	}
}

func main() {

	tenedores := make([]sync.Mutex, 5)

	go filosofo(1, tenedores[0], tenedores[1]) // 1 proceso
	go filosofo(2, tenedores[1], tenedores[2]) // 1 proceso
	go filosofo(3, tenedores[2], tenedores[3]) // 1 proceso
	go filosofo(4, tenedores[3], tenedores[4]) // 1 proceso
	filosofo(5, tenedores[4], tenedores[0])    // 1 proceso
}
