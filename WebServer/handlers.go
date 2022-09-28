package main

import (
	"fmt"
	"net/http"
)

// Handler para manejar la ruta principal
// Parametros: escritor w del tipo http.ResponseWriter y objeto request
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	//Impresion en el navegador
	//parametros: escritor-objeto encargado de responder al cliente
	//y mensaje escrito a travez del escritor
	fmt.Fprintf(w, "Hello World from handlers")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}
