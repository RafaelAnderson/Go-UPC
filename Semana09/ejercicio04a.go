// Comunicación Cliente - Servidor - Canal de comunicación bidireccional
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Servidor
	ln, err := net.Listen("tcp", "localhost:9080") // Punto de conexión
	if err != nil {
		fmt.Println("Falla al resolver la dirección de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}
	// Diferidos
	defer ln.Close()

	// Aceptar la conexión
	conexion, _ := ln.Accept() // Devuelve net y error

	defer conexion.Close()

	// Lectura de lo que llega al server
	bufferIn := bufio.NewReader(conexion) // Leemos de la conexion

	msg, _ := bufferIn.ReadString('\n') // Leemos de la conexión

	fmt.Println(msg)

	// Respuesta
	// Conectarse a un nodo servidor por el puerto 9080
	conexion2, _ := net.Dial("tcp", "localhost:9081")

	defer conexion2.Close()

	// Enviar datos al servidor
	fmt.Fprintln(conexion2, "Gracias cliente")
}
