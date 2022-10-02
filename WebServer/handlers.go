package main

import (
	"encoding/json"
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

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var metaData MetaData

	//Decode espera una interface generica
	err := decoder.Decode(&metaData)

	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v \n", metaData)
	// fmt.Println(metaData["name"])
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User

	//Decode espera una interface generica
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	response, err := user.toJSON()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(response)

}
