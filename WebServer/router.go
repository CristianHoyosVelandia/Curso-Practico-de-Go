package main

import (
	"net/http"
)

// Struct router para hacer request en el servidor
type Router struct {
	//Reglas para definir de que rutas pasan a que handler, mapa que pasa de strings a handler
	//mapa que tenga como llaves strings y que mapee a HandlerFunc
	rules map[string]map[string]http.HandlerFunc
}

// forma de instanciar el router, similar al NewServer() del archivo servidor.go
func NewRouter() *Router {
	return &Router{
		//a diferencia del servidor, aqui el router debe empezar en un estado vacio, creamos un mapa vacio
		// rules: make(map[string]map[string]http.HandlerFunc),

		//ahora como necesitamos crear un mapa de mapas que contenga los handlers agregamos a la sentencia de rules
		// el map map esta definido como METODO, /URL y handler
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// creo una funcion receiver para que sea detectada por el handler,
// aca detecto si el handler existe en el mapa de rutas.
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

// Metodo ServeHTTP de router para poder implementar en el handler el atributo s.router en server.go
// parametros: el primero es el escritor, el segundo es el request en donde viene la informacion
// no olvidar colocar ServeHTTP con letras mayusculas
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	//manejo del mensaje de manera dinamica
	//r es la referencia al router y usar la funcion FindHandler
	//el FindHandler compara el request con el mapa de reglas para saber si existe o no.
	//los valores son asignados a las variables 'handler' y 'exist'
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)

	//Evaluacion del booleano del handler para saber si existe o no, error 404
	if !exist {
		//WriteHeader es para indicar el status del request
		w.WriteHeader(http.StatusNotFound)
		//el return nos ayuda a romper la funcion si esto no existe el handler
		return
	}

	//verificio si el methodo existe para esa url, en caso de que no exista.
	//imprimimos un header diciendo que el methodo no existe
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//handler para enviar objeto w y request
	handler(w, request)

}
