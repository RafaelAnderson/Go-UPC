// Comunicación Cliente - Servidor, ejecutar este despues de ejercicio04a.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	ln, _ := net.Listen("tcp", "localhost:9081")

	// Conectarse a un nodo servidor por el puerto 9080
	conexion, _ := net.Dial("tcp", "localhost:9080")

	defer conexion.Close()

	// Enviar datos al servidor
	fmt.Fprintln(conexion, "Nodo cliente")

	defer ln.Close()

	con, _ := ln.Accept()

	bufferIn := bufio.NewReader(con) // Leemos de la conexion

	msg, _ := bufferIn.ReadString('\n') // Leemos de la conexión

	fmt.Println(msg)
}
