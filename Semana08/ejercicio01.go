package main

import (
	"log"
	"net/http"
)

// estructura
type Alumno struct {
	Codigo string
	Nombre string
	Dni    int
}

var listarAlumnos []Alumno

// Inicializa con esta data
func cargarAlumnos() {
	listarAlumnos = []Alumno{
		{"20211862", "Juan Merino", 40127558},
		{"20214714", "Christian Cueva", 50437562},
		{"20219576", "Maria Pérez", 70753412},
	}
}

// Funciones manejadoras de las solicitudes (Services)
func mostrarHome(response http.ResponseWriter, request *http.Request) {
	log.Println("Ingresa a endpoint home")
}

// Función de enrutador
func manejadorSolicitudes() {
	mux := http.NewServeMux()

	// endopoints
	mux.HandleFunc("/home", mostrarHome)

	// puerto; log.Fatal en caso haya un error
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func main() {
	cargarAlumnos()
	manejadorSolicitudes()
}
