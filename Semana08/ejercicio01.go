package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// estructura
type Alumno struct {
	Codigo string `json:"cod"`
	Nombre string `json:"nom"`
	Dni    int    `json:"dni"`
}

var listaAlumnos []Alumno

// Inicializa con esta data
func cargarAlumnos() {
	listaAlumnos = []Alumno{
		{"20211862", "Juan Merino", 40127558},
		{"20214714", "Christian Cueva", 50437562},
		{"20219576", "Maria Pérez", 70753412},
	}
}

// Funciones manejadoras de las solicitudes (Services); response devuelve algo a la petición
func mostrarHome(response http.ResponseWriter, request *http.Request) {
	log.Println("Ingresa a endpoint home")

	response.Header().Set("Content-Type", "text/html") // cabecera del mensaje
	// cuerpo del mensaje
	io.WriteString(response, `
	<html>
	<head></head>
	<body><h2>Mi API de alumnos</h2></body>
	</html>
	`)
}

func listarAlumnos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json") // contenido del mensaje json

	// Codificar
	jsonBytes, _ := json.MarshalIndent(listaAlumnos, "", " ") // Devuelve en formato JSON
	io.WriteString(response, string(jsonBytes))
}

func buscarAlumno(response http.ResponseWriter, request *http.Request) {
	// Recuperacion de parametros
	// codigo := request.FormValue("codigo") // http://localhost:9000/buscar?codigo=20211862
	cod := request.FormValue("codigo")
	// dni := request.FormValue("dni")

	log.Println(cod)

	response.Header().Set("Content-Type", "application/json")

	// logica de búsqueda
	for _, Alumno := range listaAlumnos {
		if Alumno.Codigo == cod {
			jsonBytes, _ := json.MarshalIndent(Alumno, "", " ")
			io.WriteString(response, string(jsonBytes))
		}
	}
}

func agregarAlumno(response http.ResponseWriter, request *http.Request) {
	// Validar que peticion se realice por el método POST
	if request.Method == "POST" {
		if request.Header.Get("Content-Type") == "application/json" {
			// Realiza la lógica de agregar alumno
			log.Println("Ingresa al método agregar alumno")

			// Recuperar el cuerpo del mensaje
			cuerpoMsg, err := ioutil.ReadAll(request.Body)

			if err != nil {
				http.Error(response, "Error interno al leer el body", http.StatusInternalServerError)
			}

			// Lógica para agregar alumno
			var oAlumno Alumno

			// Decodificar
			json.Unmarshal(cuerpoMsg, &oAlumno)
			listaAlumnos = append(listaAlumnos, oAlumno) // Sobreescribir la lista con el nuevo alumno

			// Return el mensaje de satisfactorio
			response.Header().Set("Content-Type", "application/json")

			io.WriteString(response, `
				{
					"respuesta":"Se agrega el alumno nuevo."
				}
			`)
		} else {
			http.Error(response, "Contenido inválido", http.StatusBadRequest)
		}
	} else {
		// lanzar error
		http.Error(response, "Método inválido", http.StatusMethodNotAllowed)
	}
}

// Función de enrutador
func manejadorSolicitudes() {
	mux := http.NewServeMux()

	// endopoints
	mux.HandleFunc("/home", mostrarHome)
	mux.HandleFunc("/listar", listarAlumnos)
	mux.HandleFunc("/buscar", buscarAlumno)
	mux.HandleFunc("/agregar", agregarAlumno)

	// puerto; log.Fatal en caso haya un error
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func main() {
	cargarAlumnos()
	manejadorSolicitudes()
}
